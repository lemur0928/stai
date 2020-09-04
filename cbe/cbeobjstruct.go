package cbe

type CBEObjStruct interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp() float64

	GetGlobalID() string
	GetTrackingID() int
	GetName() string
	GetProbability() float64
	GetConfidence() float64
	GetBox() BoundingBox
	GetLocation() Geo
	GetTag(key string) interface{}

	SetFrameInfo(frameInfo CBE)
	SetGlobalID(globalID string)
	SetTrackingID(trackingID int)
	SetName(name string)
	SetProbability(probability float64)
	SetConfidnece(confidence float64)
	SetBox(box BoundingBox)
	SetLocation(geo Geo)
	SetTag(key string, value interface{})

	Distance(CBEObjStruct) float64
	Overlap(CBEObjStruct) bool
	IsName(name string) bool
	Carry(obj CBEObjStruct) bool
}
