![LICENSE](https://img.shields.io/github/license/gsxhnd/owl-cli)
![Release](https://github.com/gsxhnd/owl-cli/workflows/Release/badge.svg?branch=v1.1.0)

## owl-cli

## Install

### Binary release
```
wget https://github.com/gsxhnd/owl-cli/releases/latest/download/owl-darwin-amd64
wget https://github.com/gsxhnd/owl-cli/releases/latest/download/owl-linux-amd64
wget https://github.com/gsxhnd/owl-cli/releases/latest/download/owl-windows-amd64.exe

mv owl-darwin-amd64 /usr/local/bin/owl
chmod +x /usr/local/bin/owl
```

### go install
```bash
go install github.com/gsxhnd/owl-cli/bin/owl
```

## Usage
```bash
usage: owl COMMAND [arg...]

commands:
   get  retrieve the value of a key
   put  set the value of a key
   version show version

```

### get
```bash
usage: owl get [flags] [arg...]
flags:
   -e, --endpoint string   etcd endpoint (default "http://127.0.0.1")
arg:
   the key what you want value at the etcd
```

### put
```bash
usage: owl put [flags] [arg...]
example:
    owl put /conf/test.yaml ../mock/test.yaml
flags:
    -e, --endpoint string   etcd endpoint (default "http://127.0.0.1")
arg:
    the key what you want value at the etcd
```

## Example
### get
```bash
owl get -e 127.0.0.1:2379 /conf/test.yaml
```
### put
```bash
owl put -e local_dev:2379 /conf/test.yaml ./mock/test.yaml
```