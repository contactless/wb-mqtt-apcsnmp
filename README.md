# APC UPS snmp driver


Конфигурация WB (заменить IP на нужный):

```
echo 'APCSNMP_OPTIONS="-snmp 10.6.2.111 -devname apcups1"' > /etc/default/wb-mqtt-apcsnmp
service wb-mqtt-apcsnmp restart
```

для того, чтобы создать одно устройство с именем apcups1, связанное с агентом 10.6.2.111.


Для того, чтобы создать несколько устройств для разных агентов, дописываем необходимые аргументы
нужное количество раз, разделяя символом "!". Например:

```
echo 'APCSNMP_OPTIONS="-snmp 10.6.2.111 -devname apcups1 ! -snmp 10.6.2.112 -devname apcups2 ! -snmp 10.6.2.113 -devname apcups3"' > /etc/default/wb-mqtt-apcsnmp
service wb-mqtt-apcsnmp restart
```
