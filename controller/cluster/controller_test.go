package cluster

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	kubeFake "k8s.io/client-go/kubernetes/fake"
)

type ClusterControllerHooksTestSuite struct {
	suite.Suite
	controller *clusterController
}

func (suite *ClusterControllerHooksTestSuite) SetupTest() {
	kubeClientset := kubeFake.NewSimpleClientset()
	clientset := fake.NewSimpleClientset()
	controller := NewClusterController(clientset, kubeClientset)

	var ok bool
	suite.controller, ok = controller.(*clusterController)
	suite.Require().True(ok)
}

func (suite *ClusterControllerHooksTestSuite) TestAddHook() {
	suite.Require().Equal(0, len(suite.controller.hooks))

	// Add hook
	hook := NewEventsHook(1) // Any hook will do
	err := suite.controller.AddHook(hook)
	suite.Assert().Nil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Hook already exists
	err = suite.controller.AddHook(hook)
	suite.Assert().NotNil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Add another hook
	anotherHook := NewEventsHook(1) // Any hook will do
	err = suite.controller.AddHook(anotherHook)
	suite.Assert().Nil(err)
	suite.Require().Equal(2, len(suite.controller.hooks))
}

func (suite *ClusterControllerHooksTestSuite) TestRemoveHook() {
	hook := NewEventsHook(1)
	anotherHook := NewEventsHook(1)
	err := suite.controller.AddHook(hook)
	suite.Assert().Nil(err)
	err = suite.controller.AddHook(anotherHook)
	suite.Assert().Nil(err)
	suite.Require().Equal(2, len(suite.controller.hooks))

	// Remove hook
	err = suite.controller.RemoveHook(hook)
	suite.Assert().Nil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Try to remove the same hook again
	err = suite.controller.RemoveHook(hook)
	suite.Assert().NotNil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Remove another hook
	err = suite.controller.RemoveHook(anotherHook)
	suite.Assert().Nil(err)
	suite.Require().Equal(0, len(suite.controller.hooks))
}

func TestClusterControllerHooksTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterControllerHooksTestSuite))
}
