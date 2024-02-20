package datautil

import (
	"indigodeltasierra/appconstants"
)

func GetSensorIDForSensorDesc(sensorDesc string) int {
	//This routine is the cross reference to get Sensor ID for a Sensor description
	var sensorID int = 0

	switch sensorDesc {
	case appconstants.SENSOR_TYPE_ALTIMETER:
		sensorID = appconstants.SENSOR_TYPE_ID_ALTIMETER

	case appconstants.SENSOR_TYPE_BATTERY:
		sensorID = appconstants.SENSOR_TYPE_ID_BATTERY

	case appconstants.SENSOR_TYPE_COMPUTE:
		sensorID = appconstants.SENSOR_TYPE_ID_COMPUTE

	case appconstants.SENSOR_TYPE_FIRE:
		sensorID = appconstants.SENSOR_TYPE_ID_FIRE

	case appconstants.SENSOR_TYPE_GEIGER:
		sensorID = appconstants.SENSOR_TYPE_ID_GEIGER

	case appconstants.SENSOR_TYPE_GPS:
		sensorID = appconstants.SENSOR_TYPE_ID_GPS

	case appconstants.SENSOR_TYPE_GYRO:
		sensorID = appconstants.SENSOR_TYPE_ID_GYRO

	case appconstants.SENSOR_TYPE_LOCK:
		sensorID = appconstants.SENSOR_TYPE_ID_LOCK

	case appconstants.SENSOR_TYPE_MOTION:
		sensorID = appconstants.SENSOR_TYPE_ID_MOTION

	case appconstants.SENSOR_TYPE_RADIO:
		sensorID = appconstants.SENSOR_TYPE_ID_RADIO

	case appconstants.SENSOR_TYPE_SEAL:
		sensorID = appconstants.SENSOR_TYPE_ID_SEAL

	case appconstants.SENSOR_TYPE_SPECTROMETER:
		sensorID = appconstants.SENSOR_TYPE_ID_SPECTROMETER

	case appconstants.SENSOR_TYPE_TEMPHUMIDITY:
		sensorID = appconstants.SENSOR_TYPE_ID_TEMPHUMIDITY
	}

	return sensorID
}

func GetSensorDescForSensorID(sensorID int) string {
	//This routine is the cross reference to get Sensor ID for a Sensor description
	var sensorDesc string = ""

	switch sensorID {
	case appconstants.SENSOR_TYPE_ID_ALTIMETER:
		sensorDesc = appconstants.SENSOR_TYPE_ALTIMETER

	case appconstants.SENSOR_TYPE_ID_BATTERY:
		sensorDesc = appconstants.SENSOR_TYPE_BATTERY

	case appconstants.SENSOR_TYPE_ID_COMPUTE:
		sensorDesc = appconstants.SENSOR_TYPE_COMPUTE

	case appconstants.SENSOR_TYPE_ID_FIRE:
		sensorDesc = appconstants.SENSOR_TYPE_FIRE

	case appconstants.SENSOR_TYPE_ID_GEIGER:
		sensorDesc = appconstants.SENSOR_TYPE_GEIGER

	case appconstants.SENSOR_TYPE_ID_GPS:
		sensorDesc = appconstants.SENSOR_TYPE_GPS

	case appconstants.SENSOR_TYPE_ID_GYRO:
		sensorDesc = appconstants.SENSOR_TYPE_GYRO

	case appconstants.SENSOR_TYPE_ID_LOCK:
		sensorDesc = appconstants.SENSOR_TYPE_LOCK

	case appconstants.SENSOR_TYPE_ID_MOTION:
		sensorDesc = appconstants.SENSOR_TYPE_MOTION

	case appconstants.SENSOR_TYPE_ID_RADIO:
		sensorDesc = appconstants.SENSOR_TYPE_RADIO

	case appconstants.SENSOR_TYPE_ID_SEAL:
		sensorDesc = appconstants.SENSOR_TYPE_SEAL

	case appconstants.SENSOR_TYPE_ID_SPECTROMETER:
		sensorDesc = appconstants.SENSOR_TYPE_SPECTROMETER

	case appconstants.SENSOR_TYPE_ID_TEMPHUMIDITY:
		sensorDesc = appconstants.SENSOR_TYPE_TEMPHUMIDITY
	}

	return sensorDesc
}
