package util

import (
	"context"
	"encoding/json"
	"tit/database"

	"tit/model"

	"github.com/carmel/gooxml/spreadsheet"
	"github.com/xshrim/gol"
	"github.com/xshrim/gol/tk"
	//	"go.mongodb.org/mongo-driver/bson"
)

// func convertProject(fpath string) {
// 	var projs []interface{}
// 	for line := range tk.IterFile(fpath) {
// 		data := strings.Split(line, ",")
// 		proj := model.Project{
// 			Pid:        data[2],
// 			Name:       data[0],
// 			Code:       data[1],
// 			Dept:       "综合业务开发部",
// 			Group:      data[3],
// 			Business:   "",
// 			Level:      data[4],
// 			User:       data[5],
// 			Risk:       data[9],
// 			Note:       data[10],
// 			Developer:  data[6],
// 			Maintainer: data[7],
// 			Operator:   data[8],
// 		}
// 		projs = append(projs, proj)

// 	}

// 	// mgo := database.New("tit", "project")
// 	// result, err := mgo.InsertMany(projs)
// 	// gol.Prtln(result, err)
// 	jsb, _ := json.Marshal(projs)
// 	gol.Prtln(string(jsb))

// }

// func convertRvitem(fpath string) {
// 	var rvs []interface{}
// 	for line := range tk.IterFile(fpath) {
// 		data := strings.Split(line, ",")
// 		gol.Prtln(line, len(data))
// 		rv := model.Rvitem{
// 			Rid:      data[0],
// 			Category: data[1],
// 			Subject:  data[2],
// 			Require:  data[3],
// 			Level:    data[4],
// 			Content:  data[5],
// 			Result:   data[6],
// 			Note:     data[7],
// 		}
// 		rvs = append(rvs, rv)
// 	}

// 	// mgo := database.New("tit", "project")
// 	// result, err := mgo.InsertMany(projs)
// 	// gol.Prtln(result, err)
// 	rvb, _ := json.Marshal(rvs)
// 	gol.Prtln(string(rvb))

// }

func syncProject() {
	lst := make(map[string][2]string)
	data := tk.ReadFile("../sys.json")
	items := tk.Jsquery(string(data), ".data.list").([]any)
	for _, item := range items {
		it := item.(map[string]any)
		pid := it["systemNo"].(string)
		name := it["systemName"].(string)
		biz := it["businessDepart"].(string)
		lst[pid] = [2]string{name, biz}
		//gol.Prtln(pid, name, biz)
	}

	ss, _ := spreadsheet.Open("../proj.xlsx")
	st := ss.Sheets()[0]
	for idx, row := range st.Rows() {
		if idx == 0 {
			continue
		}
		pid := row.Cell("C").GetString()
		pname := row.Cell("A").GetString()
		name := lst[pid][0]
		biz := lst[pid][1]
		cell := row.Cell("L")
		cell.SetString(biz)

		gol.Prtln(pid, pname, name, biz)
	}

	ss.SaveToFile("../project.xlsx")
}

func project2json(fpath string) {
	var projs []model.Project
	ss, _ := spreadsheet.Open(fpath)
	st := ss.Sheets()[0]
	for idx, row := range st.Rows() {
		if idx == 0 {
			continue
		}
		name := row.Cell("A").GetString()
		code := row.Cell("B").GetString()
		pid := row.Cell("C").GetString()
		dept := row.Cell("D").GetString()
		group := row.Cell("E").GetString()
		business := row.Cell("F").GetString()
		level := row.Cell("G").GetString()
		user := row.Cell("H").GetString()
		developer := row.Cell("I").GetString()
		maintainer := row.Cell("J").GetString()
		operator := row.Cell("K").GetString()
		risk := row.Cell("L").GetString()
		note := row.Cell("M").GetString()

		proj := model.Project{
			Pid:        pid,
			Name:       name,
			Code:       code,
			Dept:       dept,
			Group:      group,
			Business:   business,
			Level:      level,
			User:       user,
			Risk:       risk,
			Developer:  developer,
			Maintainer: maintainer,
			Operator:   operator,
			Note:       note,
		}
		projs = append(projs, proj)
	}

	projb, _ := json.Marshal(projs)

	tk.WriteFile("../project.json", projb)
}

func impProject(fpath string) {
	var projects []model.Project
	data := tk.ReadFile(fpath)
	if err := json.Unmarshal(data, &projects); err != nil {
		gol.Prtln(err)
	}

	gol.Prtln(projects)
	mgo := database.New("tit", "project")

	tmp := make([]interface{}, len(projects))
	for i, v := range projects {
		tmp[i] = v
	}
	result, err := mgo.InsertMany(tmp)
	gol.Prtln(result, err)
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
	js["uid"] = "chenpan1"

	// var bs bson.M
	// result := mgo.FindOne("id", "chenpan1")
	// result.Decode(&bs)
	// gol.Prtln(bs)

	data, _ := mgo.FindMany(js)

	var users []model.User
	data.All(context.TODO(), &users)
	gol.Prtln(users)
}

func rvitem2json(fpath string) {
	var rvs []model.Rvitem
	ss, _ := spreadsheet.Open(fpath)
	st := ss.Sheets()[0]
	for idx, row := range st.Rows() {
		if idx == 0 {
			continue
		}
		rid := row.Cell("A").GetString()
		category := row.Cell("B").GetString()
		subject := row.Cell("C").GetString()
		require := row.Cell("D").GetString()
		level := row.Cell("E").GetString()
		content := "不涉及"
		tips := row.Cell("F").GetString()
		result := row.Cell("G").GetString()
		note := row.Cell("H").GetString()

		rv := model.Rvitem{
			Rid:      rid,
			Category: category,
			Subject:  subject,
			Require:  require,
			Level:    level,
			Content:  content,
			Tips:     tips,
			Result:   result,
			Note:     note,
		}
		rvs = append(rvs, rv)
	}

	rvb, _ := json.Marshal(rvs)

	tk.WriteFile("../rvitem.json", rvb)
}

func impRvitem(fpath string) {
	var rvs []model.Rvitem
	data := tk.ReadFile(fpath)
	if err := json.Unmarshal(data, &rvs); err != nil {
		gol.Prtln(err)
	}

	gol.Prtln(rvs)
	mgo := database.New("tit", "rvitem")

	tmp := make([]interface{}, len(rvs))
	for i, v := range rvs {
		tmp[i] = v
	}
	result, err := mgo.InsertMany(tmp)
	gol.Prtln(result, err)
}
