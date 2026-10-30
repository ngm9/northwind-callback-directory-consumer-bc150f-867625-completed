package com.northwind.directory.service;

import com.northwind.directory.client.CountriesClient;
import com.northwind.directory.client.CountryPayload;
import org.springframework.stereotype.Service;

@Service
public class DirectoryService {

    private final CountriesClient client;

    public DirectoryService(CountriesClient client) {
        this.client = client;
    }

    public String getPhoneNumbers(String country, String phone) {
        CountryPayload payload = client.fetchCountry(country);
        CountryPayload.CountryRecord record = payload.getData().get(0);
        String callingCode = record.getCallingCodes().get(0);
        return "+" + callingCode + " " + phone;
    }
}
