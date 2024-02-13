package datamodels

//This struct is used to hold configuration values loaded at startup

type Config struct {
	QubzCount           int    `json:"qubzCount"`
	InternalRandomizer  int    `json:"useInternalRandomizer"`
	EventInterval       int    `json:"eventInterval"`
	EventCycleCount     int    `json:"eventCycleCount"`
	EmailAddress        string `json:"emailAddress"`
	QueueEndpoint       string `json:"queueEndpoint"`
	QueueTopic          string `json:"queueTopic"`
	LogLocation         string `json:"logLocation"`
	QubzNameFile        string `json:"qubzNameFile"`
	LocationsFile       string `json:"locationsFile"`
	RoutesFile          string `json:"routesFile"`
	ClassOfServiceFile  string `json:"classOfServiceFile"`
	SensorTypesFile     string `json:"sensorTypesFile"`
	ShipmentTypesFile   string `json:"shipmentTypesFile"`
	TransportModesFile  string `json:"transportModesFile"`
	SensorRangesFile    string `json:"sensorRangesFile"`
	OutputChannel       string `json:"outputChannel"`
	EventOutputLocation string `json:"eventOutputLocation"`
}
