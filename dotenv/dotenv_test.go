package dotenv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DotEnvTestSuite struct {
	suite.Suite
	tmpDir string
}

func (s *DotEnvTestSuite) SetupSuite() {}

func (s *DotEnvTestSuite) TearDownSuite() {}

func (s *DotEnvTestSuite) SetupTest() {
	s.tmpDir = s.T().TempDir()
}

func (s *DotEnvTestSuite) TearDownTest() {}

// Load tests
func (s *DotEnvTestSuite) TestLoad_FileNotExist() {
	de := &DotEnv{Path: filepath.Join(s.tmpDir, ".env")}

	err := de.Load()

	s.Error(err)
	var target *DotEnvNotExistError
	s.ErrorAs(err, &target)
	s.Equal(de.Path, target.Path)
}

func (s *DotEnvTestSuite) TestLoad_Success() {
	envPath := filepath.Join(s.tmpDir, ".env")
	s.Require().NoError(os.WriteFile(envPath, []byte("TEST_DOTENV_LOAD_KEY=hello\n"), 0644))
	defer os.Unsetenv("TEST_DOTENV_LOAD_KEY")

	de := &DotEnv{Path: envPath}

	err := de.Load()

	s.NoError(err)
	s.Equal("hello", os.Getenv("TEST_DOTENV_LOAD_KEY"))
}

// Validate tests
func (s *DotEnvTestSuite) TestValidate_ExampleNotExist() {
	de := &DotEnv{
		Path:    filepath.Join(s.tmpDir, ".env"),
		Example: filepath.Join(s.tmpDir, ".env.example"),
	}

	err := de.Validate()

	s.Error(err)
	var target *DotEnvNotExistError
	s.ErrorAs(err, &target)
}

func (s *DotEnvTestSuite) TestValidate_MissingKeys() {
	examplePath := filepath.Join(s.tmpDir, ".env.example")
	s.Require().NoError(os.WriteFile(examplePath, []byte("TEST_DOTENV_REQUIRED_KEY=\n"), 0644))
	os.Unsetenv("TEST_DOTENV_REQUIRED_KEY")

	de := &DotEnv{
		Path:    filepath.Join(s.tmpDir, ".env"),
		Example: examplePath,
	}

	err := de.Validate()

	s.Error(err)
	var target *DotEnvMissingKeysError
	s.ErrorAs(err, &target)
	s.Contains(target.Keys, "TEST_DOTENV_REQUIRED_KEY")
}

func (s *DotEnvTestSuite) TestValidate_Success() {
	examplePath := filepath.Join(s.tmpDir, ".env.example")
	s.Require().NoError(os.WriteFile(examplePath, []byte("TEST_DOTENV_PRESENT_KEY=\n"), 0644))
	s.T().Setenv("TEST_DOTENV_PRESENT_KEY", "somevalue")

	de := &DotEnv{
		Path:    filepath.Join(s.tmpDir, ".env"),
		Example: examplePath,
	}

	err := de.Validate()

	s.NoError(err)
}

func TestDotEnvTestSuite(t *testing.T) {
	suite.Run(t, new(DotEnvTestSuite))
}
