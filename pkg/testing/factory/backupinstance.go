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
	def.Field("Status.Phase", crv1.MySQLBackupScheduled),
	def.Field("ObjectMeta.Namespace", "default"),
	def.AfterBuild(func(model interface{}) error {
		instance, ok := model.(*crv1.MySQLBackupInstance)
		if !ok {
			return fmt.Errorf("invalid type of model in ObjectMeta.Name function")
		}

		// Generate name in After build to avoid flaky tests when schedule is not yet generated.
		minute := randomdata.Number(0, 59)
		hour := randomdata.Number(0, 23)
		day := randomdata.Number(1, 31)
		month := randomdata.Number(1, 12)
		year := randomdata.Number(1900, 3000)
		date := fmt.Sprintf("%d-%02d-%02d-%02d-%02d", year, month, day, hour, minute)

		instance.Name = fmt.Sprintf("%s-%s", instance.Spec.Schedule, date)
		return nil
	}),
)
