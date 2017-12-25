package factory

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/nauyey/factory/def"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

var MySQLClusterFactory = def.NewFactory(crv1.MySQLCluster{}, "",
	def.DynamicField("ObjectMeta.Name", func(model interface{}) (interface{}, error) {
		return fmt.Sprintf("cluster-%s", randomdata.RandStringRunes(16)), nil
	}),
	def.Field("ObjectMeta.Namespace", "default"),
	def.DynamicField("Spec.Name", func(model interface{}) (interface{}, error) {
		cluster := model.(*crv1.MySQLCluster)
		return fmt.Sprintf("%s-name", cluster.Name), nil
	}),
	def.DynamicField("Spec.Port", func(model interface{}) (interface{}, error) {
		return fmt.Sprintf("%d", randomdata.Number(1024, 49151)), nil
	}),
	def.DynamicField("Spec.User", func(model interface{}) (interface{}, error) {
		return randomdata.GenerateProfile(randomdata.RandomGender).Login.Username, nil
	}),
	def.DynamicField("Spec.Password", func(model interface{}) (interface{}, error) {
		return randomdata.GenerateProfile(randomdata.RandomGender).Login.Password, nil
	}),
	def.DynamicField("Spec.Database", func(model interface{}) (interface{}, error) {
		cluster := model.(*crv1.MySQLCluster)
		return fmt.Sprintf("%s-db", cluster.Spec.User), nil
	}),
)
