package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

var (
	logger = logrus.New()
)

func init() {
	logger.SetOutput(getOutputFile())
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})

	//logger.AddHook(&ReportHook{})

	// 设置不同的formatter到不同的hook中
	consoleHook := NewColorfulConsoleHook(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})

	logger.AddHook(consoleHook)
}

func Debug(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Debugln(replaceNextLine(msg))
}

func Info(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Infoln(replaceNextLine(msg))
}

func Warn(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Warnln(replaceNextLine(msg))
}

func Error(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Errorln(replaceNextLine(msg))
}

func Fatal(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Fatalln(replaceNextLine(msg))
}

func replaceNextLine(v ...interface{}) string {
	str := fmt.Sprint(v...)
	return strings.ReplaceAll(str, "\n", "\\n")
}

func getOutputFile() *os.File {
	fileName := fmt.Sprintf("log%s.log", time.Now().Format("060102"))

	outputDir := "logging/log_files"

	baseDir, err := os.Getwd()
	if err != nil {
		logger.Warn(err)
	}

	logFilePath := filepath.Join(baseDir, outputDir, fileName)

	if stat, _ := os.Stat(logFilePath); stat != nil {
		file, err := os.OpenFile(logFilePath, os.O_APPEND, os.ModeAppend)
		if err != nil {
			logger.Errorln(err)
			return nil
		}
		return file
	}

	file, err := os.Create(logFilePath)
	if err != nil {
		logger.Errorln(err)
		return nil
	}

	return file
}

func fieldsToMap(fields ...interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, field := range fields {
		for key, value := range StructToMap(field) {
			result[key] = value
		}
	}

	return result
}

func StructToMap(obj interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj)

	// 如果是指针，则获取指针指向的元素
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	result := make(map[string]interface{})

	// 确保处理的是结构体
	if val.Kind() != reflect.Struct {
		return result // 返回空map，因为传入的不是结构体或结构体指针
	}

	typ := val.Type() // 获取实际类型，而不是指针类型

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i).Interface() // 获取字段的值
		fieldName := typ.Field(i).Name         // 获取字段的名称
		result[fieldName] = fieldValue         // 将字段名和字段值添加到结果map中
	}

	return result
}
