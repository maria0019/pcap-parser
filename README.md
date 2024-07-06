# Go HTTP packets parser

CLI application for system packets parsing.

## Supported sources

1. File `.pcap`. Parser processes the file in an offline mode - when file is parsed then the script is finished.
2. Network interface. 
Data is processed in a live mode, the script captures packets going through the selected network interface.

## Supported protocols

Currently, the application supports the `HTTP` (`HTTP/1.1`) only.
Application runs for a protocol, set in the `.env` settings. 
The application can be extended with other protocols support:
- implement the protocol parser in the `service/parser/`;
- set needed protocol in the `.env` config for an appropriate environment.

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
sudo tcpdump port 80  -w www/ebpf/pparse/test.pcap
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
All statistics are passed to the Stdout in format compatible with logs parsers.

File:
```bash
go run main.go -file test.pcap -interval 10

time="2024-07-05T16:22:20+03:00" level=info msg="Parser app started"
time="2024-07-05T16:22:20+03:00" level=info msg="Run parser" filePath=test.pcap metricsInterval=10 netInterface= protocol=HTTP
time="2024-07-05T16:22:20+03:00" level=info msg=Stats at="2024-07-05 12:15:45" avgResponseTimeMs=65 hasData=true requestPerUrl="[{\"url\":\"captive.apple.com\",\"count\":11},{\"url\":\"shop.paridon.com\",\"count\":18}]"
time="2024-07-05T16:22:20+03:00" level=info msg=Stats at="2024-07-05 12:15:55" avgResponseTimeMs=74 hasData=true requestPerUrl="[{\"url\":\"shop.paridon.com\",\"count\":6},{\"url\":\"captive.apple.com\",\"count\":6}]"
time="2024-07-05T16:22:20+03:00" level=info msg=Stats at="2024-07-05 12:16:05" avgResponseTimeMs=0 hasData=false requestPerUrl="[]"
time="2024-07-05T16:22:20+03:00" level=info msg=Stats at="2024-07-05 12:16:15" avgResponseTimeMs=0 hasData=false requestPerUrl="[]"
```
Network interface:
```bash
sudo go run main.go -net wlp0s20f3 -interval 10

time="2024-07-05T17:02:09+03:00" level=info msg="Parser app started"
time="2024-07-05T17:02:09+03:00" level=info msg="Run parser" filePath= metricsInterval=10 netInterface=wlp0s20f3 protocol=HTTP
time="2024-07-05T17:02:19+03:00" level=info msg=Stats at="2024-07-05 17:02:19" avgResponseTimeMs=65 hasData=true requestPerUrl="[{\"url\":\"shop.paridon.com\",\"count\":4},{\"url\":\"captive.apple.com\",\"count\":9}]"
time="2024-07-05T17:02:29+03:00" level=info msg=Stats at="2024-07-05 17:02:29" avgResponseTimeMs=84 hasData=true requestPerUrl="[{\"url\":\"shop.paridon.com\",\"count\":1}]"
time="2024-07-05T17:02:39+03:00" level=info msg=Stats at="2024-07-05 17:02:39" avgResponseTimeMs=93 hasData=true requestPerUrl="[{\"url\":\"shop.paridon.com\",\"count\":12}]"
time="2024-07-05T17:02:49+03:00" level=info msg=Stats at="2024-07-05 17:02:49" avgResponseTimeMs=95 hasData=true requestPerUrl="[{\"url\":\"shop.paridon.com\",\"count\":6}]"
time="2024-07-05T17:02:59+03:00" level=info msg=Stats at="2024-07-05 17:02:59" avgResponseTimeMs=62 hasData=true requestPerUrl="[{\"url\":\"shop.paridon.com\",\"count\":3},{\"url\":\"captive.apple.com\",\"count\":8}]"
```

