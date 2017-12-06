// snapshot controller
package snapshot

import (
	"begoceph/controllers"
)

type SnapshotController struct {
	controllers.BaseController
}

// @Title create snapshot for volume
// @Description create snapshot for volume
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404
// @router / [post]
func (sc *SnapshotController) Create() {

}

// @Title  delete snapshot for volume
// @Description delete snapshot for volume by snapshot name
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404
// @router / [delete]
func (sc *SnapshotController) Delete() {

}

// @Title rollback snapshot for volume
// @Description rollback snapshot for volume
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404
// @router /rollback [post]
func (sc *SnapshotController) RollBack() {

}

// @Title  purge snapshot for volume
// @Description purge snapshot for volume
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404
// @router /purge [post]

func (sc *SnapshotController) Purge() {

}

// @Title clone snapshot for new volume
// @Description clone snapshot for new volume
// @Param   key     path    string  true        "description"
// @Success 200
// @Failure 400
// @Failure 404
// @router / [post]
func (sc *SnapshotController) Clone() {

}
