package util

import (
	"context"
	"encoding/json"
	"strings"
	"titserver/database"

	"titserver/model"

	"github.com/xshrim/gol"
	"github.com/xshrim/gol/tk"
	//	"go.mongodb.org/mongo-driver/bson"
)

func impProject(fpath string) {
	var projs []interface{}
	for line := range tk.IterFile(fpath) {
		data := strings.Split(line, ",")
		proj := model.Project{
			ID:         data[2],
			Name:       data[0],
			Code:       data[1],
			Dept:       "综合业务开发部",
			Group:      data[3],
			Level:      data[4],
			User:       data[5],
			Risk:       data[9],
			Note:       data[10],
			Developer:  data[6],
			Maintainer: data[7],
			Operator:   data[8],
		}
		projs = append(projs, proj)

	}

	mgo := database.New("tit", "project")
	result, err := mgo.InsertMany(projs)
	gol.Prtln(result, err)
	// jsb, _ := json.Marshal(projs)
	// gol.Prtln(string(jsb))

}

func impUser(fpath string) {
	var users []model.User
	data := tk.ReadFile(fpath)
	if err := json.Unmarshal(data, &users); err != nil {
		gol.Prtln(err)
	}

	gol.Prtln(users)
	mgo := database.New("tit", "user")

	tmp := make([]interface{}, len(users))
	for i, v := range users {
		tmp[i] = v
	}
	result, err := mgo.InsertMany(tmp)
	gol.Prtln(result, err)
}

func findUser() {
	mgo := database.New("tit", "user")
	js := make(map[string]any)
	js["dept"] = "综合业务开发部"
	js["id"] = "chenpan1"

	// var bs bson.M
	// result := mgo.FindOne("id", "chenpan1")
	// result.Decode(&bs)
	// gol.Prtln(bs)

	data, _ := mgo.FindMany(js)

	var users []model.User
	data.All(context.TODO(), &users)
	gol.Prtln(users)
}
