package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["begoceph/controllers/volume:VolumeController"] = append(beego.GlobalControllerRouter["begoceph/controllers/volume:VolumeController"],
		beego.ControllerComments{
			Method: "Format",
			Router: `/:volume/format`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["begoceph/controllers/volume:VolumeController"] = append(beego.GlobalControllerRouter["begoceph/controllers/volume:VolumeController"],
		beego.ControllerComments{
			Method: "Resize",
			Router: `/:volume/resize`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
