# APC UPS snmp driver


Конфигурация WB (заменить IP на нужный):

```
echo 'APCSNMP_OPTIONS="-snmp 10.6.2.111"' > /etc/default/wb-mqtt-apcsnmp
service wb-mqtt-apcsnmp restart
```
