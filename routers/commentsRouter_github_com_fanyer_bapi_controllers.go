package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:StudyController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:StudyController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:SubmissionController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:SubmissionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/fanyer/bapi/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
