package controller

type ControllerHook interface {
	OnAdd(object interface{})
	OnUpdate(object interface{})
	OnDelete(object interface{})
}
