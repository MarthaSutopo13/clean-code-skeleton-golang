package config

import (
	"basesvc_v2/pkg/adapter/db/model"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/configor"
)

var AllConfig struct {
	APPName string
	Env     string
	Dev     struct {
		Port     string
		Services struct{}
		Database struct {
			Postgres struct {
				DB model.DbConfig
			}
			Log struct {
				Mongo struct {
					Host       string
					Database   string
					Collection string
				}
				File struct {
					Location string
				}
			}
			Cache struct {
				Redis struct {
					Host string
					Port string
				}
			}
		}
		App struct {
			Timeout     int64
			AllowOrigin []string
		}
	}
	Prod struct {
		Port     string
		Services struct{}
		Database struct {
			Postgres struct {
				DB model.DbConfig
			}
			Log struct {
				Mongo struct {
					Host       string
					Database   string
					Collection string
				}
				File struct {
					Location string
				}
			}
			Cache struct {
				Redis struct {
					Host string
					Port string
				}
			}
		}
		App struct {
			Timeout     int64
			AllowOrigin []string
		}
	}
}

var UsedConfig struct {
	Port     string
	Services struct{}
	Database struct {
		Postgres struct {
			DB model.DbConfig
		}
		Log struct {
			Mongo struct {
				Host       string
				Database   string
				Collection string
			}
			File struct {
				Location string
			}
		}
		Cache struct {
			Redis struct {
				Host string
				Port string
			}
		}
	}
	App struct {
		Timeout     int64
		AllowOrigin []string
	}
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {

	configor.Load(&AllConfig, basepath+"/config.yml")
	if AllConfig.Env == "prod" {
		UsedConfig = AllConfig.Prod
	} else {
		UsedConfig = AllConfig.Dev
	}
}
