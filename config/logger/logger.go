package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type LoggerFormatter struct {
	// Format defines the logging tags
	// Optional. Default: ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n
	LogFormat string
	// LogFormatType string value is either file or out
	LogFormatType string
	// LogFormatType see https://programming.guide/go/format-parse-string-time-date-example.html
	LogFormatDate string
	// LogFormatTimeZone can be specified, such as "UTC" and "America/New_York" and "Asia/Chongqing", etc
	// Optional. Default: "Local"
	LogFormatTimeZone string
	// Filename logger
	LogFormatFileName string
}

var defaultConfig = LoggerFormatter{
	LogFormat:         "[${ip}]:${port} ${pid} ${locals:requestid}] ${status} - ${method} ${path}​\n",
	LogFormatType:     "out",
	LogFormatDate:     "02-Jan-2006",
	LogFormatTimeZone: "Asia/Jakarta",
}

func LogFormatterNew(format string, formatType string, formatDate string, timezone string, filename string) *LoggerFormatter {
	if format == "" {
		format = defaultConfig.LogFormat
	}
	if formatType == "" {
		formatType = defaultConfig.LogFormatType
	}
	if formatDate == "" {
		formatDate = defaultConfig.LogFormatDate
	}
	if timezone == "" {
		timezone = defaultConfig.LogFormatTimeZone
	}
	if filename == "" {
		filename = defaultConfig.LogFormatFileName
	}
	if format == "file" && filename == "" {
		format = defaultConfig.LogFormat
		log.Println("Set config to default out because filename is empty")
	}
	return &LoggerFormatter{
		LogFormat:         format,
		LogFormatType:     formatType,
		LogFormatDate:     formatDate,
		LogFormatTimeZone: timezone,
		LogFormatFileName: filename,
	}
}

func (lg *LoggerFormatter) SetLoggerFormat(App *fiber.App) {
	if lg.LogFormatType == "out" {
		App.Use(logger.New(logger.Config{
			Format:     lg.LogFormat,
			TimeFormat: lg.LogFormatDate,
			TimeZone:   lg.LogFormatTimeZone,
		}))
	} else {
		file, err := os.OpenFile(fmt.Sprintf("./%s.log", lg.LogFormatFileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer file.Close()
		App.Use(logger.New(logger.Config{
			Format:     lg.LogFormat,
			TimeFormat: lg.LogFormatDate,
			TimeZone:   lg.LogFormatTimeZone,
			Output:     file,
		}))
	}

}
