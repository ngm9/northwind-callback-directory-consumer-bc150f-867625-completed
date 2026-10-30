package com.northwind.directory.client;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class CountryPayload {

    private List<CountryRecord> data;

    public List<CountryRecord> getData() {
        return data;
    }

    public void setData(List<CountryRecord> data) {
        this.data = data;
    }

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static class CountryRecord {
        private String name;
        private List<String> callingCodes;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public List<String> getCallingCodes() {
            return callingCodes;
        }

        public void setCallingCodes(List<String> callingCodes) {
            this.callingCodes = callingCodes;
        }
    }
}
