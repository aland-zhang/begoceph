// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com

package routers

import (
	"begoceph/controllers"
	"begoceph/controllers/snapshot"
	"begoceph/controllers/volume"

	"github.com/astaxie/beego"
)

func init() {
	beego.NewNamespace(controllers.ApiPrefix,
		beego.NSNamespace("/volume",
			beego.NSInclude(&volume.VolumeController{})),
		beego.NSNamespace("/Volume/:volume/snapshot",
			beego.NSInclude(&snapshot.SnapshotController{}),
		),
	)
}
