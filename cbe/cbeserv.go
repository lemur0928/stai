package cbe

// CBEServ : CBE service
type CBEServ interface {
	NewCBEFrame() CBE
	NewCBEObject() CBEObjStruct
}
