package tt

import (
	"github.com/bigobject-inc/stailib/tt/space"
	"github.com/bigobject-inc/stailib/tt/space/cctv"
)

// TrackTrace track and trace service
type TrackTrace interface {
	GetServGeometry() space.Geometry
	NewCCTV(id string, geolocation [2]float64, resolution []interface{}, cameraMatrix [][]float64, rmat [][]float64, tvec [][]float64, dist [][]float64) cctv.CCTV
}
