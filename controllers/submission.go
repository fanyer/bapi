package controllers

import (
	"github.com/fanyer/bapi/models"

	"github.com/astaxie/beego"
)

// Operations about study
type SubmissionController struct {
	beego.Controller
}

// @Title Get
// @Description find study by studyid
// @Param	studyId		path 	string	true		"the studyid you want to get"
// @Success 200 {study} models.Study
// @Failure 403 :studyId is empty
// @router /:studyId [get]
// func (o *StudyController) Get() {
// 	studyId := o.Ctx.Input.Param(":studyId")
// 	if studyId != "" {
// 		ob, err := models.GetOne(studyId)
// 		if err != nil {
// 			o.Data["json"] = err.Error()
// 		} else {
// 			o.Data["json"] = ob
// 		}
// 	}
// 	o.ServeJSON()
// }

// @Title GetAll
// @Description get all studys
// @Success 200 {study} models.Study
// @Failure 403 :studyId is empty
// @router / [get]
func (o *SubmissionController) GetAll() {
	obs := models.GetAllSubmission()
	o.Data["json"] = obs
	o.ServeJSON()
}
