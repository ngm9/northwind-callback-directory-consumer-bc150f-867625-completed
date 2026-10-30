# Northwind Support — Callback Directory Consumer

## Task Overview

Northwind Support agents log a customer's international callback number by supplying a country name and the customer's national phone number to an internal directory service. The service resolves the country's calling code from an upstream countries API and returns a formatted dialable number. The service already runs and its committed tests pass, but those tests only cover a single-word country with exactly one calling code — so real support tickets around multi-word countries, unknown countries, and malformed agent requests slip through unnoticed. Because agents rely on this service in real time while a customer is on the line, request construction, input handling, and the response contract must be correct and predictable.

This repository contains **four independent language folders** — `python/`, `javascript/`, `go/`, and `java/`. Each is a complete, self-contained, layered service. **Pick ONE folder and do all your work inside it.**

## Objectives

- Ensure a country name containing spaces reaches the upstream as a correctly-formed request, so multi-word countries resolve to their real calling code.
- Resolve territories that expose multiple calling codes onto the calling code the carrier expects, and keep the resolved format exactly `+<CallingCode> <PhoneNumber>`.
- Make an unknown country a successful response that carries the documented value, rather than a server error.
- Give malformed requests (missing or blank parameters) a clean, meaningful client-side response instead of an opaque server failure.
- Return the resolved value through a consistent, machine-parseable JSON response contract, and source the upstream base URL and request timeout from configuration.

## How to Verify

- Querying `Afghanistan` with `656445445` yields `+93 656445445`.
- Querying `Puerto Rico` with `564593986` yields `+1939 564593986` (the higher-indexed calling code wins).
- Querying `Oceania` with `987574876` returns `-1` on a successful response, not an error.
- A request missing or blanking a query parameter returns a clear client error rather than a 500.
- The committed tests cover all three sample cases at both the service layer and through the endpoint, and the whole suite runs offline without touching the network.

## Helpful Tips

- Consider how a value with a space must travel through a URL, and where in the layers that transformation belongs.
- Think about which layer should reject bad input and how a well-behaved API distinguishes a client mistake from a server fault.
- Explore how the resolution rules treat a list of calling codes with more than one entry, and what an empty upstream result should mean.
- Review how the upstream client is swapped out in the existing tests so you can keep new tests offline.
- Consider keeping the response shape uniform across success and not-found answers so consumers parse it the same way.
