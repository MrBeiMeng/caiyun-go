package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

// ColorfulConsoleHook 自定义的Console hook，用于控制台输出带颜色的日志
type ColorfulConsoleHook struct {
	formatter logrus.Formatter
}

// NewColorfulConsoleHook 创建一个新的ColorfulConsoleHook实例
func NewColorfulConsoleHook(formatter logrus.Formatter) *ColorfulConsoleHook {
	return &ColorfulConsoleHook{
		formatter: formatter,
	}
}

// Fire 实现Hook接口中的Fire方法，用于处理日志事件
func (hook *ColorfulConsoleHook) Fire(entry *logrus.Entry) error {

	bytes, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(bytes)
	return err
}

// Levels 实现Hook接口中的Levels方法，指定hook处理的日志级别
func (hook *ColorfulConsoleHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//// ReportHook 实现 logrus.Hook 接口
//type ReportHook struct{}
//
//func (hook *ReportHook) Levels() []logrus.Level {
//	// 返回 hook 感兴趣的日志级别
//	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
//}
//
//func (hook *ReportHook) Fire(entry *logrus.Entry) error {
//	// 构建报告主题
//	subject := fmt.Sprintf("日志级别: %s", entry.Level)
//
//	// 构建 Markdown 正文
//	var fields []string
//	for key, value := range entry.Data {
//		fields = append(fields, fmt.Sprintf("**%s**: %v", key, value))
//	}
//	markdownBody := fmt.Sprintf("### 消息\n%s\n\n### 字段\n%s", entry.Message, strings.Join(fields, "\n"))
//
//	// 调用报告方法
//	go reporter.Report(subject, markdownBody)
//	return nil
//}
