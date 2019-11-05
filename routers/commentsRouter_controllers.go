package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["server/controllers:UsersController"] = append(beego.GlobalControllerRouter["server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UsersController"] = append(beego.GlobalControllerRouter["server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UsersController"] = append(beego.GlobalControllerRouter["server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UsersController"] = append(beego.GlobalControllerRouter["server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/getAll`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UsersController"] = append(beego.GlobalControllerRouter["server/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
