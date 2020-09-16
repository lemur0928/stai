package cbe

// CBE CBEFrame
type CBE interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp() int
	GetMicrosecond() int64
	GetObjNumber() int

	GetObjByGlobalID(globalID string) CBEObjStruct
	GetObjByTrackingID(trackingID int) CBEObjStruct
	GetObjByIndex(index int) CBEObjStruct
	GetTObjects() []CBEObjStruct

	SetFrameID(frameID string)
	SetSensorID(sensorID string)

	SetTimeStamp(timestamp int)
	SetMicrosecond(timestampMicrosecond int64)

	SetTObjects(tObjects []CBEObjStruct)
}
