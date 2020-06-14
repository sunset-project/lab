package lab

import (
	"testing"

	"github.com/sunset-project/lab/reporting"
)

// Session represents a test session, within a single test function
type Session interface {
	Context(...interface{})
	Test(...interface{})
	Assertion() Assertion
}

// NewSession creates a new test session
func NewSession(t *testing.T, reporter reporting.Reporter) Session {
	return NewTest(t, reporter)
}

// StartSession is a utility function to interact with a test Session without holding its reference. The returned tuple is the Context function, the Test function and the Assert function that work on the Session object
func StartSession(t *testing.T) (func(...interface{}), func(...interface{}), Assertion) {
	reporter := DefaultSessionReporter()
	session := NewSession(t, reporter)
	return session.Context, session.Test, session.Assertion()
}

// DefaultSessionReporter returns the default configured Reporter for `lab`
func DefaultSessionReporter() reporting.Reporter {
	return reporting.NewProxyReporter(reporting.StdoutReporter())
}