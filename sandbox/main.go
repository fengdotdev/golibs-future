package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

type Stream[T any] interface {
	GetError() chan error
	GetValue() chan T
	Get() (chan T, chan error)
	Pause()
	Resume()
	Stop()
}

var _ Stream[int] = (*IntStream)(nil)

type IntStream struct {
	value chan int
	err   chan error
	pause chan bool
}

func NewIntStream() *IntStream {
	return &IntStream{
		value: make(chan int, 1),
		err:   make(chan error, 1),
		pause: make(chan bool),
	}
}

// Get implements future.Stream.
func (i *IntStream) Get() (chan int, chan error) {
	return i.value, i.err
}

// GetError implements future.Stream.
func (i *IntStream) GetError() chan error {
	return i.err
}

// GetValue implements future.Stream.
func (i *IntStream) GetValue() chan int {
	return i.value
}

// Pause implements future.Stream.
func (i *IntStream) Pause() {
	i.pause <- true
}

// Resume implements future.Stream.
func (i *IntStream) Resume() {
	i.pause <- false
}

// Stop implements future.Stream.
func (i *IntStream) Stop() {
	close(i.value)
	close(i.err)
}

func NumberGenerator() *IntStream {
	stream := NewIntStream()

	// clock at 1 second

	go func() {
		i := 0
		paused := false
		for {
			select {
			case p, ok := <-stream.pause:
				if !ok {
					return
				}
				paused = p
			default:
				if paused {
					continue
				}
				select {
				case stream.value <- i:
					i++
				default:
					// evitar bloqueo si nadie está leyendo
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return stream
}
func TimeUP[T any](Stream Stream[T], timer time.Duration) {
	go func() {
		time.Sleep(timer)
		Stream.Stop()
	}()
}

func IntermitenPause[T any](stream Stream[T], timer time.Duration) {
	go func() {
		for {

			stream.Pause()
			time.Sleep(timer)
			stream.Resume()

		}
	}()
}

func StreamReader2(stream *IntStream) {
	for {
		select {
		case v, ok := <-stream.GetValue():
			if !ok {
				return
			}
			println(v)
		case err := <-stream.GetError():
			println(err.Error())
		}
	}
}

func main() {
	stream := NumberGenerator()
	TimeUP(stream, 30*time.Second)
	IntermitenPause(stream, 5*time.Second)

}

func SomeStream(ctx context.Context) chan []byte {

	// Create a channel to send data
	dataChan := make(chan []byte)

	// Start a goroutine to send data
	go func() {
		for i := 0; i < 10; i++ {
			dataChan <- []byte(fmt.Sprintf("Data chunk %d", i))
		}
		close(dataChan)
	}()

	return dataChan
}

type FutureStream struct {
	Channel   chan []byte
	ErrorChan chan error
}

func NewFutureStream() FutureStream {
	return FutureStream{
		Channel:   make(chan []byte),
		ErrorChan: make(chan error),
	}
}
func FIleStream2(ctx context.Context, filepath string) FutureStream {
	future := FutureStream{
		Channel:   make(chan []byte, 10),
		ErrorChan: make(chan error, 1),
	}

	go func() {
		defer close(future.Channel)
		defer close(future.ErrorChan)

		file, err := os.Open(filepath)
		if err != nil {
			future.ErrorChan <- err
			return
		}
		defer file.Close()

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				future.ErrorChan <- err
				return
			}
			if n == 0 {
				break
			}

			select {
			case future.Channel <- buf[:n]:
			case <-ctx.Done():
				future.ErrorChan <- ctx.Err()
				return
			}
		}

		future.ErrorChan <- nil
	}()

	return future
}

func FIleStream(ctx context.Context, filepath string) FutureStream {

	future := NewFutureStream()

	// Start a goroutine to send data
	go func() {

		file, err := os.Open(filepath)
		if err != nil {
			future.ErrorChan <- err
			return
		}
		defer file.Close()
		// Read the file in chunks
		buf := make([]byte, 1024)

		for {
			n, err := file.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				future.ErrorChan <- err
				return
			}
			if n == 0 {
				break
			}
			// Send the data to the channel
			future.Channel <- buf[:n]
		}

		close(future.Channel)
		close(future.ErrorChan)
	}()

	return future
}



func StreamFileReader(ctx context.Context, future FutureStream) {

	for data := range future.Channel {
		fmt.Println("Received:", string(data))
	}

	if err := <-future.ErrorChan; err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Stream completed successfully")
	}
}
