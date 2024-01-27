package datamodels

// This struct represents a single element entry
type ElementReading struct {
	ElementName     string  `json:"ElementName"`
	PartsPerMillion float64 `json:"PartsPerMillion"`
}

// This struct represents a single spectrometer sensor reading
type SpectrometerReading struct {
	EventState                  int              `json:"EventState"`
	SpectrometerState           int              `json:"SpectrometerState"`
	Elements                    []ElementReading `json:"Elements"`
	Opocs                       float64          `json:"Opocs"`
	Explosives                  float64          `json:"Explosives"`
	Urates                      float64          `json:"Urates"`
	WeaponsGradeNuclearMaterial float64          `json:"Wgnm"`
}

// This struct represents a single spectrometer event
type SpectrometerEvent struct {
	EventHeader QubzEventHeader       `json:"Header"`
	SensorData  []SpectrometerReading `json:"SensorData"`
}
