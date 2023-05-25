.PHONY: create-kind-cluster delete-kind-cluster reset-kind-cluster run-workflow

create-kind-cluster:
	kind create cluster --config dev/kind/cluster-config.yaml
	kind get kubeconfig > .kind-kubeconfig.yaml

delete-kind-cluster:
	kind delete cluster --name kind
	rm .kind-kubeconfig.yaml

reset-kind-cluster: delete-kind-cluster create-kind-cluster

run-workflow-std:
	KUBECONFIG=.kind-kubeconfig.yaml go run ./workflows/std/main.go
