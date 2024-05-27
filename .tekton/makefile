
all:
	kubectl apply -f pipeline/clone-build-push.yaml
	kubectl apply -f task/choose-dockerfile.yaml
	kubectl apply -f task/docker-build.yaml
	kubectl apply -f task/image-name.yaml
	kubectl apply -f task/image-tag.yaml
	kubectl apply -f task/milvus-env.yaml

clean:
	kubectl delete -f pipeline/clone-build-push.yaml
	kubectl delete -f task/choose-dockerfile.yaml
	kubectl delete -f task/docker-build.yaml
	kubectl delete -f task/image-name.yaml
	kubectl delete -f task/image-tag.yaml
	kubectl delete -f task/milvus-env.yaml
