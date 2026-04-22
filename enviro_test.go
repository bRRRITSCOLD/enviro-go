package enviro

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnvTestSuite struct {
	suite.Suite
	tmpDir string
}

func (s *EnvTestSuite) SetupSuite() {}

func (s *EnvTestSuite) TearDownSuite() {}

func (s *EnvTestSuite) SetupTest() {
	s.tmpDir = s.T().TempDir()
}

func (s *EnvTestSuite) TearDownTest() {}

type testEnvConfig struct {
	AppName string `env:"TEST_ENV_APP_NAME"`
	Port    int    `env:"TEST_ENV_APP_PORT,notEmpty" envDefault:"8080"`
}

func (s *EnvTestSuite) TestParse_EmptyConfig_UsesDefaults() {
	s.T().Setenv("TEST_ENV_APP_NAME", "myapp")
	s.T().Setenv("TEST_ENV_APP_PORT", "9090")

	result, err := Parse[testEnvConfig](EnvConfig{})

	s.NoError(err)
	s.NotNil(result)
	s.Equal("myapp", result.Config().AppName)
	s.Equal(9090, result.Config().Port)
}

func (s *EnvTestSuite) TestParse_DotEnvNotExist() {
	_, err := Parse[testEnvConfig](EnvConfig{
		Path: filepath.Join(s.tmpDir, ".env"),
	})

	s.Error(err)
}

func (s *EnvTestSuite) TestParse_WithDotEnv_LoadsValues() {
	envPath := filepath.Join(s.tmpDir, ".env")
	s.Require().NoError(os.WriteFile(envPath, []byte("TEST_ENV_APP_NAME=fromfile\nTEST_ENV_APP_PORT=7070\n"), 0644))
	defer func() {
		os.Unsetenv("TEST_ENV_APP_NAME")
		os.Unsetenv("TEST_ENV_APP_PORT")
	}()

	result, err := Parse[testEnvConfig](EnvConfig{Path: envPath})

	s.NoError(err)
	s.NotNil(result)
	s.Equal("fromfile", result.Config().AppName)
	s.Equal(7070, result.Config().Port)
}

func (s *EnvTestSuite) TestConfig_ReturnsUnderlyingConfig() {
	s.T().Setenv("TEST_ENV_APP_NAME", "conftest")
	s.T().Setenv("TEST_ENV_APP_PORT", "4000")

	result, err := Parse[testEnvConfig](EnvConfig{})
	s.Require().NoError(err)

	cfg := result.Config()

	s.Equal("conftest", cfg.AppName)
	s.Equal(4000, cfg.Port)
}

func TestEnvTestSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}
