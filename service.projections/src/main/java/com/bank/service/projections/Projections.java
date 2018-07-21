package com.bank.service.projections;

import akka.actor.ActorSystem;
import eventstore.IndexedEvent;
import eventstore.EventData;
import eventstore.Content;
import eventstore.SubscriptionObserver;
import eventstore.j.EsConnection;
import eventstore.j.EsConnectionFactory;

import org.json.JSONObject;

import java.io.Closeable;
import java.lang.reflect.Method;
import java.nio.ByteBuffer;

public class Projections {
    public static void main(String[] args) {
        final ActorSystem system = ActorSystem.create();
        final EsConnection connection = EsConnectionFactory.create(system);
        final Closeable closeable = connection.subscribeToAll(new SubscriptionObserver<IndexedEvent>() {
            @Override
            public void onLiveProcessingStart(Closeable subscription) {
                system.log().info("live processing started");
            }

            @Override
            public void onEvent(IndexedEvent event, Closeable subscription) {
                final EventData data = event.event().data();
                final Content content = data.data();

                final String type = data.eventType();

                Record record = SubscribeToAllExample.createRecord(type, content);

                if (record != null) {
                    system.log().info(record.toJSON().toString());
                }
            }

            @Override
            public void onError(Throwable e) {
                system.log().error(e.toString());
            }

            @Override
            public void onClose() {
                system.log().error("subscription closed");
            }
        }, false, null);
    }

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
