package com.bank.service.projections;

import org.json.JSONObject;

public class JSONRecord extends Record {
    private JSONObject payload;

    JSONRecord(String type, JSONObject payload) {
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