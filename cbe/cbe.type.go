package cbe

type Location struct {
	Latitude  float64 // $$ 緯度
	Longitude float64 // $$ 經度
	Geohash   string
}

type BoundingBox struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
