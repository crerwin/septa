package septa

type TransitViewResult struct {
	Bus []struct {
		Trip        string `json:"trip"`
		Lat         string `json:"lat"`
		Lng         string `json:"lng"`
		Label       string `json:"label"`
		VehicleID   string `json:"VehicleID"`
		BlockID     string `json:"BlockID"`
		Direction   string `json:"Direction"`
		Destination string `json:"destination"`
		Offset      string `json:"Offset"`
		OffsetSec   string `json:"Offset_sec"`
	} `json:"bus"`
}
