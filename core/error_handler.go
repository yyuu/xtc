package core

import (
  "fmt"
  "log"
  "os"
)

const (
  LOG_DEBUG = iota
  LOG_INFO
  LOG_WARN
  LOG_ERROR
)

type ErrorHandler struct {
  level int
  logger *log.Logger
  errors int
}

func NewErrorHandler(level int) *ErrorHandler {
  return &ErrorHandler { level, log.New(os.Stderr, "", log.LstdFlags), 0 }
}

func (self *ErrorHandler) skip(criteria int) bool {
  return criteria < self.level
}

func levelName(level int) string {
  switch level {
    case LOG_DEBUG: return "DEBUG"
    case LOG_INFO:  return "INFO"
    case LOG_WARN:  return "WARN"
    case LOG_ERROR: return "ERROR"
    default:        return "UNKNOWN"
  }
}

func (self *ErrorHandler) logf(level int, format string, v...interface{}) {
  if ! self.skip(level) {
    s := fmt.Sprintf("%s: %s", levelName(level), fmt.Sprintf(format, v...))
    self.logger.Print(s)
  }
}

func (self *ErrorHandler) log(level int, v...interface{}) {
  if ! self.skip(level) {
    s := fmt.Sprintf("%s: %s", levelName(level), fmt.Sprint(v...))
    self.logger.Print(s)
  }
}

func (self *ErrorHandler) Debugf(format string, v...interface{}) {
  self.logf(LOG_DEBUG, format, v...)
}

func (self *ErrorHandler) Debug(v...interface{}) {
  self.log(LOG_DEBUG, v...)
}

func (self *ErrorHandler) Infof(format string, v...interface{}) {
  self.logf(LOG_INFO, format, v...)
}

func (self *ErrorHandler) Info(v...interface{}) {
  self.log(LOG_INFO, v...)
}

func (self *ErrorHandler) Warnf(format string, v...interface{}) {
  self.logf(LOG_WARN, format, v...)
}

func (self *ErrorHandler) Warn(v...interface{}) {
  self.log(LOG_WARN, v...)
}

func (self *ErrorHandler) Errorf(format string, v...interface{}) {
  self.logf(LOG_ERROR, format, v...)
  self.errors++
}

func (self *ErrorHandler) Error(format string, v...interface{}) {
  self.log(LOG_ERROR, v...)
  self.errors++
}

func (self *ErrorHandler) Fatalf(format string, v...interface{}) {
  self.logger.Fatalf(format, v...)
}

func (self *ErrorHandler) Fatal(v...interface{}) {
  self.logger.Fatal(v...)
}

func (self *ErrorHandler) Panicf(format string, v...interface{}) {
  self.logger.Panicf(format, v...)
}

func (self *ErrorHandler) Panic(v...interface{}) {
  self.logger.Panic(v...)
}

func (self *ErrorHandler) ErrorOccured() bool {
  return 0 < self.errors
}

func (self *ErrorHandler) GetErrors() int {
  return self.errors
}
