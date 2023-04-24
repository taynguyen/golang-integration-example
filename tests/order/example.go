package order

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func (s *ExampleSuite) SetupSuite() {
	println("SetupSuite")
}

func (s *ExampleSuite) TearDownSuite() {
	println("TearDownSuite")
}

func (s *ExampleSuite) SetupTest() {
	println("SetupTest")
}

func (s *ExampleSuite) SetupSubTest() {
	println("SetupSubTest")
}

func (s *ExampleSuite) TearDownTest() {
	println("TearDownTest")
}

func (s *ExampleSuite) TearDownSubTest() {
	println("TearDownSubTest")
}

func (s *ExampleSuite) TestCreate1() {
	println("TestCreate1")
}

func (s *ExampleSuite) TestCreate2() {
	println("TestCreate2")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}
