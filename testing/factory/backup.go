package factory

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/nauyey/factory/def"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// MySQLBackupFactory generates cluster with random data for testing.
var MySQLBackupFactory = def.NewFactory(crv1.MySQLBackup{}, "",
	def.DynamicField("ObjectMeta.Name", func(model interface{}) (interface{}, error) {
		return fmt.Sprintf("backup-%s", randomdata.RandStringRunes(16)), nil
	}),
	def.Field("ObjectMeta.Namespace", "default"),
	def.DynamicField("Spec.Time", func(model interface{}) (interface{}, error) {
		minute := randomAny(randomdata.Number(0, 59))
		hour := randomAny(randomdata.Number(0, 23))
		day := randomAny(randomdata.Number(1, 31))
		month := randomAny(randomdata.Number(1, 12))
		weekday := randomAny(randomdata.Number(1, 7))
		year := randomAny(randomdata.Number(1900, 3000))
		return fmt.Sprintf("%s %s %s %s %s %s", minute, hour, day, month, weekday, year), nil
	}),
	def.Trait("ChangeDefaults",
		def.Field("Spec.Storage", resource.MustParse("1Gi")),
	),
)

func randomAny(value int) string {
	if randomdata.Boolean() {
		return "*"
	}
	return fmt.Sprintf("%d", value)
}
