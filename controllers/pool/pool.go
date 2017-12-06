package pool

// get ceph pool status infomation
import (
	"begoceph/controllers"
)

type PoolController struct {
	controllers.BaseController
}

// @Title get rbd pool status
// @Description get rbd pool status
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404
// @router /status [get]
func (p *PoolController) Status() {

}
