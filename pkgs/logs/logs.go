package logs

import (
	"io"
	"os"
	"sync"

	"todo_list/src/conf"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

const TIMESTAMP = "2006/01/02 15:04:05.00"
const LOG_PATh = "./pkgs/logs/server.log"

func Init() {
	once.Do(func() {
		var writers []io.Writer
		// Dev: Colored console only
		console := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: TIMESTAMP,
		}
		writers = append(writers, console)
		if !conf.Env.IS_DEV {
			// Production: Stdout + Rotating file
			fileWriter := &lumberjack.Logger{
				Filename:   LOG_PATh,
				MaxSize:    100,
				MaxBackups: 3,
				MaxAge:     28,
				Compress:   true,
			}
			writers = append(writers, fileWriter)
		}
		// Optimized for Zerolog specifically
		multi := zerolog.MultiLevelWriter(writers...)
		log.Logger = zerolog.New(multi).With().Timestamp().Caller().Logger()
		zerolog.TimeFieldFormat = TIMESTAMP
	})
}
