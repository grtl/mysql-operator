package controller

import "context"

// Controller is base type for all custom resource controllers.
type Controller interface {
	// Run starts the controller listening on custom resource changes.
	Run(ctx context.Context) error
	// AddHook allows to inject new hook into the controller.
	AddHook(hook Hook) error
	// RemoveHook removes hook from the controller.
	RemoveHook(hook Hook) error
}
