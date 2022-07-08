package util

import "testing"

func TestConvertProject(t *testing.T) {
	impProject("../reboot.csv")
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
