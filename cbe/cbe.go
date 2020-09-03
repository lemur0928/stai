package cbe


type CBE interface {
	GetFrameID() string
	GetSensorID() string
	GetTimeStamp()	float64
	GetObjNumber() int
	GetAllObjects() []TObject
	GetObjByGlobalID(string) *cbe.TObject
	GetObjByTrackingID(int)	*cbe.TObject
	GetObjByIndex(int)	*cbe.TObject
	
	SetFrameInfo()
	SetGlobalIDIndex()
	SetTrackingIDIndex()	
}

