module github.com/dfds/provider-confluent

go 1.16

require (
	github.com/crossplane/crossplane-runtime v0.15.0
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/crossplane/provider-template v0.0.0-20211117150009-765d3e591445
	github.com/google/uuid v1.3.0 // indirect
	github.com/pkg/errors v0.9.1
	go.dfds.cloud v0.1.3
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.2
)
