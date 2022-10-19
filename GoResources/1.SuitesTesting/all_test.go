package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {
	suite.Run(t, new(ExampleSuite))
}

func (es *ExampleSuite) TestTrue() {
	es.T().Log(".....Running TestTrue....")
	es.True(true)

}

func (es *ExampleSuite) TestFalse() {
	es.T().Log(".....Running TestFalse....")
	es.True(true)

}

func (es *ExampleSuite) SetupSuite() {
	es.T().Log("SetupSuite")
}

func (es *ExampleSuite) TearDownSuite() {
	es.T().Log("TearDownSuite")
}

func (es *ExampleSuite) SetupTest() {
	es.T().Log("SetupTest")
}

func (es *ExampleSuite) TearDownTest() {
	es.T().Log("TearDownTest")
}

func (es *ExampleSuite) BeforeTest(suiteName, testName string) {
	es.T().Log("BeforeTest")
}

func (es *ExampleSuite) AfterTest(suiteName, testName string) {
	es.T().Log("AfterTest")
}
