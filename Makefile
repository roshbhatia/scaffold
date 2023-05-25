.PHONY: create-kind-cluster delete-kind-cluster reset-kind-cluster run-starter run-worker

create-kind-cluster:
	kind create cluster --config dev/kind/cluster-config.yaml
	kind get kubeconfig > .kind-kubeconfig.yaml

delete-kind-cluster:
	kind delete cluster --name kind
	rm .kind-kubeconfig.yaml

reset-kind-cluster: delete-kind-cluster create-kind-cluster

run-starter:
	SCAFFOLD_CONFIG_PATH=dev/scaffold/std/config.yaml KUBECONFIG=.kind-kubeconfig.yaml go run ./src/cmd/std/starter/main.go

run-worker:
	SCAFFOLD_CONFIG_PATH=dev/scaffold/std/config.yaml KUBECONFIG=.kind-kubeconfig.yaml go run ./src/cmd/std/worker/main.go
