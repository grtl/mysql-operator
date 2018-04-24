package controller

import "errors"

// Base implements basic functions for controllers.
type Base struct {
	hooks []Hook
}

// NewControllerBase returns new Base.
func NewControllerBase() Base {
	return Base{
		hooks: []Hook{},
	}
}

// GetHooks returns hooks registered in the controller
func (c *Base) GetHooks() []Hook {
	return c.hooks
}

// AddHook adds a new hook to the controller.
func (c *Base) AddHook(hook Hook) error {
	for _, h := range c.hooks {
		if h == hook {
			return errors.New("Given hook is already installed in the current controller")
		}
	}
	c.hooks = append(c.hooks, hook)
	return nil
}

// RemoveHook removes given hook from the controller.
func (c *Base) RemoveHook(hook Hook) error {
	for i, h := range c.hooks {
		if h == hook {
			// Removing hooks is not that common so we can afford it in O(n)
			c.hooks = append(c.hooks[:i], c.hooks[i+1:]...)
			return nil
		}
	}
	return errors.New("Given hook is not installed in the current controller")
}
