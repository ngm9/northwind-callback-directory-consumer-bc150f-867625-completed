package com.northwind.directory.service;

import com.northwind.directory.client.CountriesClient;
import com.northwind.directory.client.CountryPayload;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

class DirectoryServiceTest {

    private CountriesClient stubClient(CountryPayload payload) {
        return country -> payload;
    }

    private CountryPayload payloadWith(String name, List<String> codes) {
        CountryPayload.CountryRecord record = new CountryPayload.CountryRecord();
        record.setName(name);
        record.setCallingCodes(codes);
        CountryPayload payload = new CountryPayload();
        payload.setData(List.of(record));
        return payload;
    }

    @Test
    void resolvesSingleCallingCode() {
        DirectoryService service = new DirectoryService(
                stubClient(payloadWith("Afghanistan", List.of("93"))));
        assertEquals("+93 656445445", service.getPhoneNumbers("Afghanistan", "656445445"));
    }
}
