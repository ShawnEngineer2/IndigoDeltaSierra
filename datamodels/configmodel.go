package datamodels

//This struct is used to hold configuration values loaded at startup

type Config struct {
	QubzCount       int    `json:"qubzCount"`
	EventInterval   int    `json:"eventInterval"`
	EventCycleCount int    `json:"eventCycleCount"`
	EmailAddress    string `json:"emailAddress"`
	QueueEndpoint   string `json:"queueEndpoint"`
	QueueTopic      string `json:"queueTopic"`
	LogLocation     string `json:"logLocation"`
}
