package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func SpectrometerInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Spectrometer sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Spectrometer.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].Spectrometer.SpectrometerState = 1
	(*qubzMatrix)[matrixIndex].Spectrometer.Explosives = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_EXPLOSIVES, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Spectrometer.Opocs = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_OPOCS, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Spectrometer.Urates = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_URATES, qubzStateDS, consoleLogger, fileLogger)
	(*qubzMatrix)[matrixIndex].Spectrometer.WeaponsGradeNuclearMaterial = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_WGNM, qubzStateDS, consoleLogger, fileLogger)

	//Add elements and initial values
	(*qubzMatrix)[matrixIndex].Spectrometer.Elements = make([]datamodels.ElementReading, 60)

	var elementName string = ""
	var partsPerMillion float64 = 0

	for elindex := 0; elindex < 60; elindex++ {

		switch elindex {
		case 0:
			elementName = "Li"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_LITHIUM, qubzStateDS, consoleLogger, fileLogger)

		case 1:
			elementName = "Li-Aerosol"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_LITHIUM_AEROSOLS, qubzStateDS, consoleLogger, fileLogger)

		case 2:
			elementName = "He"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_HELIUM, qubzStateDS, consoleLogger, fileLogger)

		case 3:
			elementName = "Ne"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_NEON, qubzStateDS, consoleLogger, fileLogger)

		case 4:
			elementName = "Ar"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_ARGON, qubzStateDS, consoleLogger, fileLogger)

		case 5:
			elementName = "Kr"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_KRYPTON, qubzStateDS, consoleLogger, fileLogger)

		case 6:
			elementName = "Xe"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_XENON, qubzStateDS, consoleLogger, fileLogger)

		case 7:
			elementName = "Rn"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_RADON, qubzStateDS, consoleLogger, fileLogger)

		case 8:
			elementName = "Og"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_OGANESSON, qubzStateDS, consoleLogger, fileLogger)

		case 9:
			elementName = "CO2"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_CARBON_DIOXIDE, qubzStateDS, consoleLogger, fileLogger)

		case 10:
			elementName = "CO"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_CARBON_MONOXIDE, qubzStateDS, consoleLogger, fileLogger)

		case 11:
			elementName = "C3H8"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_PROPANE, qubzStateDS, consoleLogger, fileLogger)

		case 12:
			elementName = "C2H2"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_ACETYLENE, qubzStateDS, consoleLogger, fileLogger)

		case 13:
			elementName = "H"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_HYDROGEN, qubzStateDS, consoleLogger, fileLogger)

		case 14:
			elementName = "N"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_NITROGEN, qubzStateDS, consoleLogger, fileLogger)

		case 15:
			elementName = "O"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_OXYGEN, qubzStateDS, consoleLogger, fileLogger)

		case 16:
			elementName = "CnH2N+2"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_PETROLEUM, qubzStateDS, consoleLogger, fileLogger)

		case 17:
			elementName = "CH4"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_METHANE, qubzStateDS, consoleLogger, fileLogger)

		case 18:
			elementName = "C2H6"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_ETHANE, qubzStateDS, consoleLogger, fileLogger)

		case 19:
			elementName = "C4H10"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_BUTANE, qubzStateDS, consoleLogger, fileLogger)

		case 20:
			elementName = "C5H12"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_PENTANE, qubzStateDS, consoleLogger, fileLogger)

		case 21:
			elementName = "C6H14"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_HEXANE, qubzStateDS, consoleLogger, fileLogger)

		case 22:
			elementName = "C"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_CARBON, qubzStateDS, consoleLogger, fileLogger)

		case 23:
			elementName = "S"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_SULFUR, qubzStateDS, consoleLogger, fileLogger)

		case 24:
			elementName = "KNO3"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_POTASSIUM_NITRATE, qubzStateDS, consoleLogger, fileLogger)

		case 25:
			elementName = "U"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_URANIUM, qubzStateDS, consoleLogger, fileLogger)

		case 26:
			elementName = "Pu"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_PLUTONIUM, qubzStateDS, consoleLogger, fileLogger)

		case 27:
			elementName = "Be"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_BERYLLIUM, qubzStateDS, consoleLogger, fileLogger)

		case 28:
			elementName = "B"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_BORON, qubzStateDS, consoleLogger, fileLogger)

		case 29:
			elementName = "F"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_FLUORINE, qubzStateDS, consoleLogger, fileLogger)

		case 30:
			elementName = "Na"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_SODIUM, qubzStateDS, consoleLogger, fileLogger)

		case 31:
			elementName = "Mg"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_MAGNESIUM, qubzStateDS, consoleLogger, fileLogger)

		case 32:
			elementName = "Al"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_ALUMINUM, qubzStateDS, consoleLogger, fileLogger)

		case 33:
			elementName = "Si"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_SILICON, qubzStateDS, consoleLogger, fileLogger)

		case 34:
			elementName = "P"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_PHOSPHORUS, qubzStateDS, consoleLogger, fileLogger)

		case 35:
			elementName = "Cl"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_CHLORINE, qubzStateDS, consoleLogger, fileLogger)

		case 36:
			elementName = "K"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_POTASSIUM, qubzStateDS, consoleLogger, fileLogger)

		case 37:
			elementName = "Ca"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_CALCIUM, qubzStateDS, consoleLogger, fileLogger)

		case 38:
			elementName = "Sc"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_SCANDIUM, qubzStateDS, consoleLogger, fileLogger)

		case 39:
			elementName = "Ti"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_TITANIUM, qubzStateDS, consoleLogger, fileLogger)

		case 40:
			elementName = "V"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_VANADIUM, qubzStateDS, consoleLogger, fileLogger)

		case 41:
			elementName = "Cr"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_CHROMIUM, qubzStateDS, consoleLogger, fileLogger)

		case 42:
			elementName = "Mn"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_MANGANESE, qubzStateDS, consoleLogger, fileLogger)

		case 43:
			elementName = "Fe"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_IRON, qubzStateDS, consoleLogger, fileLogger)

		case 44:
			elementName = "Co"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_COBALT, qubzStateDS, consoleLogger, fileLogger)

		case 45:
			elementName = "Ni"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_NICKEL, qubzStateDS, consoleLogger, fileLogger)

		case 46:
			elementName = "Cu"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_COPPER, qubzStateDS, consoleLogger, fileLogger)

		case 47:
			elementName = "Zn"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_ZINC, qubzStateDS, consoleLogger, fileLogger)

		case 48:
			elementName = "Ga"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_GALLIUM, qubzStateDS, consoleLogger, fileLogger)

		case 49:
			elementName = "Ge"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_GERMANIUM, qubzStateDS, consoleLogger, fileLogger)

		case 50:
			elementName = "As"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_ARSENIC, qubzStateDS, consoleLogger, fileLogger)

		case 51:
			elementName = "Se"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_SELENIUM, qubzStateDS, consoleLogger, fileLogger)

		case 52:
			elementName = "Br"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_BROMINE, qubzStateDS, consoleLogger, fileLogger)

		case 53:
			elementName = "Rb"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_RUBIDIUM, qubzStateDS, consoleLogger, fileLogger)

		case 54:
			elementName = "Sr"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_STRONTIUM, qubzStateDS, consoleLogger, fileLogger)

		case 55:
			elementName = "Y"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_YTTRIUM, qubzStateDS, consoleLogger, fileLogger)

		case 56:
			elementName = "Pd"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_PALLADIUM, qubzStateDS, consoleLogger, fileLogger)

		case 57:
			elementName = "Ag"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_SILVER, qubzStateDS, consoleLogger, fileLogger)

		case 58:
			elementName = "Pb"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_LEAD, qubzStateDS, consoleLogger, fileLogger)

		case 59:
			elementName = "Th"
			partsPerMillion = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SPECTROMETER_THORIUM, qubzStateDS, consoleLogger, fileLogger)
		}

		(*qubzMatrix)[matrixIndex].Spectrometer.Elements[elindex].ElementName = elementName
		(*qubzMatrix)[matrixIndex].Spectrometer.Elements[elindex].PartsPerMillion = partsPerMillion
	}

}
