package util

import "testing"

// func TestConvertProject(t *testing.T) {
// 	convertProject("../project.csv")
// }

// func TestConvertRvitem(t *testing.T) {
// 	convertRvitem("../rvform.csv")
// }

func TestSyncProject(t *testing.T) {
	syncProject()
}

func TestProject2json(t *testing.T) {
	project2json("../project.xlsx")
}

func TestImpProject(t *testing.T) {
	impProject("../project.json")
}

func TestImpUser(t *testing.T) {
	impUser("../user.json")
}

func TestFindUser(t *testing.T) {
	findUser()
}

func TestRvitem2json(t *testing.T) {
	rvitem2json("../rvitem.xlsx")
}

func TestImpRvitem(t *testing.T) {
	impRvitem("../rvitem.json")
}
