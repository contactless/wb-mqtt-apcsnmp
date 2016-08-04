package main

import (
	"github.com/contactless/wb-mqtt-apcsnmp/apcsnmp"
	"flag"
	"github.com/contactless/wbgo"
	"time"
)

func main() {
	snmpAddress := flag.String("snmp", "localhost", "snmp device address (host or host:port)")
	broker := flag.String("broker", "tcp://localhost:1883", "MQTT broker url")
        devName := flag.String("devname", "apcupc", "MQTT driver device name for this instance")
	debug := flag.Bool("debug", false, "Enable debugging")
	flag.Parse()
	if *debug {
		wbgo.SetDebuggingEnabled(true)
	}

        for {
                time.Sleep(1 * time.Second)
        }

	if driver, err := apcsnmp.NewApcSnmpDriver(*snmpAddress, *broker, *devName); err != nil {
		panic(err)
	} else {
		if err := driver.Start(); err != nil {
			panic(err)
		}
		for {
			time.Sleep(1 * time.Second)
		}
	}
}
