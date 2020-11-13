package rest

import (
	"basesvc_v2/internal/config"
	restAPI "basesvc_v2/pkg/infrastructure/rest/api"
	"basesvc_v2/pkg/shared/logger"
	"fmt"
	"log"
)

func main() {
	config := config.GetConfig()
	fmt.Println(config.Logger)
	logConfig := logger.Configuration{
		EnableConsole:     config.Logger.Console.Enable,
		ConsoleJSONFormat: config.Logger.Console.JSON,
		ConsoleLevel:      config.Logger.Console.Level,
		EnableFile:        config.Logger.File.Enable,
		FileJSONFormat:    config.Logger.File.JSON,
		FileLevel:         config.Logger.File.Level,
		FileLocation:      config.Logger.File.Path,
	}

	if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
		log.Fatalf("Could not instantiate log %v", err)
	}

	restAPI.RunServer()
}
