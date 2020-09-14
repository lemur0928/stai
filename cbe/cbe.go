package cbe

type CBE interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp() int64

	GetObjNumber() int

	GetObjByGlobalID(globalID string) CBEObjStruct
	GetObjByTrackingID(trackingID int) CBEObjStruct
	GetObjByIndex(index int) CBEObjStruct
	GetTObjects() []CBEObjStruct

	SetFrameID(frameID string)
	SetSensorID(sensorID string)

	SetTimeStamp(timestamp int64)

	SetTObjects(tObjects []CBEObjStruct)
}
