# Globalping SDK

## Overview

Monitor, debug, and benchmark your internet infrastructure from a globally distributed network of probes.

## Client guidelines

If you implement an application interacting with the API, please consider the "client guidelines"
section of each endpoint below to provide the best UX and reduce the chance of your application breaking in the future.

### General guidelines for non-browser based apps:

- Set a `User-Agent` header. The recommended format and approach is [as here](https://github.com/jsdelivr/data.jsdelivr.com/blob/60c5154d26c403ba9dd403a8ddc5e42a31931f0d/config/default.js#L9).
- Set an `Accept-Encoding` header with a value of either `br` (preferred) or `gzip`, depending on what your client can support. The compression has a significant impact on the response size.
- When requesting the measurement status, implement ETag-based client-side caching using the `ETag`/`If-None-Match` headers.

## Endpoints

https://api.globalping.io/v1/


### Available Operations

