.PHONY: create-kind-cluster delete-kind-cluster reset-kind-cluster start-worker start-workflow

create-kind-cluster:
	kind create cluster --config dev/kind/cluster-config.yaml
	kind get kubeconfig > .kind-kubeconfig.yaml

delete-kind-cluster:
	kind delete cluster --name kind
	rm .kind-kubeconfig.yaml

reset-kind-cluster: delete-kind-cluster create-kind-cluster

start-worker:
	KUBECONFIG=.kind-kubeconfig.yaml go run ./workflows/std/worker/main.go

start-workflow:
	KUBECONFIG=.kind-kubeconfig.yaml go run ./workflows/std/starter/main.go
