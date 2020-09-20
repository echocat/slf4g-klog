package bridge

import (
	"github.com/echocat/slf4g"
	"github.com/echocat/slf4g-logr/bridge"
	"k8s.io/klog/v2"
)

func Configure() {
	ConfigureWith(log.GetGlobalLogger())
}

func ConfigureWith(logger log.Logger) {
	klog.SetLogger(logr.CreateFor(logger))
}
