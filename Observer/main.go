/*
Observer Pattern is a behavioral design pattern.
This pattern allows an instance (called subject) to publish events to other multiple instances (called observers).
These observers subscribe to the subject and hence get notified by events in case of any change happening in the subject.
*/

package main

import "fmt"

// Observer defines the interface for the Observer
type Observer interface {
	Name() string
	CallBack(string)
}

// Subject defines the interface for the Subject
type Subject interface {
	Register(o *Observer)
	Unregister(o *Observer)
	NotifyAll(msg string)
}

// SubjectImpl implements the Observer interface
type SubjectImpl struct {
	observers map[string]Observer
}

func (s *SubjectImpl) Register(o Observer) {
	s.observers[o.Name()] = o
}

func (s *SubjectImpl) Unegister(o Observer) {
	delete(s.observers, o.Name())
}

func (s *SubjectImpl) NotifyAll(msg string) {
	for _, o := range s.observers {
		o.CallBack(msg)
	}
}

// ObserverA implements A observer
type ObserverA struct {
}

func (oa *ObserverA) Name() string {
	return "A"
}

func (oa *ObserverA) CallBack(msg string) {
	fmt.Printf("A receive a message: %v\n", msg)
}

// ObserverB implements B observer
type ObserverB struct {
}

func (ob *ObserverB) Name() string {
	return "B"
}

func (ob *ObserverB) CallBack(msg string) {
	fmt.Printf("B receive a message: %v\n", msg)
}

func main() {
	s := &SubjectImpl{
		observers: make(map[string]Observer),
	}

	oa := &ObserverA{}
	s.Register(oa)

	ob := &ObserverB{}
	s.Register(ob)

	s.NotifyAll("message 1")
	s.NotifyAll("message 2")
}

/*
OUTPUT:
A receive a message: message 1
B receive a message: message 1
A receive a message: message 2
B receive a message: message 2
*/
