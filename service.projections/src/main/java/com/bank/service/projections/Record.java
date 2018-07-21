package com.bank.service.projections;

import org.json.JSONObject;

public interface Record {
    public abstract JSONObject toJSON();
}