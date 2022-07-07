package util

import "testing"

func TestImpProject(t *testing.T) {
	impProject("../reboot.csv")
}

func TestImpUser(t *testing.T) {
	impUser("../user.json")
}

func TestFindUser(t *testing.T) {
	findUser()
}
