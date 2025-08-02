package main

import (
	"encoding/json"
	"os"
	"path"

	wl_fs "github.com/wsva/lib_go/fs"
	wl_int "github.com/wsva/lib_go_integration"
)

type MainConfig struct {
	CA         CA         `json:"CA"`
	CertServer CertServer `json:"CertServer"`
}

var (
	MainConfigFile = "pki_config.json"
	PKIPath        = wl_int.DirPKI
)

var mainConfig MainConfig

func initGlobals() error {
	basepath, err := wl_fs.GetExecutableFullpath()
	if err != nil {
		return err
	}
	MainConfigFile = path.Join(basepath, MainConfigFile)
	PKIPath = path.Join(basepath, PKIPath)

	contentBytes, err := os.ReadFile(MainConfigFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contentBytes, &mainConfig)
	if err != nil {
		return err
	}
	return nil
}
