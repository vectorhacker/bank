package com.bank.service.projections;

import akka.actor.ActorSystem;
import eventstore.IndexedEvent;
import eventstore.EventData;
import eventstore.Content;
import eventstore.SubscriptionObserver;
import eventstore.j.EsConnection;
import eventstore.j.EsConnectionFactory;
import eventstore.Position;

import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;

import org.json.JSONObject;

import java.io.Closeable;
import java.lang.reflect.Method;
import java.nio.ByteBuffer;
import java.util.Properties;

public class Projections {
    public static void main(String[] args) {

        Properties props = new Properties();
        props.put("bootstrap.servers", "localhost:9092");
        props.put("acks", "all");
        props.put("retries", 1);
        props.put("batch.size", 16384);
        props.put("linger.ms", 1);
        props.put("buffer.memory", 33554432);
        props.put("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        props.put("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");

        Producer<String, String> producer = new KafkaProducer<>(props);

        final ActorSystem system = ActorSystem.create();
        final EsConnection connection = EsConnectionFactory.create(system);
        final Closeable closeable = connection.subscribeToAllFrom(new SubscriptionObserver<IndexedEvent>() {
            @Override
            public void onLiveProcessingStart(Closeable subscription) {
                system.log().info("live processing started");
            }

            @Override
            public void onEvent(IndexedEvent event, Closeable subscription) {
                final String streamId = event.event().streamId().streamId();
                final EventData data = event.event().data();
                final Content content = data.data();
                
                // skip streams that contain $
                if (streamId.contains("$")) {
                    system.log().info("skipping " + streamId);
                    return;
                }

                final String type = data.eventType();

                Record record = Projections.createRecord(type, content);

                if (record != null) {
                    String category = streamId.split("-")[0];
                    producer.send(new ProducerRecord<String, String>("ce-" + category, streamId,
                            record.toJSON().toString()));
                    producer.send(new ProducerRecord<String, String>("et-" + type, streamId,
                            record.toJSON().toString()));
                    producer.send(new ProducerRecord<String, String>(streamId,
                            streamId, record.toJSON().toString()));
                }
            }

            @Override
            public void onError(Throwable e) {
                system.log().error(e.toString());
            }

            @Override
            public void onClose() {
                system.log().error("subscription closed");
                producer.close();
            }
        }, Position.First(), false, null);
    }

    /**
     * Creates a new record based of the event type and content.
     * 
     * @param eventType the event's type
     * @param content   the content of the event
     * @returns A brand new record implementing Record interface
     */
    public static Record createRecord(String eventType, Content content) {
        if (content.contentType().toString() == "ContentType.Json") {
            final String payloadString = content.value().utf8String();
            final JSONObject jsonPayload = new JSONObject(payloadString);
            return new JSONRecord(eventType, jsonPayload);
        } else if (content.contentType().toString() == "ContentType.Binary") {
            ByteBuffer buffer = ByteBuffer.allocate(content.value().length());

            content.value().copyToBuffer(buffer);
            final byte bytePayload[] = buffer.array();

            return new BinaryRecord(eventType, bytePayload);
        }

        return null;
    }
}
