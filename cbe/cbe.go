package cbe

type CBE interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp() float64
	GetObjNumber() int

	GetObjByGlobalID(GlobalID string) CBEObjStruct
	GetObjByTrackingID(trackingID int) CBEObjStruct
	GetObjByIndex(index int) CBEObjStruct

	SetFrameID(frameID string)
	SetSensorID(sensorID string)
	SetTimeStampFloat(timestampFloat float64)
	SetTimeStamp(timestamp int)
	SetDateTime(dateTime string)
	SetTObjects(tObjects []CBEObjStruct)

	SetFrameInfo()
	SetGlobalIDIndex()
	SetTrackingIDIndex()
}
