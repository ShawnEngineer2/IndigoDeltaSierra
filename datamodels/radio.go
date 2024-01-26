package datamodels

// This struct represents a single radio sensor reading
type RadioReading struct {
	EventState              int    `json:"EventState"`
	RadioSensorState        int    `json:"RadioSensorState"`
	RadioPowerState         int    `json:"RadioPowerState"`
	RadioFirmwareState      int    `json:"RadioFirmwareState"`
	RadioFirmwareVersion    string `json:"RadioFirmwareVersion"`
	RadioSignalStrength     int    `json:"RadioSignalStrength"`
	GPSSignalStrength       int    `json:"GPSSignalStrength"`
	LocalWiFiSignalStrength int    `json:"LocalWiFiSignalStrength"`
	GPSHandshakeStable      int    `json:"GPSHandshakeStable"`
	LocalWiFiHandshake      int    `json:"LocalWiFiHandshake"`
	WiFiUplinkState         int    `json:"WiFiUplinkState"`
	WiFiDownlinkState       int    `json:"WiFiDownlinkState"`
	WiFiDownloadSpeed       int    `json:"WiFiDownloadSpeed"`
	WiFiUploadSpeed         int    `json:"WiFiUploadSpeed"`
}

// This struct represents a single radio event
type RadioEvent struct {
	EventHeader QubzEventHeader `json:"Header"`
	SensorData  []RadioReading  `json:"SensorData"`
}
