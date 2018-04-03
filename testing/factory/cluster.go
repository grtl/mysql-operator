package factory

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/nauyey/factory/def"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// MySQLClusterFactory generates cluster with random data for testing.
var MySQLClusterFactory = def.NewFactory(crv1.MySQLCluster{}, "",
	def.DynamicField("ObjectMeta.Name", func(model interface{}) (interface{}, error) {
		return fmt.Sprintf("cluster-%s", randomdata.RandStringRunes(16)), nil
	}),
	def.Field("ObjectMeta.Namespace", "default"),
	def.DynamicField("Spec.Secret", func(model interface{}) (interface{}, error) {
		cluster, ok := model.(*crv1.MySQLCluster)
		if !ok {
			return nil, fmt.Errorf("invalid type of model in Spec.Secret function")
		}
		return fmt.Sprintf("%s-secret", cluster.Name), nil
	}),
	def.Field("Spec.Storage", resource.MustParse("1Gi")),
	def.Trait("ChangeDefaults",
		def.DynamicField("Spec.Replicas", func(model interface{}) (interface{}, error) {
			return uint32(randomdata.Number(3, 1<<8)), nil
		}),
		def.DynamicField("Spec.Port", func(model interface{}) (interface{}, error) {
			return int32(randomdata.Number(1<<12, 1<<16)), nil
		}),
		def.DynamicField("Spec.Image", func(model interface{}) (interface{}, error) {
			major := randomdata.Number(1, 10)
			minor := randomdata.Number(1, 10)
			return fmt.Sprintf("mysql:v%d.%d", major, minor), nil
		}),
	),
)
