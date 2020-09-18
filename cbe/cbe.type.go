package cbe

// Location latitude, longitude, geohash
type Location struct {
	Latitude  float64 `json:"latitude"`  // $$ 緯度
	Longitude float64 `json:"longitude"` // $$ 經度
	Geohash   string  `json:"geohash"`
}

// BoundingBox upper-left point, width, height
type BoundingBox struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
