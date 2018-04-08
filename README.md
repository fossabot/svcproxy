# svcproxy
HTTP app-agnostic proxy allows to gather metrics and automatically issue certificates using ACME based CA, like Let's Encrypt

# Configuration example

svcproxy uses simple YAML configuration files like this working example:
```
---
listener:
  # Which port to listen for HTTP requests
  httpAddr: :8080
  # Which port to listen for HTTPS requests
  httpsAddr: :8443
autocert:
  cache:
    # Cache backend to use
    # Currently available:
    # - sql
    backend: sql
    backendOptions:
      # Driver to use by backend
      # Currently avaialble:
      # - mysql
      # - postgres
      driver: mysql
      # DSN(Data Source Name) to be passed to driver
      dsn: root@tcp(127.0.0.1:3306)/svcproxy
      # PSK(Pre-shared key) to encrypt/decrypt cached data
      encryptionKey: testkey
services:
  - frontend:
      # FQDN service is gonna response by
      fqdn: myservice.local
    backend:
      # Service backend to handle requests behind proxy
      url: http://localhost:8082
```

Some options could be passed as Environment variables:
 * `CONFIG_PATH` - path to YAML configuration file in file system

# TODO
 - [ ] Redirect from HTTP to HTTPS(configurable)
 - [ ] HTTPS-only service
 - [ ] Authentication(?)
 - [ ] Fix cache tests
