package apcsnmp

import (
	"fmt"
	"github.com/alouca/gosnmp"
	"github.com/contactless/wbgo"
	"log"
	"strconv"
	"time"
)

const (
	DEV_NAME = "apcups"
)

type ConverterFunc func(in string) string

func AsIs(in string) string { return in }

func Scale(coef float64) ConverterFunc {
	return func(in string) string {
		v, err := strconv.ParseFloat(in, 64)
		if err != nil {
			wbgo.Warn.Printf("can't convert numeric value: %s", in)
			return in
		}
		if coef == 1 {
			return strconv.FormatInt(int64(v), 10)
		}
		return strconv.FormatFloat(v*coef, 'f', 1, 64)
	}
}

var Num = Scale(1)

func ConvDuration(in string) string {
	v, err := strconv.ParseInt(in, 10, 0)
	if err != nil {
		wbgo.Warn.Printf("can't convert duration value: %v", v)
		return in
	}
	return (time.Duration(v) * 10 * time.Millisecond).String()
}

type ApcVar struct {
	MqttName, Name, Units string
	Converter             ConverterFunc
}

// TBD: make a command line option for DumpVars
func DumpVars(snmp *gosnmp.GoSNMP) {
	for _, apcVar := range apcVars {
		resp, err := snmp.Get(mibMap[apcVar.Name])
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range resp.Variables {
			val := fmt.Sprintf("%v", v.Value)
			log.Printf("Response for %s: %s : %s : %s \n",
				apcVar.MqttName, v.Name,
				apcVar.Converter(val),
				v.Type.String())
		}
	}
}

type ApcUpsDevice struct {
	wbgo.DeviceBase
	snmp         *gosnmp.GoSNMP
	cachedValues map[string]string
}

func newApcUpsDevice(snmp *gosnmp.GoSNMP, name, title string) *ApcUpsDevice {
	return &ApcUpsDevice{
		DeviceBase:   wbgo.DeviceBase{DevName: name, DevTitle: title},
		snmp:         snmp,
		cachedValues: make(map[string]string),
	}
}

// values are only received from UPS
func (dev *ApcUpsDevice) AcceptValue(name, value string) {}

// read-only device
func (dev *ApcUpsDevice) AcceptOnValue(name, value string) bool { return false }

func (dev *ApcUpsDevice) IsVirtual() bool { return false }

func (dev *ApcUpsDevice) Poll() {
	for _, apcVar := range apcVars {
		resp, err := dev.snmp.Get(mibMap[apcVar.Name])
		if err != nil {
			wbgo.Error.Printf("error reading %s: %s", apcVar.Name, err)
			continue
		}
		if len(resp.Variables) == 0 {
			continue
		}
		val := apcVar.Converter(fmt.Sprintf("%v", resp.Variables[0].Value))
		oldVal, found := dev.cachedValues[apcVar.MqttName]
		switch {
		case !found:
			paramType := "value"
			switch {
			case apcVar.Units == "-":
				paramType = "text"
			case apcVar.Units != "":
				paramType += ":" + apcVar.Units
			}
			dev.Observer.OnNewControl(
				dev, apcVar.MqttName, paramType, val, true, -1, true)
		case val != oldVal:
			dev.Observer.OnValue(dev, apcVar.MqttName, val)
		}
		dev.cachedValues[apcVar.MqttName] = val
	}
}

type ApcUpsModel struct {
	wbgo.ModelBase
	snmp *gosnmp.GoSNMP
	dev  *ApcUpsDevice
}

func NewApcUpsModel(snmp *gosnmp.GoSNMP) *ApcUpsModel {
	return &ApcUpsModel{snmp: snmp}
}

func (model *ApcUpsModel) Start() error {
	devTitle := "APC UPS"
	resp, err := model.snmp.Get(mibMap["upsBasicIdentModel.0"])
	if err != nil || len(resp.Variables) == 0 {
		wbgo.Error.Printf("couldn't get model string: %s", err)
	} else {
		devTitle = resp.Variables[0].Value.(string)
	}
	model.dev = newApcUpsDevice(model.snmp, DEV_NAME, devTitle)
	model.Observer.OnNewDevice(model.dev)
	return nil
}

func (model *ApcUpsModel) Poll() {
	if model.dev != nil {
		model.dev.Poll()
	}
}
