package nanoevent

import (
	"testing"
)

func TestNew(t *testing.T) {
	eventEmitter := New()

	if eventEmitter == nil {
		t.Errorf("Expected eventEmitter to not be nil, but got %v", eventEmitter)
	}
}

func TestEmptyListeners(t *testing.T) {
	eventEmitter := New()

	listenersLength := len(eventEmitter.listeners)

	if listenersLength != 0 {
		t.Errorf("Expected listeners to be empty, but got length %d", listenersLength)
	}
}

func TestSubscribeListeners(t *testing.T) {
	eventEmitter := New()

	eventEmitter.On("event1", func() {})
	eventEmitter.On("event2", func() {})
	eventEmitter.On("event2", func() {})

	eventsLength := len(eventEmitter.listeners)

	if eventsLength != 2 {
		t.Errorf("Expected events length to be 2, but got length %d", eventsLength)
	}

	event1ListenersLength := len(eventEmitter.listeners["event1"])

	if event1ListenersLength != 1 {
		t.Errorf("Expected event1 listeners length to be 1, but got length %d", event1ListenersLength)
	}

	event2ListenersLength := len(eventEmitter.listeners["event2"])

	if event2ListenersLength != 2 {
		t.Errorf("Expected event2 listeners length to be 2, but got length %d", event2ListenersLength)
	}
}

func TestEmitEventToListeners(t *testing.T) {
	eventEmitter := New()

	counter := 0

	eventEmitter.On("count", func(number int) {
		counter += number
	})

	eventEmitter.Emit("count", 1)
	eventEmitter.Emit("count", 1)

	if counter != 2 {
		t.Errorf("Expected count to be equal to 2, but got %d", counter)
	}
}

func TestUnsubscribeListener(t *testing.T) {
	eventEmitter := New()

	counter := 0

	unsubscribe := eventEmitter.On("count", func(number int) {
		counter += number
	})

	eventEmitter.Emit("count", 1)

	unsubscribe()

	eventEmitter.Emit("count", 1)

	if counter != 1 {
		t.Errorf("Expected count to be equal to 1, but got %d", counter)
	}
}
