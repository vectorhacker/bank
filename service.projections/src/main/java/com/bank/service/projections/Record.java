package com.bank.service.projections;

import org.json.JSONObject;

public abstract class Record {
    protected String type;

    public String getType() {
        return type;
    }

    public abstract JSONObject toJSON();
}