package controller

// Hook represents actions that can be injected into the controller.
type Hook interface {
	// OnAdd runs after the controller finishes processing the added object.
	OnAdd(object interface{})
	// OnUpdate runs after the controller finishes processing the updated object.
	OnUpdate(object interface{})
	// OnDelete runs after the controller finishes processing the deleted object.
	OnDelete(object interface{})
}
