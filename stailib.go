package stailib

import (
	"github.com/bigobject-inc/stailib/tt"
)

// STAILibrary STAI library
type STAILibrary interface {
	GetServTrackTrace() tt.TrackTrace
}
