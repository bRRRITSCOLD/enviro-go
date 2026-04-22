package enviro

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnvErrorsTestSuite struct {
	suite.Suite
}

func (s *EnvErrorsTestSuite) SetupSuite() {}

func (s *EnvErrorsTestSuite) TearDownSuite() {}

func (s *EnvErrorsTestSuite) SetupTest() {}

func (s *EnvErrorsTestSuite) TearDownTest() {}

func (s *EnvErrorsTestSuite) TestParseError() {
	cause := errors.New("missing required field")
	err := &ParseError{Cause: cause}
	s.Equal("ParseError: failed to parse env: missing required field", err.Error())
}

func TestEnvErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(EnvErrorsTestSuite))
}
