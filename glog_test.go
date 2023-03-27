package myglog

import (
	"os"
	"testing"
)

const (
	sampleMsg  = "<sample message>"
	sampleCode = 200
)

func TestTrace(t *testing.T) {
	V(2).Trace("Trace: ", sampleMsg)
	V(2).Traceln("Traceln:", sampleMsg)
	V(2).Tracef("Tracef: %s, code: %d\n", sampleMsg, sampleCode)
}

func TestDebug(t *testing.T) {
	V(1).Debug("Debug: ", sampleMsg)
	V(1).Debugln("Debugln:", sampleMsg)
	V(1).Debugf("Debugf: %s, code: %d\n", sampleMsg, sampleCode)
}

func TestInfo(t *testing.T) {
	Info("Info: ", sampleMsg)
	Infoln("Infoln:", sampleMsg)
	Infof("Infof: %s, code: %d\n", sampleMsg, sampleCode)
}

func TestWarn(t *testing.T) {
	Warn("Warn: ", sampleMsg)
	Warnln("Warnln:", sampleMsg)
	Warnf("Warnf: %s, code: %d\n", sampleMsg, sampleCode)
}

func TestError(t *testing.T) {
	Error("Error: ", sampleMsg)
	Errorln("Errorln:", sampleMsg)
	Errorf("Errorf: %s, code: %d\n", sampleMsg, sampleCode)
}

// enabling TestFatal will cause test to fail with code 255
/*
func TestFatal(t *testing.T) {
	Fatal("Fatal: ", sampleMsg)
	Fatalln("Fatalln:", sampleMsg)
	Fatalf("Fatalf: %s, code: %d\n", sampleMsg, sampleCode)
}
*/

func TestMain(m *testing.M) {
	// configure logging
	logFlags := DefaultFlags()
	logFlags.Verbosity = Level(2)
	ApplyFlags(logFlags)

	code := m.Run()
	os.Exit(code)
}
