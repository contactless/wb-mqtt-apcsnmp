package apcsnmp

import (
	"github.com/alouca/gosnmp"
	"github.com/contactless/wbgo"
)

const (
	DRIVER_CLIENT_ID = "apcups"
	POLL_INTERVAL_MS = 1000
)

func NewApcSnmpDriver(snmpAddress, brokerAddress string) (*wbgo.Driver, error) {
	snmp, err := gosnmp.NewGoSNMP(snmpAddress, "public", gosnmp.Version2c, 5)
	if err != nil {
		wbgo.Error.Fatal(err)
	}
	model := NewApcUpsModel(snmp)
	driver := wbgo.NewDriver(model, wbgo.NewPahoMQTTClient(brokerAddress, DRIVER_CLIENT_ID, false))
	driver.SetPollInterval(POLL_INTERVAL_MS)
	return driver, nil
}
