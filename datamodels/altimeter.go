package datamodels

// This struct represents a single altimeter sensor reading
type AltimeterReading struct {
	EventState     int `json:"EventState"`
	AltimeterState int `json:"AltimeterState"`
	Altitude       int `json:"Altitude"`
}

// This struct represents a single altimeter sensor event
type AltimeterEvent struct {
	EventHeader QubzEventHeader    `json:"Header"`
	SensorData  []AltimeterReading `json:"SensorData"`
}
