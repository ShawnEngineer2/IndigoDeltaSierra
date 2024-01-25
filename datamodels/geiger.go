package datamodels

// This struct represents a single geiger counter sensor reading
type GeigerReading struct {
	EventState         int `json:"State"`
	GeigerCounterState int `json:"GeigerCounterState"`
	GeigerReading      int `json:"GeigerReading"`
}

// This struct represents a single geiger counter event
type GeigerEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []GeigerReading `json:"SensorData"`
}
