package managers

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	ManagementFolder = ".dotmgmt"
	ManagementFile   = "mgmt"
)

type Config struct {
	Home string
	Path string
}

func (c *Config) New() *Config {
	c.Path = c.GetPath()
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	c.Home = u.HomeDir

	return c
}

func (c *Config) Exists() bool {
	_, err := os.Stat(c.Path)
	if err != nil {
		return false
	}

	return true
}

func (c *Config) MakeFile() {
	os.Mkdir(filepath.Join(c.Home, ManagementFolder), 0755)
	os.Create(c.Path)
}

func (c *Config) GetPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	return filepath.Join(usr.HomeDir, ManagementFolder, ManagementFile)
}
