package cbe

type CBE interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp() int
	GetTimeStampFloat() float64
	GetDateTime() string
	GetObjNumber() int

	GetObjByGlobalID(GlobalID string) CBEObjStruct
	GetObjByTrackingID(trackingID int) CBEObjStruct
	GetObjByIndex(index int) CBEObjStruct
	GetTObjects() []CBEObjStruct

	SetFrameID(frameID string)
	SetSensorID(sensorID string)
	SetTimeStampFloat(timestampFloat float64)
	SetTimeStamp(timestamp int)
	SetDateTime(timestampFloat float64)
	SetTObjects(tObjects []CBEObjStruct)

	SetFrameInfo()
	SetGlobalIDIndex()
	SetTrackingIDIndex()
}
