package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

func SpectrometerInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Spectrometer sensor

	for i := range *qubzMatrix {

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Spectrometer.EventState = appconstants.SENSOR_STATE_CURRENT
		(*qubzMatrix)[i].Spectrometer.SpectrometerState = 1
		(*qubzMatrix)[i].Spectrometer.Explosives = 0.045
		(*qubzMatrix)[i].Spectrometer.Opocs = 0.021
		(*qubzMatrix)[i].Spectrometer.Urates = 480
		(*qubzMatrix)[i].Spectrometer.WeaponsGradeNuclearMaterial = 0.0003

		//Add elements and initial values
		(*qubzMatrix)[i].Spectrometer.Elements = make([]datamodels.ElementReading, 60)

		var elementName string = ""
		var partsPerMillion float64 = 0

		for elindex := 0; elindex < 60; elindex++ {

			switch elindex {
			case 0:
				elementName = "Li"
				partsPerMillion = 79

			case 1:
				elementName = "Li-Aerosol"
				partsPerMillion = 1.13245

			case 2:
				elementName = "He"
				partsPerMillion = 2.85739

			case 3:
				elementName = "Ne"
				partsPerMillion = 3.00562

			case 4:
				elementName = "Ar"
				partsPerMillion = 3.00293

			case 5:
				elementName = "Kr"
				partsPerMillion = 1.0027

			case 6:
				elementName = "Xe"
				partsPerMillion = 0.215

			case 7:
				elementName = "Rn"
				partsPerMillion = 0.71000

			case 8:
				elementName = "Og"
				partsPerMillion = 0.7941

			case 9:
				elementName = "CO2"
				partsPerMillion = 219.47295

			case 10:
				elementName = "CO"
				partsPerMillion = 16.27364

			case 11:
				elementName = "C3H8"
				partsPerMillion = 4000.72618

			case 12:
				elementName = "C2H2"
				partsPerMillion = 2987.82719

			case 13:
				elementName = "H"
				partsPerMillion = 11.57362

			case 14:
				elementName = "N"
				partsPerMillion = 643748.87295

			case 15:
				elementName = "O"
				partsPerMillion = 257042.00872

			case 16:
				elementName = "CnH2N+2"
				partsPerMillion = 152893.90098

			case 17:
				elementName = "CH4"
				partsPerMillion = 827.27000

			case 18:
				elementName = "C2H6"
				partsPerMillion = 315.87295

			case 19:
				elementName = "C4H10"
				partsPerMillion = 184.02039

			case 20:
				elementName = "C5H12"
				partsPerMillion = 228.98173

			case 21:
				elementName = "C6H14"
				partsPerMillion = 985.98273

			case 22:
				elementName = "C"
				partsPerMillion = 37.82098

			case 23:
				elementName = "S"
				partsPerMillion = 0.92671

			case 24:
				elementName = "KNO3"
				partsPerMillion = 1.13245
			}

			(*qubzMatrix)[i].Spectrometer.Elements[elindex].ElementName = elementName
			(*qubzMatrix)[i].Spectrometer.Elements[elindex].PartsPerMillion = partsPerMillion
		}

	}

}
