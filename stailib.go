package stailib

import (
	"github.com/ptoyoohri/stailib/cbe"
	"github.com/ptoyoohri/stailib/tt"
)

// STAILibrary STAI library
type STAILibrary interface {
	GetServTrackTrace() tt.TrackTrace
	GetServCBE() cbe.CBE
	GetServCBEObjStruct() cbe.CBEObjStruct
}
