package conf

import (
	"os"
	"{{ .AppName }}/assets/conf"
	"github.com/go-ini/ini"
	"{{ .AppName }}/lib/log"
)

var (
	Server   ServerCfg
	Database DatabaseCfg
	Logger   LoggerCfg
)

func Init(customConf string) error {
	cfgFile, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, conf.MustAsset("sample/app.ini"))

	if customConf != "" && fileExists(customConf) {
		if err = cfgFile.Append(customConf); err != nil {
			return err
		}
		log.Info("with custom config %q", customConf)
	} else {
		log.Info("custom config %q not found. will use default conf", customConf)
	}

	mapTo(cfgFile, "server", &Server)
	mapTo(cfgFile, "database", &Database)
	mapTo(cfgFile, "logger", &Logger)

	return err
}

func mapTo(cfg *ini.File, section string, v interface{}) {

	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal("Cfg.MapTo Cfg err: %v", err)
	}
	// log.Debug("----%v -->%v", section, v)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}


func MustInit(customConf string) {
	err := Init(customConf)
	if err != nil {
		panic(err)
	}
}
