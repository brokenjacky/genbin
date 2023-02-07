package binTemplate

var logTpl = `
package {{.Package}}

import "log"

type Log interface {
    Infof(template string, args ...interface{})
}

type DefaultLogger struct {
}

func (l *DefaultLogger) Infof(template string, args ...interface{}) {
    log.Printf(template, args...)
}

var (
    Logger Log = &DefaultLogger{}
)

`

func init() {
    Register("log", logTpl)
}
