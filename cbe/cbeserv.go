package cbe

//	"github.com/ptoyoohri/stailib/tt/space"
//	"github.com/ptoyoohri/stailib/tt/space/cctv"

// TrackTrace track and trace service
type CBEServ interface {
	NewCBEFrame() CBE
	NewCBEObject() CBEObjStruct
}
