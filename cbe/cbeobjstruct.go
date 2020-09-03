package cbe 

type CBEObjStruct interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp()	float64

	GetGlobalID() string
	GetTrackingID() string
	GetName()	string
	GetProbability()	float64
	GetConfidence()		float64
	GetBox()	BoundingBox
	GetLocation()	GeoCoordinate
	GetTag(string)	interface{}

	SetGlobalID(string)
	
	Distance(CBEObjStruct) int
	Overlap(CBEObjStruct) bool
	IsName(string) bool
	Carry(CBEObjStruct) bool
}
