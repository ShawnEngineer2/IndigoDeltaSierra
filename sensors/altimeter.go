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

func AltimeterSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine sets values for the Altimeter sensor

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
	(*qubzMatrix)[matrixIndex].Altimeter.AltimeterState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_ALTIMETER_ALTIMETER_STATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].Altimeter.Altitude = altitudeValue
	(*qubzMatrix)[matrixIndex].Altimeter.EventState = appconstants.SENSOR_STATE_CURRENT

}

func AltimeterGetEvent(qubzMatrixCurrent *[]datamodels.QubzMatrix, qubzMatrixPrevious *[]datamodels.QubzMatrix, matrixIndex int, eventHeader *datamodels.QubzEventHeader, consoleLogger *slog.Logger, fileLogger *slog.Logger) datamodels.AltimeterEvent {

	//Create Event Instance
	eventInstance := datamodels.AltimeterEvent{}

	//Fill in Header
	eventInstance.EventHeader.EventTimestamp = datautil.GetRFC3339TimeString()
	eventInstance.EventHeader.QubzId = eventHeader.QubzId
	eventInstance.EventHeader.RouteAssignment = eventHeader.RouteAssignment
	eventInstance.EventHeader.ShipmentType = eventHeader.ShipmentType
	eventInstance.EventHeader.SensorType.SensorTypeId = appconstants.SENSOR_TYPE_ID_ALTIMETER
	eventInstance.EventHeader.SensorType.SensorTypeDescription = appconstants.SENSOR_TYPE_ALTIMETER
	eventInstance.EventHeader.EventUUID = datautil.GetUUID()

	//Initialize Sensor Data
	eventInstance.SensorData = make([]datamodels.AltimeterReading, 2)

	//Set current State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].AltimeterState = (*qubzMatrixCurrent)[matrixIndex].Altimeter.AltimeterState
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].Altitude = (*qubzMatrixCurrent)[matrixIndex].Altimeter.Altitude
	eventInstance.SensorData[appconstants.SENSOR_STATE_CURRENT].EventState = (*qubzMatrixCurrent)[matrixIndex].Altimeter.EventState

	//Set previous State data
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].AltimeterState = (*qubzMatrixPrevious)[matrixIndex].Altimeter.AltimeterState
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].Altitude = (*qubzMatrixPrevious)[matrixIndex].Altimeter.Altitude
	eventInstance.SensorData[appconstants.SENSOR_STATE_PREVIOUS].EventState = (*qubzMatrixPrevious)[matrixIndex].Altimeter.EventState

	//Return the completed event
	return eventInstance

}
