## mqtt-test Is a tool for testing and benchmarking Mainflux IIoT MQTT messiging

[Mainflux](https://github.com/mainflux/mainflux) is modern, scalable, secure open source and patent-free IoT cloud platform written in Go.

It supports various network protocols (i.e. HTTP, MQTT, WebSocket, CoAP), but we are interested mostly in MQTT, the IoT mostly used protocol.

Performance are tested simulating real world cases, with multiple clients and messages over TLS so you can easly configure the tool regarding to your needs and measures the results. Thanks to goroutines you can simulate concurrent high load and huge number of clients and massive load.

Example, how much it will take for:
```
10 clients sends 100 messages with QOS 2 using TLS
```
Using this testing approach you can easly test Mainflux MQTT performance and your infrastructure, HA, scaling capabilities etc...

## Requirments
You must have one Mainflux Thing provisioned and connected to one Mainflux channel.

TBD

## Configuration
Tool is configurable trough the config.yaml file in the project root.
