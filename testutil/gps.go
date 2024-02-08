package testutil

import (
	"indigodeltasierra/customerror"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/eventemitter"
)

func GPSTestEvent(outputFilePath string) {
	//This routine is used to test the GPS JSON

	jsondata := datamodels.GPSEvent{}

	jsondata.EventHeader.EventTimestamp = "2024-01-26T23:38:14.123Z"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2
	jsondata.EventHeader.TransportMode = 6
	jsondata.EventHeader.SensorType.SensorTypeId = 5
	jsondata.EventHeader.SensorType.SensorTypeDescription = "gps"

	expectedLatitude := datamodels.GPSCoordinate{
		Hours:   134,
		Minutes: 18,
		Seconds: 32,
	}

	expectedLongitude := datamodels.GPSCoordinate{
		Hours:   -15,
		Minutes: -45,
		Seconds: -2,
	}

	expectedPostion := datamodels.GPSPosition{
		Latitude:  expectedLatitude,
		Longitude: expectedLongitude,
	}

	actualLatitude := datamodels.GPSCoordinate{
		Hours:   134,
		Minutes: 18,
		Seconds: 19,
	}

	actualLongitude := datamodels.GPSCoordinate{
		Hours:   -15,
		Minutes: -41,
		Seconds: -57,
	}

	actualPosition := datamodels.GPSPosition{
		Latitude:  actualLatitude,
		Longitude: actualLongitude,
	}

	jsondata.SensorData = make([]datamodels.GPSReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].ExpectedGPSPosition = expectedPostion
	jsondata.SensorData[0].ActualGPSPosition = actualPosition

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].ExpectedGPSPosition = expectedPostion
	jsondata.SensorData[1].ActualGPSPosition = actualPosition

	err := eventemitter.EventToFile(jsondata, outputFilePath)

	customerror.CheckAndPanic(err)
}
