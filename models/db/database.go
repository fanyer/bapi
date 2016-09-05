package db

import (
	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func Conn() *mgo.Session {
	return session.Copy()
}

/*
func Close() {
	session.Close()
}
*/

func init() {
	url := beego.AppConfig.String("mongodb::url")

	sess, err := mgo.Dial(url)
	fmt.Println("DB first")
	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
}
