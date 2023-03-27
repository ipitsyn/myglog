// Go support for leveled logs, analogous to https://code.google.com/p/google-glog/
//
// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Flags structure

package myglog

type Flags struct {
	ToStderr        bool     // log to standard error instead of files
	AlsoToStderr    bool     // log to standard error as well as files
	Verbosity       Level    // log level for V logs
	StderrThreshold severity // logs at or above this threshold go to stderr
	VModule         string   // comma-separated list of pattern=N settings for file-filtered logging
	TraceLocation   string   // when logging hits line file:N, emit a stack trace
	LogDir          string   // directory in which to write logs
}

func DefaultFlags() *Flags {
	return &Flags{
		ToStderr:        true,
		AlsoToStderr:    false,
		Verbosity:       0,
		StderrThreshold: ErrorLog,
	}
}

func ApplyFlags(f *Flags) {
	logging.toStderr = f.ToStderr
	logging.alsoToStderr = f.AlsoToStderr

	logging.stderrThreshold.set(f.StderrThreshold)
	logging.vmodule.set(f.VModule)
	logging.traceLocation.set(f.TraceLocation)

	logging.mu.Lock()
	defer logging.mu.Unlock()
	logging.setVState(f.Verbosity, logging.vmodule.filter, false)

	logDir = f.LogDir
}
