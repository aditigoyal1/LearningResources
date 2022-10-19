# What is Kafka?
Kafka is an event streaming platform, it allows us to publish, subscribe, store and process events.

1. Features
2. Architecture
3. Key Teminologies
4. Installation and  Pre-requisties
5. Setting up single-node Kafka Cluster
6. Setting up Multi-node Kafka cluster
7. ZooKeeper
8. Setting Up single node zookeeper
9. Setting up Multi-node zookeeper
10. Producer Details
11. Consumer Details
12. Topic Creation and Alteration
13. Kafka Security
    PlainText
    Sasl_PlainText
    SSL(TLS encryption)
    SASL_SSl

14. Open Sorce Lib 
15. Cluster Monitoring
16. JMX exporter
17. Prometheus
18. Grafana


For zookeeper installation
1. Java 8 or above
2. RAM size minimum 512MB

Steps:
1. Untar Zookeeper Binary
2. Create file zoo.cfg within "conf" directory.
3. Write down the properties in zoo.cfg file.

        # The number of milliseconds of each tick
        tickTime=2000
        # The number of ticks that the initial 
        # synchronization phase can take
        initLimit=10
        # The number of ticks that can pass between 
        # sending a request and getting an acknowledgement
        syncLimit=5
        # the directory where the snapshot is stored.
        # do not use /tmp for storage, /tmp here is just 
        # example sakes.
        dataDir=/tmp/zookeeper
        # the port at which the clients will connect
        clientPort=2181
        # the maximum number of client connections.
        # increase this if you need to handle more clients
        maxClientCnxns=60
        4lw.commands.whitelist=*
4. Start Zookeeper

        bin/zkServer.sh start [In background]
        bin/zkServer.sh start-foreground [In foreground]


5. Validate whether Zookeeper server is running or not.
   
        echo status | nc localhost 2181


### Kafka Setup

1. Untar kafka binary 
2. Edit the server-properties within config directory
   
        broker.id=1
        listeners=PLAINTEXT://localhost:9092
3. start Kafka
   
        bin/kafka-server-start.sh config/server.properties [ IN-foreground]  
        bin/kafka-server-start.sh -daemon config/server.properties [ In Background ]
4. Validate Kafka server status
   
        echo dump |nc localhost 2181 | grep brokers

### Create Topic from command Line

1. Create topic "myTopic" 
     - Command: bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic  mytopic --partitions 1 --replication-factor 1
2. List the Topic 
    -  Command: bin/kafka-topics.sh --bootstrap-server localhost:9092 --list 
3. Describe the Topic
    -  Command: bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic mytopic
    -  OUTPUT: 
    -  Topic: mytopic	TopicId: GN7SKv6PSY60JrAHLgw90w	PartitionCount: 1	ReplicationFactor: 1	Configs: 
	Topic: mytopic	Partition: 0	Leader: 1	Replicas: 1	Isr: 1



### Working with console Producer
- bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic mytopic


### Working with console Consumer
- bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic mytopic --from-beginning

Every consumer is connected to consumer-group
### Commands to check topic properties and consumer group properties.
1. Check the consumer group
   - bin/kafka-consumer-groups.sh -bootstrap-server localhost:9092 --list
2. Describe the consumer group 
   - bin/kafka-consumer-groups.sh -bootstrap-server localhost:9092 --describe --group console-consumer-75948
  
Using consumer-offset, Kafka cluster stores the information of all consumers. 
One More topic is created by Kafka i.e __consumer_offsets
bin/kafka-topics.sh --bootstrap-server localhost:9092 --list 


     




