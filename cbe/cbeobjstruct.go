package cbe

// CBEObjStruct :
type CBEObjStruct interface {
	// get info from frame
	GetFrameID() string
	GetSensorID() string
	GetTimestamp() int

	// get object info
	GetGlobalID() string
	GetBackLink() string
	GetSequenceID() string
	GetTrackingID() int
	GetName() string
	GetProbability() float64
	GetConfidence() float64
	GetBox() BoundingBox
	GetLocation() Location
	GetTag(label, name string) (v interface{})

	// set object info
	SetGlobalID(globalID string)
	SetBackLink(backLink string)
	SetTrackingID(trackingID int)
	SetName(name string)
	SetProbability(probability float64)
	SetConfidnece(confidence float64)
	SetBox(box BoundingBox)
	SetLocation(geo Location)
	SetTag(label, name string, value interface{})

	SetFrameInfo(CBE)
	Distance(CBEObjStruct) float64
	Overlap(CBEObjStruct) bool
	IsName(name string) bool
	Carry(obj CBEObjStruct) bool
}
