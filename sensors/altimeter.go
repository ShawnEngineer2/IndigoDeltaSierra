package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func AltimeterInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Altimeter sensor

	var altitudeValue float64 = 0

	//Generate an appropriate value based on the passed transport mode
	switch (*qubzMatrix)[matrixIndex].TransportMode {
	case appconstants.TRANSPORT_MODE_DRAGON:
		altitudeValue = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_ALTIMETER_ALTITUDE_DRAGON, qubzStateDS, consoleLogger, fileLogger)
	case appconstants.TRANSPORT_MODE_FALCON:
		altitudeValue = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_ALTIMETER_ALTITUDE_FALCON, qubzStateDS, consoleLogger, fileLogger)
	case appconstants.TRANSPORT_MODE_PLANE:
		altitudeValue = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_ALTIMETER_ALTITUDE_PLANE, qubzStateDS, consoleLogger, fileLogger)
	case appconstants.TRANSPORT_MODE_SHIP:
		altitudeValue = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_ALTIMETER_ALTITUDE_SHIP, qubzStateDS, consoleLogger, fileLogger)
	case appconstants.TRANSPORT_MODE_TRAIN:
		altitudeValue = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_ALTIMETER_ALTITUDE_TRAIN, qubzStateDS, consoleLogger, fileLogger)
	case appconstants.TRANSPORT_MODE_TRUCK:
		altitudeValue = datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_ALTIMETER_ALTITUDE_TRUCK, qubzStateDS, consoleLogger, fileLogger)
	default:
		altitudeValue = 0
	}

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].Altimeter.AltimeterState = 1
	(*qubzMatrix)[matrixIndex].Altimeter.Altitude = altitudeValue
	(*qubzMatrix)[matrixIndex].Altimeter.EventState = appconstants.SENSOR_STATE_CURRENT

}
