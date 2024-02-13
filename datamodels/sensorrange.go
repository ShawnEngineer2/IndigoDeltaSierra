package datamodels

// This struct holds a sensor range record
type SensorRange struct {
	SensorTypeId      int     `json:"SensorTypeId"`
	SensorType        string  `json:"SensorType"`
	SensorDataPoint   string  `json:"SensorDataPoint"`
	SensorDataPointId int     `json:"SensorDataPointId"`
	TransportMode     int     `json:"TransportMode"`
	DataPointTypeId   int     `json:"DataPointTypeId"`
	NumberScale       int     `json:"NumberScale"`
	NominalMin        float64 `json:"NominalMin"`
	NominalMax        float64 `json:"NominalMax"`
	ExceptionMin      float64 `json:"ExceptionMin"`
	ExceptionMax      float64 `json:"ExceptionMax"`
}
