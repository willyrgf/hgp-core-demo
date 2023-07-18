# Hashicorp Go Plugin Core Demo


## Compile plugin
```sh
 go build  -o comp ./plugins/comp/main.go
```

## Compile core
```sh
 go build  -o core ./main.go
```

## Run test
```sh
PLUGIN="comp" ./core
```
```sh
2023-07-18T10:24:07.094-0300 [DEBUG] plugin: starting plugin: path=/bin/sh args=[sh, -c, ./comp]
2023-07-18T10:24:07.095-0300 [DEBUG] plugin: plugin started: path=/bin/sh pid=11843
2023-07-18T10:24:07.095-0300 [DEBUG] plugin: waiting for RPC address: path=/bin/sh
2023-07-18T10:24:07.104-0300 [DEBUG] plugin: using plugin: version=1
2023-07-18T10:24:07.104-0300 [DEBUG] plugin.sh: plugin address: network=unix address=/var/folders/7r/sy0t398s141_ps7tvyrsk8740000gn/T/plugin2682679746 timestamp=2023-07-18T10:24:07.104-0300
2023-07-18T10:24:07.105-0300 [TRACE] plugin.stdio: waiting for stdio data
Report:
message:"test OK"
2023-07-18T10:24:07.106-0300 [DEBUG] plugin.stdio: received EOF, stopping recv loop: err="rpc error: code = Unavailable desc = error reading from server: EOF"
2023-07-18T10:24:07.106-0300 [INFO]  plugin: plugin process exited: path=/bin/sh pid=11843
2023-07-18T10:24:07.106-0300 [DEBUG] plugin: plugin exited
```
