package factory

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/nauyey/factory/def"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// MySQLBackupInstanceFactory generates a Backup instance schedule with random data for testing.
var MySQLBackupInstanceFactory = def.NewFactory(crv1.MySQLBackupInstance{}, "",
	def.DynamicField("Spec.Schedule", func(model interface{}) (interface{}, error) {
		return fmt.Sprintf("backup-%s", randomdata.RandStringRunes(16)), nil
	}),
	def.DynamicField("ObjectMeta.Name", func(model interface{}) (interface{}, error) {
		instance, ok := model.(*crv1.MySQLBackupInstance)
		if !ok {
			return nil, fmt.Errorf("invalid type of model in ObjectMeta.Name function")
		}

		minute := randomdata.Number(0, 59)
		hour := randomdata.Number(0, 23)
		day := randomdata.Number(1, 31)
		month := randomdata.Number(1, 12)
		year := randomdata.Number(1900, 3000)
		date := fmt.Sprintf("%d-%d-%d-%d-%d", minute, hour, day, month, year)

		return fmt.Sprintf("%s-%s", instance.Spec.Schedule, date), nil
	}),
	def.Field("ObjectMeta.Namespace", "default"),
)
