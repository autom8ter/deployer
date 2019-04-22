package deployer_test

import (
	"github.com/autom8ter/deployer"
	"github.com/autom8ter/objectify"
	"testing"
)

var util = objectify.Default()

func TestNewDeployer(t *testing.T) {
	d, err := deployer.NewDeployer(false)
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := util.Validate(d); err != nil {
		t.Fatal(err.Error())
	}
}
