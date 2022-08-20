module github.com/chaosblade-io/chaosblade-exec-os

go 1.13

require (
	github.com/chaosblade-io/chaosblade-spec-go v1.5.1-0.20220423030509-6d8dbd90b300
	github.com/containerd/cgroups v1.0.2-0.20210605143700-23b51209bf7b
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c
	github.com/shirou/gopsutil v3.21.6+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.7 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//TODO remove
replace (
	github.com/chaosblade-io/chaosblade-spec-go => ../chaosblade-spec-go
)
