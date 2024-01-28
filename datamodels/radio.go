package datamodels

// This struct represents a single radio sensor reading
type RadioReading struct {
	EventState              int     `json:"EventState"`
	RadioSensorState        int     `json:"RadioSensorState"`
	RadioPowerState         int     `json:"RadioPowerState"`
	RadioFirmwareState      int     `json:"RadioFirmwareState"`
	RadioFirmwareVersion    string  `json:"RadioFirmwareVersion"`
	RadioSignalStrength     float64 `json:"RadioSignalStrength"`
	GPSSignalStrength       float64 `json:"GPSSignalStrength"`
	LocalWiFiSignalStrength float64 `json:"LocalWiFiSignalStrength"`
	GPSHandshakeStable      int     `json:"GPSHandshakeStable"`
	LocalWiFiHandshake      int     `json:"LocalWiFiHandshake"`
	WiFiUplinkState         int     `json:"WiFiUplinkState"`
	WiFiDownlinkState       int     `json:"WiFiDownlinkState"`
	WiFiDownloadSpeed       float64 `json:"WiFiDownloadSpeed"`
	WiFiUploadSpeed         float64 `json:"WiFiUploadSpeed"`
}

// This struct represents a single radio event
type RadioEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []RadioReading  `json:"SensorData"`
}
