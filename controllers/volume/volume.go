package volume

import (
	"begoceph/controllers"
)

type VolumeController struct {
	controllers.BaseController
}

// @Title format volume filesystem
// @Description format volume filesystem (ext4,xfs)
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404 volume not found
// @router /:volume/format [get]
func (v *VolumeController) Format() {

}

// @Title volume resize
// @Description format volume filesystem (ext4, )
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404 volume not found
// @router /:volume/resize [get]
func (v *VolumeController) Resize() {

}
