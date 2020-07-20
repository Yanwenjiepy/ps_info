package logger

import (
	"errors"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"ps_info/pkg/config"
)

const (
	UnavailableLogLevel = "unavailable log level"

	UnavailableLogFile = "unavailable log file path"

	UnavailableLogFileMaxSize = "unavailable log file max size"

	UnavailableLogFileMaxAge = "unavailable log file max age"

	UnavailableLogFileMaxBackups = "unavailable log file max backups"

	UnavailableCompressFlag = "unavailable compress flag"

	UnavailableLocalTimeFlag = "unavailable local time flag"
)

var Log *zap.Logger

func InitLog(logConfig config.LoggerConfig) error {

	level, err := getLogLevel(logConfig.LogLevel)
	if err != nil {
		return err
	}

	fileLogger, err := newFileLogger(logConfig)
	if err != nil {
		return err
	}

	fileWriter := zapcore.AddSync(fileLogger)
	consoleErrWriter := zapcore.Lock(os.Stderr)

	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	fileCore := zapcore.NewCore(fileEncoder, fileWriter, level)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleErrWriter, level)
	core := zapcore.NewTee(fileCore, consoleCore)

	Log = zap.New(core, zap.AddCaller())
	return nil
}

func newFileLogger(logConfig config.LoggerConfig) (*lumberjack.Logger, error) {

	logFilepath, err := getLogFilepath(logConfig.LogPath)
	if err != nil {
		return nil, err
	}
	log.Println("[Init logger] log file path: ", logFilepath)

	maxSize, err := getLogFileMaxSize(logConfig.LogFileMaxSize)
	if err != nil {
		return nil, err
	}
	log.Printf("[Init logger] log file max size: %d MB\n", maxSize)

	maxAge, err := getLogFileMaxAge(logConfig.LogFileMaxAge)
	if err != nil {
		return nil, err
	}
	log.Printf("[Init logger] log file max age: %d day\n", maxAge)

	maxBackups, err := getLogFileMaxBackups(logConfig.LogFileMaxBackups)
	if err != nil {
		return nil, err
	}
	log.Println("[Init logger] log file max backups: ", maxBackups)

	localTime, err := isUseLocalTime(logConfig.LocalTime)
	if err != nil {
		return nil, err
	}
	log.Println("[Init logger] log file use local time: ", localTime)

	compress, err := isCompressLogFile(logConfig.Compress)
	if err != nil {
		return nil, err
	}
	log.Println("[Init logger] log file use compress: ", compress)

	fileLogger := lumberjack.Logger{
		Filename:   logFilepath,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		LocalTime:  localTime,
		Compress:   compress,
	}

	return &fileLogger, nil
}

func getLogFilepath(filepath string) (string, error) {

	ErrUnavailableLogFilepath := errors.New(UnavailableLogFile)
	if filepath == "" {
		return "", ErrUnavailableLogFilepath
	}

	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return "", ErrUnavailableLogFilepath
	}
	_ = f.Close()

	return filepath, nil
}

func getLogLevel(level string) (zapcore.Level, error) {

	logLevel := zap.InfoLevel

	switch level {
	case "debug":
		logLevel = zap.DebugLevel

	case "info":
		logLevel = zap.InfoLevel

	case "warning":
		logLevel = zap.WarnLevel

	case "error":
		logLevel = zap.ErrorLevel

	case "fatal":
		logLevel = zap.FatalLevel

	default:
		ErrUnavailableLogLevel := errors.New(UnavailableLogLevel)
		return logLevel, ErrUnavailableLogLevel
	}

	return logLevel, nil
}

func getLogFileMaxSize(size int) (int, error) {

	if size > 0 {
		return size, nil
	}

	ErrUnavailableLogFileMaxSize := errors.New(UnavailableLogFileMaxSize)
	return 0, ErrUnavailableLogFileMaxSize
}

func getLogFileMaxAge(age int) (int, error) {

	if age < 0 {
		ErrUnavailableLogFileMaxAge := errors.New(UnavailableLogFileMaxAge)
		return 0, ErrUnavailableLogFileMaxAge
	}

	return age, nil
}

func getLogFileMaxBackups(backups int) (int, error) {

	if backups < 0 {
		ErrUnavailableLogFileMaxBackups := errors.New(UnavailableLogFileMaxBackups)
		return 0, ErrUnavailableLogFileMaxBackups
	}

	return backups, nil
}

func isUseLocalTime(isLocalTime bool) (bool, error) {
	switch isLocalTime {

	case false:
		return false, nil

	case true:
		return true, nil

	default:
		ErrUnavailableLocalTimeFlag := errors.New(UnavailableLocalTimeFlag)
		return false, ErrUnavailableLocalTimeFlag
	}
}

func isCompressLogFile(isCompress bool) (bool, error) {
	switch isCompress {

	case false:
		return false, nil

	case true:
		return true, nil

	default:
		ErrUnavailableCompressFlag := errors.New(UnavailableCompressFlag)
		return false, ErrUnavailableCompressFlag
	}
}
