
# Use Helm to install Strimzi
- helm 
- [strimzi-kafka-operator](https://artifacthub.io/packages/helm/strimzi/strimzi-kafka-operator/0.35.1)


ARM M1 환경에서 에뮬레이팅하다보니 init-kafka 를 실행할 때 init-kafka 에서 할당된 메모리 부족으로 OOM 이 발생한다. :(  
> https://github.com/orgs/strimzi/discussions/6553

스터디 실습 환경과 동일한게 좋지만 ARM M1 환경에서는 최신 카프카 버전을 사용하는 편이 환경 구성하는데 간편하므로 최신 버전을 사용한다.


```sh
# kind를 이용해 로컬 쿠버네티스 클러스터 환경을 구축한다.
 kind create cluster --image kindest/node:v1.27.3 --name local-1-27
 
# kind 클러스터를 사용하도록 kubectl 설정을 변경한다.
kubectl config use-context kind-local-1-27

# 카프카 클러스터를 생성해주기 위해 strimzi-kafka-operator 를 설치한다.
helm upgrade -i strimzi-kafka  \
strimzi/strimzi-kafka-operator \
--create-namespace \
--namespace operator \
--version 0.35.1 \
-f values.yaml

# 카프카 클러스터를 생성하기 위한 namespace 를 생성한다. 
kubectl -f resources/namespace.yaml

# kafka 클러스터를 생성한다.
kubectl -f resources/kafka.yaml
```

## Kafka Cluster 생성 확인
```sh
kubectl get pods -n kafka
```
# onlybooks-kafka2
