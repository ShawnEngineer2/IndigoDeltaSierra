package datamodels

// This struct represents a single lock sensor reading
type LockReading struct {
	EventState      int    `json:"EventState"`
	LockState       int    `json:"LockState"`
	LockEventTime   string `json:"LockEventTime"`
	LockEventType   int    `json:"LockEventType"`   //1=Lock / 2=Unlock
	LockEventMethod int    `json:"LockEventMethod"` //1=Remote Signal / 2=Local Biometrics / 3=HSM Dongle
}

// This struct represents a single lock event
type LockEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []LockReading   `json:"SensorData"`
}
