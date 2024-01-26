package datamodels

// This struct represents a single motion sensor reading
type SealReading struct {
	EventState      int    `json:"EventState"`
	SealSensorState int    `json:"SealSensorState"`
	SealState       int    `json:"SealState"`
	SealEventTime   string `json:"SealEventTime"`
}

// This struct represents a single motion event
type SealEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []MotionReading `json:"SensorData"`
}
