package logging

import (
	"fmt"
	"time"

	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", settings.AppConfigSetting.App.RuntimeRootPath, settings.AppConfigSetting.App.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		time.Now().Format(settings.AppConfigSetting.App.TimeFormat),
		settings.AppConfigSetting.App.LogFileExt,
	)
}
