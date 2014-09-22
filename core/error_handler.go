package core

import (
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
}

func NewErrorHandler(level int) *ErrorHandler {
  return &ErrorHandler { level, log.New(os.Stderr, "ErrorHandler: ", log.LstdFlags) }
}

func (self *ErrorHandler) skip(criteria int) bool {
  return criteria < self.level
}

func (self *ErrorHandler) Debugf(format string, v...interface{}) {
  if ! self.skip(LOG_DEBUG) {
    self.logger.Printf(format, v...)
  }
}

func (self *ErrorHandler) Debugln(v...interface{}) {
  if ! self.skip(LOG_DEBUG) {
    self.logger.Println(v...)
  }
}

func (self *ErrorHandler) Debug(v...interface{}) {
  if ! self.skip(LOG_DEBUG) {
    self.logger.Print(v...)
  }
}

func (self *ErrorHandler) Infof(format string, v...interface{}) {
  if ! self.skip(LOG_INFO) {
    self.logger.Printf(format, v...)
  }
}

func (self *ErrorHandler) Infoln(v...interface{}) {
  if ! self.skip(LOG_INFO) {
    self.logger.Println(v...)
  }
}

func (self *ErrorHandler) Info(v...interface{}) {
  if ! self.skip(LOG_INFO) {
    self.logger.Print(v...)
  }
}

func (self *ErrorHandler) Warnf(format string, v...interface{}) {
  if ! self.skip(LOG_WARN) {
    self.logger.Printf(format, v...)
  }
}

func (self *ErrorHandler) Warnln(v...interface{}) {
  if ! self.skip(LOG_WARN) {
    self.logger.Println(v...)
  }
}

func (self *ErrorHandler) Warn(v...interface{}) {
  if ! self.skip(LOG_WARN) {
    self.logger.Print(v...)
  }
}

func (self *ErrorHandler) Errorf(format string, v...interface{}) {
//if ! self.skip(LOG_ERROR) {
//  self.logger.Printf(format, v...)
//}
  self.logger.Fatalf(format, v...)
}

func (self *ErrorHandler) Errorln(v...interface{}) {
//if ! self.skip(LOG_ERROR) {
//  self.logger.Println(v...)
//}
  self.logger.Fatalln(v...)
}

func (self *ErrorHandler) Error(v...interface{}) {
//if ! self.skip(LOG_ERROR) {
//  self.logger.Print(v...)
//}
  self.logger.Fatal(v...)
}

func (self *ErrorHandler) Fatalf(format string, v...interface{}) {
  self.logger.Fatalf(format, v...)
}

func (self *ErrorHandler) Fatalln(v...interface{}) {
  self.logger.Fatalln(v...)
}

func (self *ErrorHandler) Fatal(v...interface{}) {
  self.logger.Fatal(v...)
}

func (self *ErrorHandler) Panicf(format string, v...interface{}) {
  self.logger.Panicf(format, v...)
}

func (self *ErrorHandler) Panicln(v...interface{}) {
  self.logger.Panicln(v...)
}

func (self *ErrorHandler) Panic(v...interface{}) {
  self.logger.Panic(v...)
}
