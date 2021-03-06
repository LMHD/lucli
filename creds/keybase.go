package creds

import (
	"path/filepath"

	"github.com/labstack/gommon/log"
	"github.com/skybet/cali"
)

var (
	keybasePrivDir = "/Volumes/Keybase/private/"

	// TODO: lmhd <-- make this a global config
	keybaseUser = "lmhd"
)

func keybaseDir() (string, error) {
	keybaseDir, err := filepath.EvalSymlinks(keybasePrivDir + keybaseUser)
	if err != nil {
		return "", err
	}
	log.Debugf("Keybase Dir: %s", keybaseDir)

	return keybaseDir, nil
}

func BindAWS(t *cali.Task, args []string) error {
	keybaseDir, err := keybaseDir()
	if err != nil {
		return err
	}

	awsDir, err := t.Bind(keybaseDir+"/aws", "/root/.aws")
	if err != nil {
		return err
	}

	t.AddBinds([]string{awsDir})

	return nil
}

func BindTerraform(t *cali.Task, args []string) error {
	keybaseDir, err := keybaseDir()
	if err != nil {
		return err
	}

	tfDir, err := t.Bind(keybaseDir+"/terraformrc", "/root/.terraformrc")
	if err != nil {
		return err
	}

	t.AddBinds([]string{tfDir})

	return nil
}

// TODO: BindVault
