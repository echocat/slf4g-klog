package bridge

import (
	"github.com/echocat/slf4g-logr/bridge"
)

type workaround struct {
	*logr.Bridge
}

func (instance *workaround) Info(msg string, keysAndValues ...interface{}) {
	instance.Bridge.Info(msg, instance.fixKeysAndValuesIfRequired(keysAndValues)...)
}

func (instance *workaround) Error(err error, msg string, keysAndValues ...interface{}) {
	instance.Bridge.Error(err, msg, instance.fixKeysAndValuesIfRequired(keysAndValues)...)
}

func (instance *workaround) fixKeysAndValuesIfRequired(keysAndValues []interface{}) []interface{} {
	if len(keysAndValues) != 1 {
		return keysAndValues
	}
	if sub, ok := keysAndValues[0].([]interface{}); ok {
		return sub
	}
	return keysAndValues
}
