# nanoevent

Tiny event emitter for Go

## Example

```go
eventEmitter := nanoevent.New()

unsubscribe := eventEmitter.On("event", func(message string) {
    fmt.Println(message)
})

eventEmitter.Emit("event", "Hello Go!") // Prints Hello Go

unsubscribe()

eventEmitter.Emit("event", "Hi Go!") // Prints nothing
```

## API

### func New

```go
func New() *EventEmitter
```

New creates a new EventEmitter.

### func (\*EventEmitter) On

```go
func (eventEmitter *EventEmitter) On(event string, listener listenerFunc) func()
```

On subscribes a listener to an event.

Returns unsubscribe function that can be used to remove subscribed listener.

### func (\*EventEmitter) Emit

```go
func (eventEmitter *EventEmitter) Emit(event string, args ...interface{})
```

Emit emits an event to the listeners.

## License

MIT
