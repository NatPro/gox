package internal

import (
	"flag"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/db"
	"github.com/maprost/gox/internal/log"
)

type baseCommand struct {
	log  args.LogFlag
	file args.FileFlag
}

func (cmd *baseCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.log.DefineFlag(fs)
	cmd.file.DefineFlag(fs)
}

func (cmd *baseCommand) init() {
	if cmd.log.LogLevel == "debug" {
		log.InitLogger(log.LevelDebug)
	} else if cmd.log.LogLevel == "warn" {
		log.InitLogger(log.LevelWarn)
	} else {
		log.InitLogger(log.LevelInfo)
	}

	// load config file
	err := gxcfg.InitConfig(cmd.file.File, gxcfg.DatabaseAccessLink)
	checkFatal(err, "Can't init config: ")
}

func startDatabases(hdd bool) error {
	for _, dbConf := range gxcfg.GetConfig().Database {
		dbx := db.New(dbConf)
		err := dbx.Run(hdd)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
