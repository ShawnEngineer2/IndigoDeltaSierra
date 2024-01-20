package datamodels

//This struct is used to hold the list of Qubz Names and IDs from the Qubz.dat file

type Qubz struct {
	QubzID   int    `json:"QubzId"`
	QubzName string `json:"QubzName"`
}
