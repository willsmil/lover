package logger

import (
	"log"
	"os"
)

var Logger *logger

const (
	ERROR = iota
)

type logger struct {
}

func init() {
	Logger = &logger{}
	writer, err := os.OpenFile("logs/lover.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic("init log failed, err:" + err.Error())
	}
	log.SetOutput(writer)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}

func Error(fomat string, arg ...interface{}) {
	log.Printf("[ERROR] "+fomat, arg...)
}
func (l *logger) Info(fomat string, arg ...interface{}) {
	log.Printf("[INFO] "+fomat, arg...)
}
func (l *logger) Debug(fomat string, arg ...interface{}) {
	log.Printf("[DEBUG] "+fomat, arg...)
}
func (l *logger) Warn(fomat string, arg ...interface{}) {
	log.Printf("[WARN] "+fomat, arg...)
}
func (l *logger) Fatal(fomat string, arg ...interface{}) {
	log.Printf("[FATAL] "+fomat, arg...)
}
