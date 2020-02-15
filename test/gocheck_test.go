// See: http://labix.org/gocheck
package test

import (
	"fmt"
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	dir string
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *C) {
	fmt.Println("SetUpSuite() called.")
}

func (s *MySuite) SetUpTest(c *C) {
	fmt.Println("SetUpTest() called.")
	s.dir = c.MkDir()
	fmt.Printf("data dir: %s\n", s.dir)
	// Use s.dir to prepare some data.
}

func (s *MySuite) TearDownTest(c *C) {
	fmt.Println("TearDownTest() called.")
}

func (s *MySuite) TearDownSuite(c *C) {
	fmt.Println("TearDownSuite() called.")
}

func (s *MySuite) TestHelloWorld(c *C) {
	fmt.Println("TestHelloWorld() called.")
	//c.Assert(42, Equals, "42")
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

func (s *MySuite) TestByeWorld(c *C) {
	fmt.Println("TestByeWorld() called.")
	// Use the data in s.dir in the test.
}

/*
=== RUN   Test
SetUpSuite() called.
SetUpTest() called.
data dir: /tmp/check-5577006791947779410/0
TestByeWorld() called.
TearDownTest() called.
SetUpTest() called.
data dir: /tmp/check-5577006791947779410/1
TestHelloWorld() called.
TearDownTest() called.
TearDownSuite() called.
OK: 2 passed
--- PASS: Test (0.00s)
*/
