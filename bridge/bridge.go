package bridge

import (
	"github.com/echocat/slf4g"
	"github.com/echocat/slf4g-logr/bridge"
	nlogr "github.com/go-logr/logr"
	"k8s.io/klog/v2"
)

func Configure() {
	ConfigureWith(log.GetRootLogger())
}

func ConfigureWith(logger log.Logger) {
	bridge := logr.CreateFor(logger)
	bridge.CallerDepth = 4 + 1 // +1 for the workaround
	klog.SetLogger(nlogr.New(&workaround{bridge}))
}
