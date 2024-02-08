package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func SpectrometerTestEvent(outputFilePath string) {
	//This routine is used to test the Geiger Counter JSON

	jsondata := datamodels.SpectrometerEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeId = 11
	jsondata.EventHeader.SensorType.SensorTypeDescription = "spectrometer"

	elementdata := make([]datamodels.ElementReading, 60)

	elementName := ""

	for i := 0; i < 60; i++ {

		switch i {
		case 0:
			elementName = "Li"
		case 1:
			elementName = "Li-Aerosol"

		case 2:
			elementName = "He"

		case 3:
			elementName = "Ne"

		case 4:
			elementName = "Ar"

		case 5:
			elementName = "Kr"

		case 6:
			elementName = "Pb"

		case 7:
			elementName = "Xe"

		case 8:
			elementName = "Rn"

		case 9:
			elementName = "Og"

		case 10:
			elementName = "CO2"

		case 11:
			elementName = "CO"

		case 12:
			elementName = "C3H8"

		case 13:
			elementName = "C2H2"

		case 14:
			elementName = "H"

		case 15:
			elementName = "N"

		case 16:
			elementName = "O"

		case 17:
			elementName = "CnH2n+2"

		case 18:
			elementName = "CH4"

		case 19:
			elementName = "C2H6"

		case 20:
			elementName = "C4H10"

		case 21:
			elementName = "Th"

		case 22:
			elementName = "C5H12"

		case 23:
			elementName = "C6H14"

		case 24:
			elementName = "C"

		case 25:
			elementName = "S"

		case 26:
			elementName = "KNO3"

		case 27:
			elementName = "U"

		case 28:
			elementName = "Pu"

		case 29:
			elementName = "Be"

		case 30:
			elementName = "B"

		case 31:
			elementName = "F"

		case 32:
			elementName = "Na"

		case 33:
			elementName = "Mg"

		case 34:
			elementName = "Al"

		case 35:
			elementName = "Si"

		case 36:
			elementName = "P"

		case 37:
			elementName = "Cl"

		case 38:
			elementName = "K"

		case 39:
			elementName = "Ca"

		case 40:
			elementName = "Sc"

		case 41:
			elementName = "Ti"

		case 42:
			elementName = "V"

		case 43:
			elementName = "Cr"

		case 44:
			elementName = "Mn"

		case 45:
			elementName = "Fe"

		case 46:
			elementName = "Co"

		case 47:
			elementName = "Ni"

		case 48:
			elementName = "Cu"

		case 49:
			elementName = "Zn"

		case 50:
			elementName = "Ga"

		case 51:
			elementName = "Ge"

		case 52:
			elementName = "As"

		case 53:
			elementName = "Se"

		case 54:
			elementName = "Br"

		case 55:
			elementName = "Rb"

		case 56:
			elementName = "Sr"

		case 57:
			elementName = "Y"

		case 58:
			elementName = "Pd"

		case 59:
			elementName = "Ag"

		}

		elementdata[i].ElementName = elementName
		elementdata[i].PartsPerMillion = 2.769

	}

	jsondata.SensorData = make([]datamodels.SpectrometerReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].Elements = elementdata
	jsondata.SensorData[0].Explosives = 1.439
	jsondata.SensorData[0].Opocs = 3.514
	jsondata.SensorData[0].SpectrometerState = 1
	jsondata.SensorData[0].Urates = 14.3287539
	jsondata.SensorData[0].WeaponsGradeNuclearMaterial = .00567

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].Elements = elementdata
	jsondata.SensorData[1].Explosives = 1.439
	jsondata.SensorData[1].Opocs = 3.514
	jsondata.SensorData[1].SpectrometerState = 1
	jsondata.SensorData[1].Urates = 14.3287539
	jsondata.SensorData[1].WeaponsGradeNuclearMaterial = .00567

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
