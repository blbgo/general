// Package general contains interfaces and code that could be used in abstract ways in a range of
// situations
package general

// DelayCloser is an interface that should be implemented by values that need to be closed but may
// require some time to do so.  Usually because they manage one or more worker threads.
//
// The implementation of DelayCloser.Close must return *immediately* and at some later time on
// *another* thread must write a single item (either nil or an error) to the provided chan.
type DelayCloser interface {
	Close(doneChan chan<- error)
}

// Shutdowner provides a way to request shutdown
type Shutdowner interface {
	// Shutdown indicates to the implementer that shutdown is needed with a possible error
	Shutdown(err error)
}

// Logger is a minimal interface to support logging
type Logger interface {
	Log(v ...interface{}) error
	Logf(format string, v ...interface{}) error
}

// Task represents a process that can be run and its inputs
type Task interface {
	// Name must return the name of the task with passible namespacing separated by slashes
	// Example "Account/Create"
	Name() string
	Inputs() []string
	Run(inputs ...string) error
}
