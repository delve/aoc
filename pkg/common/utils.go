package common

import ("github.com/sirupsen/logrus"
"runtime"
)

// Check is shorthand for dieing if e contains a value
func Check(e error) {
	if e != nil {
		logrus.Fatal(e)
	}
}

func Trace(msg ...string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	if len(msg) > 0 {
		logrus.Info(msg)
	}
	logrus.Infof("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}
