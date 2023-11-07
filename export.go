package myglog

func GetVerbosity() Level {
	return logging.verbosity.get()
}

func SetVerbosity(l Level) {
	logging.mu.Lock()
	defer logging.mu.Unlock()
	logging.setVState(l, logging.vmodule.filter, false)
}
