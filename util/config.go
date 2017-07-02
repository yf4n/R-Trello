package util

import (
	"flag"
	"github.com/larspensjo/config"
	"log"
	"os"
)

var configFile = flag.String("configfile", os.Getenv("HOME")+"/.R/config.ini", "General configuration file")
var INIT_CONFIG = make(map[string]string)

func GetIniConfig(t string, k string) string {
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}

	if cfg.HasSection(t) {
		section, err := cfg.SectionOptions(t)
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(t, v)
				if err == nil {
					INIT_CONFIG[v] = options
				}
			}
		}
	}

	return INIT_CONFIG[k]
}
