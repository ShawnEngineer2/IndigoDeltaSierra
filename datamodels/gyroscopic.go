package datamodels

// This struct represents a single gyroscope sensor reading
type GyroscopeReading struct {
	EventState            int     `json:"EventState"`
	XAxisAngle            float64 `json:"XAxisAngle"`
	YAxisAngle            float64 `json:"YAxisAngle"`
	ZAxisAngle            float64 `json:"ZAxisAngle"`
	GyroscopicSensorState int     `json:"GyroSensorState"`
}

// This struct represents a single gyroscope event
type GyroscopeEvent struct {
	EventHeader QubzEventHeader    `json:"Header"`
	SensorData  []GyroscopeReading `json:"SensorData"`
}
