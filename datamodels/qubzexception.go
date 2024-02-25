package datamodels

// This struct holds data for a Qubz sensor exception
type QubzSensorDataPointException struct {
	SensorDataPointId   int    `json:"sensorDataPointId"`
	SensorDataPointDesc string `json:"sensorDataPointDesc"`
	ValueModType        int    `json:"valueModType"`
	FixedModValue       int    `json:"fixedModValue"`
}

// This struct represents a single Qubz exception made up of one or more sensor data exceptions
type QubzException struct {
	ExceptionId           int                            `json:"exceptionId"`
	ExceptionDesc         string                         `json:"exceptionDesc"`
	SeverityLevel         int                            `json:"severityLevel"`
	ExceptionType         int                            `json:"exceptionType"`
	IntermittencyInterval int                            `json:"intermittencyInterval"`
	AffectedSensors       []QubzSensorDataPointException `json:"affectedSensors"`
}
