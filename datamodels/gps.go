package datamodels

// This struct represents a single GPS Coordinate
type GPSCoordinate struct {
	Hours   int `json:"Hours"`
	Minutes int `json:"Minutes"`
	Seconds int `json:"Seconds"`
}

// This struct represents a single GPS Position
type GPSPosition struct {
	Latitude  GPSCoordinate `json:"Latitude"`
	Longitude GPSCoordinate `json:"Longitude"`
}

// This struct represents a single GPS sensor reading
type GPSReading struct {
	EventState          int         `json:"EventState"`
	ExpectedGPSPosition GPSPosition `json:"ExpectedGPSPosition"`
	ActualGPSPosition   GPSPosition `json:"ActualGPSPosition"`
	GPSSensorState      int         `json:"GPSSensorState"`
}

// This struct represents a single compute event
type GPSEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []GPSReading    `json:"SensorData"`
}
