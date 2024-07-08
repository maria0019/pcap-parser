# Go HTTP packets parser

CLI application for system packets parsing.

## Supported sources

1. File `.pcap`. Parser processes the file in an offline mode - when file is parsed then the script is finished.
2. Network interface. 
Data is processed in a live mode, the script captures packets going through the selected network interface.

## Supported protocols

Currently, the application supports the `HTTP` (`HTTP/1.1`) only.
Application runs with the protocol that set in the `.env` settings. 
The application can be extended with other protocols support:
- implement the protocol parser in the `internal/parser/`;
- set the needed protocol in the `.env` config for an appropriate environment.

## Input options

Application can be run to process data from one of the supported data sources:

- `file` - file path to parse, expected `.pcap` file;
- `net` - network interface name;
- `interval` - statistics calculation interval in seconds. Not required, the default value is `60`.

Either `file` or `net` options should be passed.

## Usage examples

### Parse a .pcap file

The `.pcap` file can be generated using `tcpdump` util:
```bash
sudo tcpdump port 80  -w www/ebpf/pcap-parser/test.pcap
```

Execute a few requests to sites that support the `HTTP/1.1`, as an example:
```bash
curl -i --raw http://captive.apple.com
```
```bash
curl -i --raw http://shop.paridon.com/
```

Run the app with `-file` and `-interval` option:

```bash
go run main.go -file test.pcap -interval 10
```

### Listen a network interface

Select network interface
```bash
ifconfig

...

wlp0s20f3: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500

...
```
and run the script:
```bash
sudo go run main.go -net wlp0s20f3 -interval 10
```

## The output
All statistics are passed to the Stdout.

File:
```bash
go run main.go -file test.pcap -interval 10

2024/07/08T19:58:45.004 channel=application level=INFO [file=main.go:12,main] message=Parser app started data=
2024/07/08T19:58:45.004 channel=application level=INFO [file=main.go:27,main] message=Run parser data={filePath:test.pcap, netInterface:, metricsInterval:10, protocol:HTTP}
2024/07/08T19:58:45.004 channel=application level=INFO [file=log.go:8,ToStdout] message=Stats data={at:2024-07-05 12:15:45, avgResponseTimeMs:65, requestPerUrl:[{"url":"captive.apple.com","count":11},{"url":"shop.paridon.com","count":18}], hasData:true}
2024/07/08T19:58:45.005 channel=application level=INFO [file=log.go:8,ToStdout] message=Stats data={at:2024-07-05 12:15:55, avgResponseTimeMs:74, requestPerUrl:[{"url":"captive.apple.com","count":6},{"url":"shop.paridon.com","count":6}], hasData:true}
2024/07/08T19:58:45.005 channel=application level=INFO [file=log.go:8,ToStdout] message=Stats data={at:2024-07-05 12:16:05, avgResponseTimeMs:0, requestPerUrl:[], hasData:false}
2024/07/08T19:58:45.005 channel=application level=INFO [file=log.go:8,ToStdout] message=Stats data={at:2024-07-05 12:16:15, avgResponseTimeMs:0, requestPerUrl:[], hasData:false}

```

Network interface:
```bash
sudo go run main.go -net wlp0s20f3 -interval 10

2024/07/08T19:59:10.467 channel=application level=INFO [file=main.go:12,main] message=Parser app started data=
2024/07/08T19:59:10.467 channel=application level=INFO [file=main.go:27,main] message=Run parser data={protocol:HTTP, filePath:, netInterface:wlp0s20f3, metricsInterval:10}
2024/07/08T19:59:20.474 channel=application level=INFO [file=log.go:8,ToStdout] message=Stats data={at:2024-07-08 19:59:20, avgResponseTimeMs:82, requestPerUrl:[{"url":"shop.paridon.com","count":12},{"url":"captive.apple.com","count":12}], hasData:true}
2024/07/08T19:59:30.476 channel=application level=INFO [file=log.go:8,ToStdout] message=Stats data={at:2024-07-08 19:59:30, avgResponseTimeMs:66, requestPerUrl:[{"url":"captive.apple.com","count":7},{"url":"shop.paridon.com","count":5}], hasData:true}
2024/07/08T19:59:40.476 channel=application level=INFO [file=log.go:8,ToStdout] message=Stats data={at:2024-07-08 19:59:40, avgResponseTimeMs:80, requestPerUrl:[{"url":"shop.paridon.com","count":1}], hasData:true}

```
