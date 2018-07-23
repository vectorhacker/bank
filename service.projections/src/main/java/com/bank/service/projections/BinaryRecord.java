package com.bank.service.projections;

import org.json.JSONObject;

public class BinaryRecord extends Record {
    private byte payload[];

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