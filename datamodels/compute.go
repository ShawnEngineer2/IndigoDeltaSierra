package datamodels

// This struct represents a single compute sensor reading
type ComputeReading struct {
	EventState             int     `json:"EventState"`
	CPUUtilization         float64 `json:"CPUUtilization"`
	MemoryUtilization      float64 `json:"MemoryUtilization"`
	NumCPUCores            int     `json:"NumCPUCores"`
	TotalAmountOfMemory    int     `json:"TotalAmountOfMemory"`
	AmountMemoryUtilized   int     `json:"AmountMemoryUtilized"`
	TotalDiskStorage       int     `json:"TotalDiskStorage"`
	RemainingDiskStorage   int     `json:"RemainingDiskStorage"`
	ComputeSensorState     int     `json:"ComputeSensorState"`
	ComputeFirmwareState   int     `json:"ComputeFirmwareState"`
	ComputeFirmwareVersion int     `json:"ComputeFirmwareVersion"`
	OSTypeVersion          string  `json:"OSTypeVersion"`
}

// This struct represents a single compute event
type ComputeEvent struct {
	EventHeader QubzEventHeader  `json:"Header"`
	SensorData  []ComputeReading `json:"SensorData"`
}
