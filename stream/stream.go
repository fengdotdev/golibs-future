package stream

import "sync"

// Stream represents a sequence of asynchronous events.
type Stream[T any] interface {
	Listen(handler func(T)) Subscription
}

// StreamController allows for the creation and control of a Stream.
type StreamController[T any] struct {
	output chan T
	subs   []func(T)
	mu     sync.Mutex
	closed bool
}

// NewStreamController creates a new StreamController and its associated Stream.
func NewStreamController[T any]() (*StreamController[T], Stream[T]) {
	c := &StreamController[T]{
		output: make(chan T),
		subs:   make([]func(T), 0),
	}
	go c.process()
	return c, c
}

// Add adds a new event to the Stream.
func (c *StreamController[T]) Add(event T) {
	if !c.closed {
		c.output <- event
	}
}

// Close closes the Stream and prevents adding new events.
func (c *StreamController[T]) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.closed {
		close(c.output)
		c.closed = true
	}
}

func (c *StreamController[T]) process() {
	for event := range c.output {
		c.mu.Lock()
		for _, sub := range c.subs {
			// Execute handlers in a goroutine to avoid blocking the stream
			go sub(event)
		}
		c.mu.Unlock()
	}
}

// Listen subscribes to the Stream and registers a handler function.
func (c *StreamController[T]) Listen(handler func(T)) Subscription {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.subs = append(c.subs, handler)
	return &streamSubscription{
		unsubscribe: func() {
			c.mu.Lock()
			defer c.mu.Unlock()
			for i, sub := range c.subs {
				if &sub == &handler {
					c.subs = append(c.subs[:i], c.subs[i+1:]...)
					break
				}
			}
		},
	}
}

// Subscription represents an active subscription to a Stream.
type Subscription interface {
	Cancel()
}

type streamSubscription struct {
	unsubscribe func()
}

// Cancel stops the subscription from receiving further events.
func (s *streamSubscription) Cancel() {
	s.unsubscribe()
}
