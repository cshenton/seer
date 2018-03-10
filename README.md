# seer
Time series forecasting microservice

## Generating server snippets
```
# From this directory
protoc -I seer/ seer/seer.proto --go_out=plugins=grpc:seer
```