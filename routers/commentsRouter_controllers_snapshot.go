package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"] = append(beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"] = append(beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"] = append(beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"],
		beego.ControllerComments{
			Method: "Clone",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"] = append(beego.GlobalControllerRouter["begoceph/controllers/snapshot:SnapshotController"],
		beego.ControllerComments{
			Method: "RollBack",
			Router: `/rollback`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
