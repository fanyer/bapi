package models

import (
	"fmt"
	"time"

	"github.com/fanyer/bapi/models/db"

	"gopkg.in/mgo.v2/bson"
)

type Study struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"` // Mongodb _id
	Title         string        `json:"Title,omitempty" bson:"Title,omitempty"`
	UniversalID   string        `json:"UniversalID,omitempty" bson:"UniversalID,omitempty"`
	Abstract      string        `json:"Abstract,omitempty" bson:"Abstract,omitempty"`
	StudyType     string        `json:"StudyType,omitempty" bson:"StudyType,omitempty"`
	Description   string        `json:"Description,omitempty" bson:"Description,omitempty"`
	CenterName    string        `json:"CenterName,omitempty" bson:"CenterName,omitempty"`
	Organism      string        `json:"Organism,omitempty" bson:"Organism,omitempty"`
	Owner         string        `json:"Owner,omitempty" bson:"Owner,omitempty"`
	DeleteFlag    bool          `json:"DeleteFlag,omitempty" bson:"DeleteFlag,omitempty"`
	CreateTime    time.Time     `json:"CreateTime,omitempty" bson:"CreateTime,omitempty"`
	CreateTimeRaw string        `json:"CreateTimeRaw,omitempty" bson:"CreateTimeRaw,omitempty"`
}

var result []Study

func init() {

	mConn := db.Conn()

	c := mConn.DB("knowledgebase").C("study")

	c.Find(nil).All(&result)

	fmt.Println("******************")
	// fmt.Println(result)
	fmt.Println("******************")

}

// func GetOne(StudyId string) (study *Study, err error) {
// 	if v, ok := result[StudyId]; ok {
// 		return v, nil
// 	}
// 	return nil, errors.New("StudyId Not Exist")
// }

func GetAllStudy() []Study {
	return result
}
