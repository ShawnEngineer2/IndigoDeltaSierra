package datamodels

// This struct holds LeMonde location records
type Location struct {
	LocationID   int     `json:"LocationID"`
	LocationCode string  `json:"LocationCode"`
	Description  string  `json:"Description"`
	City         string  `json:"City"`
	State        string  `json:"State"`
	Country      string  `json:"Country"`
	Latitude     float32 `json:"Latitude"`
	Longitude    float32 `json:"Longitude"`
}
