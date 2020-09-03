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
	GetTag(tagName string) interface{}

	SetGlobalID(globalID string)
	SetFrameInfo(frameInfo CBE)
	//	Distance(CBEObjStruct) int
	//	Overlap(CBEObjStruct) bool
	IsName(name string) bool
	Carry(obj CBEObjStruct) bool
}
