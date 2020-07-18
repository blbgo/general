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

// LoggerFactory is a minimal interface to supports creating loggers
type LoggerFactory interface {
	New(name string) (Logger, error)
}

// Dumper is a minimal interface to support dumping large chunks of data
type Dumper interface {
	Dump(data []byte) error
	DumpObj(obj interface{}) error
}

// DumperFactory is a minimal interface to supports creating dumpers
type DumperFactory interface {
	// New creates a new Dumper with the given name
	// it may need to be closed, use "closer, ok := dumper.(io.Closer)" to find out
	New(name string) (Dumper, error)

	// Dump is a helper method that creates a new Dumper, dumps the one given buffer and closes
	// the dumper
	Dump(name string, data []byte) error

	// DumpObj is a helper method that creates a new Dumper, dumps the one given object and closes
	// the dumper
	DumpObj(name string, obj interface{}) error
}

// Task represents a process that can be run and its inputs
type Task interface {
	// Name must return the name of the task with passible namespacing separated by slashes
	// Example "Account/Create"
	Name() string
	Inputs() []string
	Run(inputs ...string) error
}

// Config allows retreving configuration values
type Config interface {
	Value(section, name string) (string, error)
}

// PersistentState interface allows reading and writing named state objects
type PersistentState interface {
	// Save persists an object under the given name
	// the state object should be json serializable
	Save(name string, state interface{}) error

	// Retrieve fills the provided state object with data that was previously saved under the
	// given name.
	Retrieve(name string, state interface{}) error
}
