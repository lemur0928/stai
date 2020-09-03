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
	GetTag(string) interface{}

	SetGlobalID(string)
	SetFrameInfo(CBE)
	//	Distance(CBEObjStruct) int
	//	Overlap(CBEObjStruct) bool
	IsName(string) bool
	//	Carry(CBEObjStruct) bool
}
