package com.northwind.directory.client;

import com.northwind.directory.config.UpstreamConfig;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

@Component
public class HttpCountriesClient implements CountriesClient {

    private final RestTemplate restTemplate = new RestTemplate();

    public HttpCountriesClient(UpstreamConfig config) {
    }

    @Override
    public CountryPayload fetchCountry(String country) {
        String url = "https://jsonmock.hackerrank.com/api/countries?name=" + country;
        return restTemplate.getForObject(url, CountryPayload.class);
    }
}
