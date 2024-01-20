package datamodels

//This struct is used to implement the Qubz Matrix - this is the centerpiece of data tracking in this app

type QubzMatrix struct {
	QubzID   int    `json:"QubzId"`
	QubzName string `json:"QubzName"`
}
