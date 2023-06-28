# Chapters
- [CHAPTER_4.md](CHAPTER_4.md)
- [CHAPTER_5.md](CHAPTER_5.md)

# Pre-requisites 
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installing-with-a-package-manager)
- [helm 3](https://helm.sh/ko/docs/intro/install/#%ED%8C%A8%ED%82%A4%EC%A7%80-%EB%A7%A4%EB%8B%88%EC%A0%80%EB%A5%BC-%ED%86%B5%ED%95%B4%EC%84%9C)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

ARM M1 환경에서 x86_64 이미지를 에뮬레이팅하다보니 init-kafka 를 실행할 때 메모리 부족으로 OOM 이 발생합니다. :(  
> https://github.com/orgs/strimzi/discussions/6553

스터디 실습 환경과 동일한게 좋지만 ARM M1 환경에서는 최신 카프카 버전을 사용하도록 구성하는데 간편하므로 최신 버전을 사용합니다.

이 문서에서 카프카 클러스터 구성은 [strimzi-kafka-operator](https://artifacthub.io/packages/helm/strimzi/strimzi-kafka-operator/0.35.1)을 사용합니다.

```sh
# kind를 이용해 로컬 쿠버네티스 클러스터 환경을 구축합니다.
kind create cluster --image kindest/node:v1.27.3 --name local-1-27
 
# kind 클러스터를 사용하도록 kubectl 설정을 변경합니다.
kubectl config use-context kind-local-1-27

# helm repo 를 추가합니다.
helm repo add strimzi https://strimzi.io/charts/

# 카프카 클러스터를 생성해주기 위해 strimzi-kafka-operator 를 설치합니다.
helm upgrade -i strimzi-kafka  \
strimzi/strimzi-kafka-operator \
--create-namespace \
--namespace operator \
--version 0.35.1 \
-f values.yaml

# 카프카 클러스터를 생성하기 위한 namespace 를 생성합니다.
kubectl -f resources/namespace.yaml

# kafka 클러스터를 생성합니다.
kubectl -f resources/kafka.yaml
```

## Kafka Cluster 생성 확인
```sh
kubectl get pods -n kafka
```
