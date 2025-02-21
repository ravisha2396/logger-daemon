# Setting Up a BusyBox Docker Container with Rsyslog and gRPC Logging

## Overview

This document provides a step-by-step guide to setting up a **BusyBox-based Docker container** that includes `rsyslog` for system logging. Since BusyBox does not come with `rsyslog`, we use a **multi-stage Docker build** to install `rsyslog` on **Alpine Linux** and then copy the necessary binaries and libraries into the final BusyBox image.

Additionally, we include a **modularized Golang gRPC logging daemon** that listens for log messages, extracts the **client process name**, and logs it to `rsyslog` with the **client application's name** in syslog entries.

---

## Project Directory Structure

```plaintext
logging-project/
├── proto/
│   ├── logservice.proto
├── server/
│   ├── server.go
│   ├── server_test.go
├── client/
│   ├── client.go
├── docker/
│   ├── Dockerfile
│   ├── rsyslog.conf
├── build/
│   ├── logger-daemon (compiled binary)
│   ├── client-app (compiled binary)
├── Makefile
```

---


