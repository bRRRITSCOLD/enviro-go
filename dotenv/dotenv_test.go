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

func (s *DotEnvTestSuite) TestLoad_FileNotExist() {
	de := &DotEnv{Path: filepath.Join(s.tmpDir, ".env")}

	err := de.Load()

	s.Error(err)
	var target *NotExistError
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

func TestDotEnvTestSuite(t *testing.T) {
	suite.Run(t, new(DotEnvTestSuite))
}
