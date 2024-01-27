package datamodels

// This struct represents a single geiger counter sensor reading
type GeigerReading struct {
	EventState         int     `json:"EventState"`
	GeigerCounterState int     `json:"GeigerCounterState"`
	GeigerReading      float64 `json:"GeigerReading"`
}

// This struct represents a single geiger counter event
type GeigerEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []GeigerReading `json:"SensorData"`
}
