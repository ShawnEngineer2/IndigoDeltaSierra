package datamodels

// This struct represents a single motion sensor reading
type MotionReading struct {
	EventState        int     `json:"EventState"`
	MotionSensorState int     `json:"MotionSensorState"`
	NumberOfContacts  int     `json:"NumberOfContacts"`
	AverageVelocity   float64 `json:"AverageVelocity"`
}

// This struct represents a single motion event
type MotionEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []MotionReading `json:"SensorData"`
}
