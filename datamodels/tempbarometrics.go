package datamodels

// This struct represents a single Temperature / Barometer sensor reading
type TempBarometricsReading struct {
	EventState                 int `json:"EventState"`
	OverallInternalTemperature int `json:"OverallInternalTemperature"`
	TemperatureSensorState     int `json:"TemperatureSensorState"`
	BarometricSensorState      int `json:"BarometricSensorState"`
	HumidityLevel              int `json:"HumidityLevel"`
}

// This struct represents a single Temperature / Barometer event
type TempBarometricsEvent struct {
	EventHeader QubzEventHeader          `json:"Header"`
	SensorData  []TempBarometricsReading `json:"SensorData"`
}
