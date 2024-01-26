package datamodels

// This struct represents a single gyroscope sensor reading
type GyroscopeReading struct {
	EventState            int `json:"EventState"`
	XAxisAngle            int `json:"XAxisAngle"`
	YAxisAngle            int `json:"YAxisAngle"`
	ZAxisAngle            int `json:"ZAxisAngle"`
	GyroscopicSensorState int `json:"GeigerReading"`
}

// This struct represents a single gyroscope event
type GyroscopeEvent struct {
	EventHeader QubzEventHeader    `json:"Header"`
	SensorData  []GyroscopeReading `json:"SensorData"`
}
