package main

import (
	"errors"
	_ "github.com/echocat/slf4g-klog/bridge/hook"
	_ "github.com/echocat/slf4g/native"
	"k8s.io/klog/v2"
)

func main() {
	klog.Error("foo", "bar")
	klog.InfoS("anInfoM\"essage", "foo", 1, "bar", 2)
	klog.ErrorS(errors.New("anError"), "anErrorMessage", "foo", 2, "bar", 2)
	klog.WarningDepth(2, "a", "b")
}
