package log

import "github.com/sirupsen/logrus"

// init sets the global logrus settings
func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

//Debug Debug wrapper method
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

//Debugf Debugf wrapper method
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args)
}

//Fatal wrapper method
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

//Error wrapper method
func Error(args ...interface{}) {
	logrus.Error(args...)
}

//Errorf wrapper method
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args)
}

//WithField wrapper method
func WithField(key string, arg interface{}) *logrus.Entry {
	return logrus.WithField(key, arg)
}

//WithFields wrapper method
func WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.WithFields(fields)
}

//WithError wrapper method
func WithError(err error) *logrus.Entry {
	return logrus.WithError(err)
}

// Method allows one to quickly log the method field
func Method(method interface{}) *logrus.Entry {
	return logrus.WithField("method", method)
}

// Function allows one to quickly log the method field
func Function(function interface{}) *logrus.Entry {
	return logrus.WithField("function", function)
}
