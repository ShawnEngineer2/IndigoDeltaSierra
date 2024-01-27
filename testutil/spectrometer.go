package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func SpectrometerTestEvent(outputFilePath string) {
	//This routine is used to test the Geiger Counter JSON

	jsondata := datamodels.SpectrometerEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	elementdata := make([]datamodels.ElementReading, 60)

	elementName := ""

	for i := 0; i < 60; i++ {

		switch i {
		case 0:
			elementName = "Li"
			break
		case 1:
			elementName = "Li-Aerosol"
			break
		case 2:
			elementName = "He"
			break
		case 3:
			elementName = "Ne"
			break
		case 4:
			elementName = "Ar"
			break
		case 5:
			elementName = "Kr"
			break
		case 6:
			elementName = "Pb"
			break
		case 7:
			elementName = "Xe"
			break
		case 8:
			elementName = "Rn"
			break
		case 9:
			elementName = "Og"
			break
		case 10:
			elementName = "CO2"
			break
		case 11:
			elementName = "CO"
			break
		case 12:
			elementName = "C3H8"
			break
		case 13:
			elementName = "C2H2"
			break
		case 14:
			elementName = "H"
			break
		case 15:
			elementName = "N"
			break
		case 16:
			elementName = "O"
			break
		case 17:
			elementName = "CnH2n+2"
			break
		case 18:
			elementName = "CH4"
			break
		case 19:
			elementName = "C2H6"
			break
		case 20:
			elementName = "C4H10"
			break
		case 21:
			elementName = "Th"
			break
		case 22:
			elementName = "C5H12"
			break
		case 23:
			elementName = "C6H14"
			break
		case 24:
			elementName = "C"
			break
		case 25:
			elementName = "S"
			break
		case 26:
			elementName = "KNO3"
			break
		case 27:
			elementName = "U"
			break
		case 28:
			elementName = "Pu"
			break
		case 29:
			elementName = "Be"
			break
		case 30:
			elementName = "B"
			break
		case 31:
			elementName = "F"
			break
		case 32:
			elementName = "Na"
			break
		case 33:
			elementName = "Mg"
			break
		case 34:
			elementName = "Al"
			break
		case 35:
			elementName = "Si"
			break
		case 36:
			elementName = "P"
			break
		case 37:
			elementName = "Cl"
			break
		case 38:
			elementName = "K"
			break
		case 39:
			elementName = "Ca"
			break
		case 40:
			elementName = "Sc"
			break
		case 41:
			elementName = "Ti"
			break
		case 42:
			elementName = "V"
			break
		case 43:
			elementName = "Cr"
			break
		case 44:
			elementName = "Mn"
			break
		case 45:
			elementName = "Fe"
			break
		case 46:
			elementName = "Co"
			break
		case 47:
			elementName = "Ni"
			break
		case 48:
			elementName = "Cu"
			break
		case 49:
			elementName = "Zn"
			break
		case 50:
			elementName = "Ga"
			break
		case 51:
			elementName = "Ge"
			break
		case 52:
			elementName = "As"
			break
		case 53:
			elementName = "Se"
			break
		case 54:
			elementName = "Br"
			break
		case 55:
			elementName = "Rb"
			break
		case 56:
			elementName = "Sr"
			break
		case 57:
			elementName = "Y"
			break
		case 58:
			elementName = "Pd"
			break
		case 59:
			elementName = "Ag"
			break
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

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].Elements = elementdata
	jsondata.SensorData[1].Explosives = 1.439
	jsondata.SensorData[1].Opocs = 3.514
	jsondata.SensorData[1].SpectrometerState = 1
	jsondata.SensorData[1].Urates = 14.3287539

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
