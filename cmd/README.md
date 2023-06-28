

```shell
kubectl run shell --rm -i --tty  --namespace kafka --image ubuntu -- bash

GOARCH=arm64 GOOS=linux go build -o producer-go cmd/transactional-producer.go

kubectl cp producer-go kafka/shell:/tmp/producer-go

kubectl exec -it shell --namespace kafka -- /tmp/producer-go
```
