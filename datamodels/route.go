package datamodels

// This struct holds LeMonde Route records
type Route struct {
	RouteID              int     `json:"RouteID"`
	OriginCode           string  `json:"OriginCode"`
	DestinationCode      string  `json:"DestinationCode"`
	ClassOfServiceID     int     `json:"ClassOfService"`
	TransportModeID      int     `json:"TransportModeID"`
	CostInDollars        float32 `json:"CostInDollars"`
	NumDaysTravel        int     `json:"NumDaysTravel"`
	OriginLatitude       float32 `json:"OriginLatitude"`
	OriginLongitude      float32 `json:"OriginLongitude"`
	DestinationLatitude  float32 `json:"DestinationLatitude"`
	DestinationLongitude float32 `json:"DestinationLongitude"`
	BearingToDestination float32 `json:"BearingToDestination"`
}
