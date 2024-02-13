package datamodels

// This struct is used to hold the state of a single sensor data point
type QubzSensorDataPoint struct {
	SensorDataPointId int
	DataPointTypeId   int
	DataPointValue    float64
}

// This struct is used to hold the sensor and exception state of a single Qubz unit
type QubzState struct {
	QubzID                int
	QubzName              string
	RouteAssignment       int
	ShipmentType          int
	TransportMode         int
	ExceptionAssignment   int
	ExceptionSeverity     int
	ExceptionIntermittent bool
	ExceptionInterval     int //Note: Only applies if Exception is intermittent. If Exception Interval is -1, the error occurs randomly (0 /1 type stuff)
	SensorDataPoints      []QubzSensorDataPoint
}
