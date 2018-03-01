package controller

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ControllerBaseHooksTestSuite struct {
	suite.Suite
	controller Base
}

func (suite *ControllerBaseHooksTestSuite) SetupTest() {
	suite.controller = NewControllerBase()
}

func (suite *ControllerBaseHooksTestSuite) TestControllerBase_AddHook() {
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

func (suite *ControllerBaseHooksTestSuite) TestControllerBase_RemoveHook() {
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

func (suite *ControllerBaseHooksTestSuite) TestControllerBase_GetHooks() {
	hook := NewEventsHook(1)
	anotherHook := NewEventsHook(1)
	err := suite.controller.AddHook(hook)
	suite.Assert().Nil(err)
	err = suite.controller.AddHook(anotherHook)
	suite.Assert().Nil(err)
	suite.Require().Equal(2, len(suite.controller.hooks))

	hooks := suite.controller.GetHooks()
	suite.Require().Equal(2, len(hooks))
	suite.Assert().Contains(hooks, hook)
	suite.Assert().Contains(hooks, anotherHook)
	suite.Require().Equal(hooks, suite.controller.hooks)
}

func TestControllerBaseHooksTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerBaseHooksTestSuite))
}
