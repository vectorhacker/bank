package com.bank.service.projections;

import org.json.JSONObject;

public class BinaryRecord implements Record {
    private byte payload[];

    private String type;

    BinaryRecord(String type, byte payload[]) {
        this.payload = payload;
        this.type = type;
    }

    @Override
    public JSONObject toJSON() {
        return new JSONObject()
            .put("type", type)
            .put("payload", payload);
    }
}