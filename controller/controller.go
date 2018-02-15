package controller

import "context"

type Controller interface {
	Run(ctx context.Context) error
	AddHook(hook ControllerHook) error
	RemoveHook(hook ControllerHook) error
}
