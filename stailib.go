package stailib

import (
	"github.com/lemur0928/stai/cbe"
	"github.com/bigobject-inc/stailib/tt"
)

// STAILibrary STAI library
type STAILibrary interface {
	GetServTrackTrace() tt.TrackTrace
	GetServCBE() cbe.CBEServ
}
