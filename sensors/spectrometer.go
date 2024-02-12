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

			case 25:
				elementName = "U"
				partsPerMillion = 1.13245

			case 26:
				elementName = "Pu"
				partsPerMillion = .00345

			case 27:
				elementName = "Be"
				partsPerMillion = 2.67500

			case 28:
				elementName = "B"
				partsPerMillion = 35.98723

			case 29:
				elementName = "F"
				partsPerMillion = 35.98723

			case 30:
				elementName = "Na"
				partsPerMillion = 2097.72846

			case 31:
				elementName = "Mg"
				partsPerMillion = 97.87263

			case 32:
				elementName = "Al"
				partsPerMillion = 393827.74628

			case 33:
				elementName = "Si"
				partsPerMillion = 35.98723

			case 34:
				elementName = "P"
				partsPerMillion = 11876.92738

			case 35:
				elementName = "Cl"
				partsPerMillion = 7892.98263

			case 36:
				elementName = "K"
				partsPerMillion = 841.72638

			case 37:
				elementName = "Ca"
				partsPerMillion = 219726.82736

			case 38:
				elementName = "Sc"
				partsPerMillion = 90726.72169

			case 39:
				elementName = "Ti"
				partsPerMillion = 479328.18374

			case 40:
				elementName = "V"
				partsPerMillion = 76917.01739

			case 41:
				elementName = "Cr"
				partsPerMillion = 23718.00928

			case 42:
				elementName = "Mn"
				partsPerMillion = 3287.08901

			case 43:
				elementName = "Fe"
				partsPerMillion = 621076.46275

			case 44:
				elementName = "Co"
				partsPerMillion = 76823.91728

			case 45:
				elementName = "Ni"
				partsPerMillion = 276982.47262

			case 46:
				elementName = "Cu"
				partsPerMillion = 189726.92702

			case 47:
				elementName = "Zn"
				partsPerMillion = 315927.27380

			case 48:
				elementName = "Ga"
				partsPerMillion = 1378.98428

			case 49:
				elementName = "Ge"
				partsPerMillion = 6529.273629

			case 50:
				elementName = "As"
				partsPerMillion = 0.78920

			case 51:
				elementName = "Se"
				partsPerMillion = .92736

			case 52:
				elementName = "Br"
				partsPerMillion = 5.82791

			case 53:
				elementName = "Rb"
				partsPerMillion = 12.72638

			case 54:
				elementName = "Sr"
				partsPerMillion = 23.72695

			case 55:
				elementName = "Y"
				partsPerMillion = 35.98723

			case 56:
				elementName = "Pd"
				partsPerMillion = 35.98723

			case 57:
				elementName = "Ag"
				partsPerMillion = 15872.92846

			case 58:
				elementName = "Pb"
				partsPerMillion = 287927.75283

			case 59:
				elementName = "Th"
				partsPerMillion = 35.98723
			}

			(*qubzMatrix)[i].Spectrometer.Elements[elindex].ElementName = elementName
			(*qubzMatrix)[i].Spectrometer.Elements[elindex].PartsPerMillion = partsPerMillion
		}

	}

}
