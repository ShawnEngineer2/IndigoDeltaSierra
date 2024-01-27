package datamodels

//This struct is used to hold the list of Qubz Names and IDs from the Qubz.dat file

type Qubz struct {
	QubzID   int    `json:"QubzId"`
	QubzName string `json:"QubzName"`
}

//This struct is used as an identifier header for published Qubz events

type QubzEventHeader struct {
	QubzId          string     `json:"QubzId"`
	EventTimestamp  string     `json:"EventTimestamp"`
	RouteAssignment int        `json:"RouteAssignment"`
	ShipmentType    int        `json:"ShipmentType"`
	SensorType      QubzSensor `json:"SensorType"`
}

//This struc is used to identify sensor types and Ids

type QubzSensor struct {
	SensorTypeId          int    `json:"SensorTypeId"`
	SensorTypeDescription string `json:"SensorTypeDescription"`
}
