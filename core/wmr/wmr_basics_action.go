package wmr

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func (WMR *WMR) LoadBasics() (err error) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.Infoln(WMR.Confer.Opts.ApptypeConf, "Basics load success. ")
	return nil
}
