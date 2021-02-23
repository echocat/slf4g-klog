module github.com/echocat/slf4g-bridge-klog/internal/demo

go 1.14

replace github.com/echocat/slf4g-klog => ../../

require (
	github.com/echocat/slf4g-klog v0.0.0
	github.com/echocat/slf4g/native v0.11.0
	k8s.io/klog/v2 v2.3.0
)
