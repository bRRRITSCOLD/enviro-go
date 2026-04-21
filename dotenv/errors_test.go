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

func (s *DotEnvErrorsTestSuite) TestDotEnvNotExistError() {
	err := &DotEnvNotExistError{Path: "/some/path/.env"}
	s.Equal("DotEnvNotExistError: /some/path/.env does not exist", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestDotEnvStatError() {
	cause := errors.New("permission denied")
	err := &DotEnvStatError{Path: "/some/path/.env", Cause: cause}
	s.Equal("DotEnvStatError: cannot stat /some/path/.env: permission denied", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestDotEnvLoadError() {
	cause := errors.New("unexpected EOF")
	err := &DotEnvLoadError{Path: "/some/path/.env", Cause: cause}
	s.Equal("DotEnvLoadError: failed to load /some/path/.env: unexpected EOF", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestDotEnvReadError() {
	cause := errors.New("read error")
	err := &DotEnvReadError{Path: "/some/path/.env.example", Cause: cause}
	s.Equal("DotEnvReadError: failed to read /some/path/.env.example: read error", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestDotEnvValidateError() {
	cause := errors.New("stat failed")
	err := &DotEnvValidateError{Path: "/some/path/.env", Cause: cause}
	s.Equal("env: cannot stat /some/path/.env: stat failed", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestDotEnvMissingKeysError_SingleKey() {
	err := &DotEnvMissingKeysError{Keys: []string{"DATABASE_URL"}}
	s.Equal("DotEnvValidateError: missing required keys (declared in example ): DATABASE_URL", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestDotEnvMissingKeysError_MultipleKeys() {
	err := &DotEnvMissingKeysError{Keys: []string{"DATABASE_URL", "SECRET_KEY", "PORT"}}
	s.Equal("DotEnvValidateError: missing required keys (declared in example ): DATABASE_URL, SECRET_KEY, PORT", err.Error())
}

func (s *DotEnvErrorsTestSuite) TestDotEnvMissingKeysError_EmptyKeys() {
	err := &DotEnvMissingKeysError{Keys: []string{}}
	s.Equal("DotEnvValidateError: missing required keys (declared in example ): ", err.Error())
}

func TestDotEnvErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(DotEnvErrorsTestSuite))
}
