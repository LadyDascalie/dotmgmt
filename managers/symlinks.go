package managers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type SymlinkCollection struct {
	Symlinks  []Symlink `json:"symlinks"`
	Timestamp string `json:"timestamp,omitempty"`
}

type Symlink struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

func (l *SymlinkCollection) Save(conf *Config) {
	j, err := json.MarshalIndent(l, "", "	")
	if err != nil {
		log.Println("Cannot marshall JSON:", err)
	}

	err = ioutil.WriteFile(conf.Path, j, 0755)
	if err != nil {
		log.Println("Cannot write config file:", err)
	}
}
