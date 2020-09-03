package stailib

import (
	"github.com/ptoyoohri/stailib/cbe"
	"github.com/ptoyoohri/stailib/tt"
)

// STAILibrary STAI library
type STAILibrary interface {
	GetServTrackTrace() tt.TrackTrace
	GetServCBE() cbe.CBEServ
	//	GetServCBEObjStruct() cbe.CBEObjStruct
}
