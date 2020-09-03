package cbe


type CBE interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp()	float64
	GetObjNumber() int
	GetAllObjects() []CBEObjStruct
	GetObjByGlobalID(string) *CBEObjStruct
	GetObjByTrackingID(int)	*CBEObjStruct
	GetObjByIndex(int)	*CBEObjStruct
	
	SetFrameInfo()
	SetGlobalIDIndex()
	SetTrackingIDIndex()	
}

