package datautil

import (
	"errors"
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

//This package contains routines for retrieving sensor range information

func GetSensorRangeInfo(sensorRangeData *[]datamodels.SensorRange, dataPointId int) (datamodels.SensorRange, error) {
	//Scan the passed Sensor Range data and retrieve the sought after data point
	for _, x := range *sensorRangeData {

		if x.SensorDataPointId == dataPointId {
			//Return the found data point
			return x, nil
		}

	}

	//If you get this far you didn't find it - return an empty data point and an error
	emptyDataPoint := datamodels.SensorRange{}
	return emptyDataPoint, errors.New(appconstants.ITEM_NOT_FOUND)

}
