package tt

import (
	"github.com/ptoyoohri/stailib/cbe"
)

// FilterCommand filter Command
type FilterCommand struct {
	Action string   `json:"action"`
	IDs    []string `json:"ids"`
}

// Node node info of a path
type Node struct {
	FrameID              string
	SensorID             string
	TimestampMicrosecond int64
	Timestamp            int
	GlobalID             string
	TrackingID           int
	Name                 string
	SequenceID           string
	BackLink             string
	Probability          float64
	Confidence           float64
	GeoLocation          cbe.Location
	BoundingBox          cbe.BoundingBox
}
