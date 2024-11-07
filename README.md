# host-updater

The goal of this project is to update a host file to update the selected IPs when the network changes.

## build

```
go build -o bin/host-updater.exe
```

## Run

to automatically select a wifi network if available :

```
bin\host-updater.exe update -w
```
