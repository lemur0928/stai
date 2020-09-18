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
	FrameID              string          `json:"frameID"`
	SensorID             string          `json:"sensorID"`
	TimestampMicrosecond int64           `json:"timestampMicrosecond"`
	Timestamp            int             `json:"timestamp"`
	GlobalID             string          `json:"globalID"`
	TrackingID           int             `json:"trackingID"`
	Name                 string          `json:"name"`
	SequenceID           string          `json:"sequenceID"`
	BackLink             string          `json:"backLink"`
	Probability          float64         `json:"probability"`
	Confidence           float64         `json:"confidence"`
	GeoLocation          cbe.Location    `json:"location"`
	BoundingBox          cbe.BoundingBox `json:"boundingBox"`
}
