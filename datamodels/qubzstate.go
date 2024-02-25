package datamodels

// This struct is used to hold the state of a single sensor data point
type QubzSensorDataPoint struct {
	SensorDataPointId int
	DataPointTypeId   int
	DataPointValue    float64
}

// This struct is used to hold the sensor and exception state of a single Qubz unit
type QubzState struct {
	QubzID                    int
	QubzName                  string
	RouteAssignment           int
	ShipmentType              int
	TransportMode             int
	ExceptionAssignment       int
	ExceptionSeverity         int
	ExceptionType             int
	ExceptionIntervalBoundary int
	CurrentExceptionInterval  int
	SensorDataPoints          []QubzSensorDataPoint
}
