package cbe

// CBEObjStruct :
type CBEObjStruct interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp() float64

	GetGlobalID() string
	GetBackLink() string
	GetTrackingID() int
	GetName() string
	GetProbability() float64
	GetConfidence() float64
	GetBox() BoundingBox
	GetLocation() Location
	GetTag(key string) interface{}

	SetFrameInfo(frameInfo CBE)
	SetGlobalID(globalID string)
	SetBackLink(backLink string)
	SetTrackingID(trackingID int)
	SetName(name string)
	SetProbability(probability float64)
	SetConfidnece(confidence float64)
	SetBox(box BoundingBox)
	SetLocation(geo Location)
	SetTag(key string, value interface{})

	Distance(CBEObjStruct) float64
	Overlap(CBEObjStruct) bool
	IsName(name string) bool
	Carry(obj CBEObjStruct) bool
}
