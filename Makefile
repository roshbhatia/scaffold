.PHONY: create-kind-cluster delete-kind-cluster reset-kind-cluster


create-kind-cluster:
	kind create cluster --config dev/kind/cluster-config.yaml

delete-kind-cluster:
	kind delete cluster --name kind

reset-kind-cluster: delete-kind-cluster create-kind-cluster
