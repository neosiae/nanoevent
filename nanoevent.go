package nanoevent

import (
	"reflect"
)

type listener interface{}

// EventEmitter struct
type EventEmitter struct {
	listeners map[string][]listener
}

// New creates a new EventEmitter
func New() *EventEmitter {
	eventEmitter := new(EventEmitter)
	eventEmitter.listeners = make(map[string][]listener)
	return eventEmitter
}

// On subscribes a listener to an event
func (eventEmitter *EventEmitter) On(event string, listener listener) func() {
	listeners := eventEmitter.listeners
	listeners[event] = append(listeners[event], listener)

	return func() {
		for i, l := range listeners[event] {
			if reflect.ValueOf(l).Pointer() == reflect.ValueOf(listener).Pointer() {
				listeners[event] = append(listeners[event][:i], listeners[event][i+1:]...)
			}
		}

		if len(listeners[event]) == 0 {
			delete(listeners, event)
		}
	}
}

// Emit emits an event to the listeners
func (eventEmitter *EventEmitter) Emit(event string, args ...interface{}) {
	listeners := eventEmitter.listeners[event]
	vargs := make([]reflect.Value, len(args))

	for i, v := range args {
		vargs[i] = reflect.ValueOf(v)
	}

	for _, listener := range listeners {
		fnValue := reflect.ValueOf(listener)
		fnValue.Call(vargs)
	}
}
