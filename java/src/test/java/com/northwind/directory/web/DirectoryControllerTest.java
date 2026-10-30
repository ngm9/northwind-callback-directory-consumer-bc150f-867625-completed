package com.northwind.directory.web;

import com.northwind.directory.client.CountriesClient;
import com.northwind.directory.client.CountryPayload;
import com.northwind.directory.service.DirectoryService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(DirectoryController.class)
class DirectoryControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @TestConfiguration
    static class StubConfig {
        @Bean
        CountriesClient countriesClient() {
            return country -> {
                CountryPayload.CountryRecord record = new CountryPayload.CountryRecord();
                record.setName("Afghanistan");
                record.setCallingCodes(List.of("93"));
                CountryPayload payload = new CountryPayload();
                payload.setData(List.of(record));
                return payload;
            };
        }

        @Bean
        DirectoryService directoryService(CountriesClient client) {
            return new DirectoryService(client);
        }
    }

    @Test
    void endpointReturnsOkForKnownCountry() throws Exception {
        mockMvc.perform(get("/phone-numbers")
                        .param("country", "Afghanistan")
                        .param("phone", "656445445"))
                .andExpect(status().isOk());
    }
}
