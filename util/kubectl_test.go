package util

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type KubectlTestSuite struct {
	suite.Suite
}

func (suite *KubectlTestSuite) SetupTest() {
	stderr, err := Create("../testing/artifacts/mysqlcrd.yaml")
	suite.Require().Nil(err, stderr.String())
}

func (suite *KubectlTestSuite) TearDownTest() {
	stderr, err := Delete("../testing/artifacts/mysqlcrd.yaml")
	suite.Require().Nil(err, stderr.String())

	stderr, err = Delete("../testing/artifacts/mycluster.yaml")
}

func (suite *KubectlTestSuite) TestKubectlApply() {
	stderr, err := Apply("../testing/artifacts/mycluster.yaml")
	suite.Nil(err, stderr.String())

	stderr, err = Apply("../testing/artifacts/mycluster.yaml")
	suite.Require().NotNil(err, "Should be an error: already exists")
	suite.Require().Contains(stderr.String(), "already exists", "Should be an error: already exists, found: "+stderr.String())
}

func (suite *KubectlTestSuite) TestKubectlCreate() {
	stderr, err := Create("../testing/artifacts/mysqlcrd.yaml")
	suite.Require().NotNil(err, "Should be an error: CRD already exists")
	suite.Require().Contains(stderr.String(), "already exists", "Should be an error: CRD already exists, found: "+stderr.String())
}

func (suite *KubectlTestSuite) TestKubectlDelete() {
	stderr, err := Delete("../testing/artifacts/mysqlcrd.yaml")
	suite.Require().Nil(err, stderr.String())

	stderr, err = Delete("../testing/artifacts/mysqlcrd.yaml")
	suite.Require().NotNil(err, "Should be an error: CRD not found")
	suite.Require().Contains(stderr.String(), "not found", "Should be an error: CRD not found, found: "+stderr.String())

	stderr, err = Create("../testing/artifacts/mysqlcrd.yaml")
	suite.Require().Nil(err, stderr.String())

	stderr, err = Apply("../testing/artifacts/mycluster.yaml")
	suite.Require().Nil(err, stderr.String())

	stderr, err = Delete("../testing/artifacts/mycluster.yaml")
	suite.Require().Nil(err, stderr.String())
}

func TestClusterControllerTestSuite(t *testing.T) {
	suite.Run(t, new(KubectlTestSuite))
}
