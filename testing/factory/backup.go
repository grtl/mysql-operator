package factory

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/nauyey/factory/def"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// MySQLBackupFactory generates cluster with random data for testing.
var MySQLBackupFactory = def.NewFactory(crv1.MySQLBackup{}, "",
	def.DynamicField("ObjectMeta.Name", func(model interface{}) (interface{}, error) {
		return fmt.Sprintf("backup-%s", randomdata.RandStringRunes(16)), nil
	}),
	def.Field("ObjectMeta.Namespace", "default"),
)
