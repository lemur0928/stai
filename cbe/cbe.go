package cbe

type CBE interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp() float64
	GetObjNumber() int

	GetObjByGlobalID(GlobalID string) CBEObjStruct
	GetObjByTrackingID(trackingID int) CBEObjStruct
	GetObjByIndex(index int) CBEObjStruct

	SetFrameInfo()
	SetGlobalIDIndex()
	SetTrackingIDIndex()
}
