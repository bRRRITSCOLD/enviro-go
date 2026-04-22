package dotenv

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DotEnvErrorsTestSuite struct {
	suite.Suite
}

func (s *DotEnvErrorsTestSuite) SetupSuite() {}

func (s *DotEnvErrorsTestSuite) TearDownSuite() {}

func (s *DotEnvErrorsTestSuite) SetupTest() {}

func (s *DotEnvErrorsTestSuite) TearDownTest() {}

func (s *DotEnvErrorsTestSuite) TestNotExistError() {
	err := &NotExistError{Path: "/some/path/.env"}
	s.Equal("NotExistError: /some/path/.env does not exist", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestStatError() {
	cause := errors.New("permission denied")
	err := &StatError{Path: "/some/path/.env", Cause: cause}
	s.Equal("StatError: cannot stat /some/path/.env: permission denied", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestLoadError() {
	cause := errors.New("unexpected EOF")
	err := &LoadError{Path: "/some/path/.env", Cause: cause}
	s.Equal("LoadError: failed to load /some/path/.env: unexpected EOF", err.Error())
}

func TestDotEnvErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(DotEnvErrorsTestSuite))
}
