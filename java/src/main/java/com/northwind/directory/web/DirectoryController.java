package com.northwind.directory.web;

import com.northwind.directory.service.DirectoryService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DirectoryController {

    private final DirectoryService service;

    public DirectoryController(DirectoryService service) {
        this.service = service;
    }

    @GetMapping("/phone-numbers")
    public String phoneNumbers(@RequestParam String country, @RequestParam String phone) {
        return service.getPhoneNumbers(country, phone);
    }
}
