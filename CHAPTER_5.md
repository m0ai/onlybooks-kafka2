# p165

```shell
kubectl apply -f ./resources/topic-peter-test04.yaml
```


```shell
kubectl cp ./resources/producer.config  kafka/my-cluster-kafka-0:/tmp/producer.config   

# 잘못된 옵션 파일을 가지고 producer 실행 
kubectl exec -it my-cluster-kafka-0 -c kafka -- bin/kafka-console-producer.sh \
  --bootstrap-server my-cluster-kafka-bootstrap:9092 \
  --topic peter-test04 \
  --producer.config /tmp/producer.config
    
```
최신 카프카에선 실행 에러 대신 메세지를 보낼 때 아래와 동일한 메세지가 출력됌

> [2023-06-28 14:39:58,824] WARN [Producer clientId=console-producer] Error while fetching metadata with correlation id 4 : {peter-test04=LEADER_NOT_AVAILABLE} (org.apache.kafka.clients.NetworkClient)
> [2023-06-28 14:39:58,926] WARN [Producer clientId=console-producer] Error while fetching metadata with correlation id 5 : {peter-test04=LEADER_NOT_AVAILABLE} (org.apache.kafka.clients.NetworkClient)
> 

# p166
메세제를 정상적으로 보낸 snapshot 파일이 없다면 브로커를 재시작하면 됌

```shell
kubectl exec -it my-cluster-kafka-0 -c kafka -- bin/kafka-dump-log.sh \
  --print-data-log \
  --files /var/lib/kafka/data-0/kafka-log0/peter-test04-0/00000000000000000004.snapshot
  
# producerId: 2000 producerEpoch: 0 coordinatorEpoch: -1 currentTxnFirstOffset: None lastTimestamp: 1687963199031 firstSequence: 0 lastSequence: 0 lastOffset: 0 offsetDelta: 0 timestamp: 1687963199031
# producerId: 2001 producerEpoch: 0 coordinatorEpoch: -1 currentTxnFirstOffset: None lastTimestamp: 1687963521642 firstSequence: 2 lastSequence: 2 lastOffset: 3 offsetDelta: 0 timestamp: 1687963521642

```


# p175

```shell
kubectl apply -f ./resources/topic-peter-test05.yaml
```
