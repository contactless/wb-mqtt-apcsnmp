package apcsnmp

var apcVars []*ApcVar = []*ApcVar{
	&ApcVar{"Internal Temperature", "upsHighPrecBatteryTemperature.0", "°C", Scale(0.1)},
	&ApcVar{"Runtime Remaining", "upsAdvBatteryRunTimeRemaining.0", "", ConvDuration},
	// 'Power' section
	&ApcVar{"Input Voltage L1", "upsPhaseInputVoltage.1.1.1", "VAC", Num},
	&ApcVar{"Input Voltage L2", "upsPhaseInputVoltage.1.1.2", "VAC", Num},
	&ApcVar{"Input Voltage L3", "upsPhaseInputVoltage.1.1.3", "VAC", Num},
	&ApcVar{"Bypass Input Voltage L1", "upsPhaseInputVoltage.2.1.1", "VAC", Num},
	&ApcVar{"Bypass Input Voltage L2", "upsPhaseInputVoltage.2.1.2", "VAC", Num},
	&ApcVar{"Bypass Input Voltage L3", "upsPhaseInputVoltage.2.1.3", "VAC", Num},
	&ApcVar{"Output Voltage L1", "upsPhaseOutputVoltage.1.1.1", "VAC", Num},
	&ApcVar{"Output Voltage L2", "upsPhaseOutputVoltage.1.1.2", "VAC", Num},
	&ApcVar{"Output Voltage L3", "upsPhaseOutputVoltage.1.1.3", "VAC", Num},
	&ApcVar{"Current L1", "upsPhaseInputCurrent.1.1.1", "Amps", Scale(0.1)},
	&ApcVar{"Current L2", "upsPhaseInputCurrent.1.1.2", "Amps", Scale(0.1)},
	&ApcVar{"Current L3", "upsPhaseInputCurrent.1.1.3", "Amps", Scale(0.1)},
	&ApcVar{"Input Frequency", "upsHighPrecInputFrequency.0", "Hz", Scale(0.1)},
	// 'Load' section
	&ApcVar{"Output Load L1", "upsPhaseOutputLoad.1.1.1", "kVA", Scale(0.001)},
	&ApcVar{"Output Load L2", "upsPhaseOutputLoad.1.1.2", "kVA", Scale(0.001)},
	&ApcVar{"Output Load L3", "upsPhaseOutputLoad.1.1.3", "kVA", Scale(0.001)},
	&ApcVar{"Output Percent Load L1", "upsPhaseOutputPercentLoad.1.1.1", "%kVA", Num},
	&ApcVar{"Output Percent Load L2", "upsPhaseOutputPercentLoad.1.1.2", "%kVA", Num},
	&ApcVar{"Output Percent Load L3", "upsPhaseOutputPercentLoad.1.1.3", "%kVA", Num},
	&ApcVar{"Output Percent Power L1", "upsPhaseOutputPercentPower.1.1.1", "%Watts", Num},
	&ApcVar{"Output Percent Power L2", "upsPhaseOutputPercentPower.1.1.2", "%Watts", Num},
	&ApcVar{"Output Percent Power L3", "upsPhaseOutputPercentPower.1.1.3", "%Watts", Num},
	&ApcVar{"Output Current L1", "upsPhaseOutputCurrent.1.1.1", "Amps", Scale(0.1)},
	&ApcVar{"Output Current L2", "upsPhaseOutputCurrent.1.1.2", "Amps", Scale(0.1)},
	&ApcVar{"Output Current L3", "upsPhaseOutputCurrent.1.1.3", "Amps", Scale(0.1)},
	&ApcVar{"Output Frequency", "upsPhaseOutputFrequency.1", "Hz", Scale(0.1)},
	// 'Battery' section
	&ApcVar{"Battery Capacity", "upsHighPrecBatteryCapacity.0", "%", Scale(0.1)},
	&ApcVar{"Nominal Battery Voltage", "upsHighPrecBatteryNominalVoltage.0", "VDC", Scale(0.1)},
	&ApcVar{"Actual Battery Bus Voltage", "upsHighPrecBatteryActualVoltage.0", "VDC", Scale(0.1)},
	&ApcVar{"Battery Current", "upsHighPrecBatteryCurrent.0", "Amps", Scale(0.1)},
	// FIXME - External Battery Cabinet Rating
	&ApcVar{"Number of Batteries", "upsAdvBatteryNumOfBattPacks.0", "", Scale(1)},
	&ApcVar{"Number of Bad Batteries", "upsAdvBatteryNumOfBadBattPacks.0", "", Num},
	// 'Intelligence module' section ("-" as units means 'text' value type)
	&ApcVar{"Firmware Revision", "upsDiagIMFirmwareRev.1", "-", AsIs},
	&ApcVar{"Serial Number", "upsDiagIMSerialNum.1", "-", AsIs},
	&ApcVar{"Manufacture Date", "upsDiagIMManufactureDate.1", "-", AsIs},
	&ApcVar{"Hardware Revision", "upsDiagIMHardwareRev.1", "-", AsIs},
	// Status
	&ApcVar{"Output Status", "upsBasicOutputStatus.0", "-", AsIs},
}
