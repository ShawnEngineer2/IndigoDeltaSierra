package datamodels

// This struct represents a single altimeter sensor reading
type FireReading struct {
	EventState           int `json:"EventState"`
	FireSensorState      int `json:"FireSensorState"`
	FireDetected         int `json:"FireDetected"`
	FireSuppressionState int `json:"FireSuppressionState"`
}

// This struct represents a single altimeter sensor event
type FireEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []FireReading   `json:"SensorData"`
}
