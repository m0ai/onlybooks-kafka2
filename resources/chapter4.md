
# p132
```bash
 kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-topics.sh  \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --topic peter-test02 \
    --describe 
     
> Topic: peter-test02     TopicId: LxoZYmB_QRuILdGQYMONVw PartitionCount: 1       ReplicationFactor: 2    Configs: min.insync.replicas=1,message.format.version=3.0-IV1
>       Topic: peter-test02     Partition: 0    Leader: 1       Replicas: 1,0   Isr: 1,0
  
    1 broker = leader 
    0 broker = follower
    
```

```bash
# leader 가 1번 브로커이므로 1번 브로커에 접속해서 확인해보면 됩니다.
kubectl exec -it my-cluster-kafka-1 -c kafka \ 
    -- cat /var/lib/kafka/data-0/kafka-log1/peter-test02-0/leader-epoch-checkpoint   

0 
1     <- 현재의 leader epoch 번호
0 0   <- leader epoch 번호 , leader epoch start offset
```


# p134
```bash
# 메세지 프로듀서 실행
 kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-console-producer.sh \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --topic peter-test02 
> message 
```

# p135
```bash
# leader 인 1번 브로커 죽이기
kubectl delete pod my-cluster-kafka-1

# 새롭게 leader 승격된 0 번 브로커에 접속해서 확인
kubectl exec -it my-cluster-kafka-0 -c kafka -- cat /var/lib/kafka/data-0/kafka-log0/peter-test02-0/leader-epoch-checkpoint   

0
2     <- 현재 leader epoch 번호 1 에서 2로 증분됌 
0 0   <- 리더에포크 번호, 리더에포크 0의 시작 오프셋
1 1   <- 리더에포크 번호, 리더에포크 1의 시작 오프셋 (새로운 메세지를 전송 받게 될 번호)
```


# p141

```shell
kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-configs.sh \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --broker 1 \
    --describe --all | grep shutdown
    
> controlled.shutdown.enable=true sensitive=false synonyms={DEFAULT_CONFIG:controlled.shutdown.enable=true}
> controlled.shutdown.max.retries=3 sensitive=false synonyms={DEFAULT_CONFIG:controlled.shutdown.max.retries=3}
> controlled.shutdown.retry.backoff.ms=5000 sensitive=false synonyms={DEFAULT_CONFIG:controlled.shutdown.retry.backoff.ms=5000}

```

# p142
```shell
# cleanup.policy=delete 정책 확인
kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-configs.sh \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --broker 1 \
    --describe --all | grep cleanup 
> transaction.abort.timed.out.transaction.cleanup.interval.ms=10000 sensitive=false synonyms={DEFAULT_CONFIG:transaction.abort.timed.out.transaction.cleanup.interval.ms=10000}
> log.cleanup.policy=delete sensitive=false synonyms={DEFAULT_CONFIG:log.cleanup.policy=delete}
> transaction.remove.expired.transaction.cleanup.interval.ms=3600000 sensitive=false synonyms={DEFAULT_CONFIG:transaction.remove.expired.transaction.cleanup.interval.ms=3600000}

#  publishing
kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-console-producer.sh \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --topic peter-test03
> log1

# sbuscribing
kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-console-consumer.sh \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --topic peter-test03 --from-beginning
> log1
```

# p143
```shell
# topic 3 에 retention.ms=0 설정
kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-configs.sh \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --topic peter-test03 --add-config retention.ms=0 --alter
> Completed updating config for topic peter-test03.
    
# retention.ms=0 설정 확인
kubectl exec -it my-cluster-kafka-0 -c kafka -- \
    bin/kafka-topics.sh  \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --topic peter-test03 --describe
> Topic: peter-test03	TopicId: W6YDA02qSTqZXzZ9-8PyYQ	PartitionCount: 1	ReplicationFactor: 3	Configs: min.insync.replicas=1,retention.ms=0,message.format.version=3.0-IV1
>	Topic: peter-test03	Partition: 0	Leader: 2	Replicas: 2,1,0	Isr: 2,1,0 
```


# p144
```shell
# topic 3 에 retention.ms=0 삭제
kubectl exec -it my-cluster-kafka-0 -c kafka \
    -- bin/kafka-configs.sh \
    --bootstrap-server my-cluster-kafka-bootstrap:9092 \
    --topic peter-test03 --delete-config retention.ms --alter
> Completed updating config for topic peter-test03.
```
