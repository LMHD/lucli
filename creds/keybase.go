package creds

import (
	"path/filepath"

	"github.com/labstack/gommon/log"
	"github.com/skybet/cali"
)

func BindAWS(t *cali.Task, args []string) error {

	// TODO: lmhd <-- make this a global config

	keybaseDir, err := filepath.EvalSymlinks("/keybase/private/lmhd")
	if err != nil {
		return err
	}
	log.Debugf("Keybase Dir: %s", keybaseDir)

	awsDir, err := t.Bind(keybaseDir+"/aws", "/root/.aws")
	if err != nil {
		return err
	}
	t.AddBinds([]string{awsDir})

	return nil
}
