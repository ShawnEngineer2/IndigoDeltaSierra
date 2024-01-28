package datamodels

// This struct represents a single battery sensor reading
type BatteryReading struct {
	EventState                   int     `json:"EventState"`
	MaxAmpHours                  int     `json:"MaxAmpHours"`
	RemainingAmpHours            int     `json:"RemainingAmpHours"`
	DrainRate                    int     `json:"DrainRate"`
	CapacitanceGelExpectedVolume int     `json:"CapacitanceGelExpectedVolume"`
	CapacitanceGelActualVolume   int     `json:"CapacitanceGelActualVolume"`
	BatteryTemperature           float64 `json:"BatteryTemperature"`
	BatterySensorState           int     `json:"BatterySensorState"`
	BatteryChargePct             float64 `json:"BatteryChargePct"`
}

// This struct represents a single battery sensor event
type BatteryEvent struct {
	EventHeader QubzEventHeader  `json:"Header"`
	SensorData  []BatteryReading `json:"SensorData"`
}
