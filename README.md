![LICENSE](https://img.shields.io/github/license/gsxhnd/owl-cli)
![Release](https://github.com/gsxhnd/owl-cli/workflows/Release/badge.svg?branch=v1.1.0)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/gsxhnd/owl-cli?label=version)
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
NAME:
   owl - owl

USAGE:
   owl [global options] command [command options] [arguments...]

COMMANDS:
   get       get value by key
   get_keys  get keys by prefix
   put       read file then put value to etcd
   version   show version
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --endpoint value, -e value  (default: "http://127.0.0.1:2379")
   --help, -h                  show help (default: false)
```

### get
```bash
NAME:
   owl get - get value by key

USAGE:
   owl get [key]

DESCRIPTION:
   the [key] what you want value at the etcd
```

### put
```bash
NAME:
   owl put - read file then put value to etcd

USAGE:
   owl put [key] [file_path]

DESCRIPTION:
   example: owl put /conf/test.yaml ../mock/test.yaml

```

### get keys
```bash
NAME:
   owl get_keys - get keys by prefix

USAGE:
   owl get_keys [prefix]

```



## Example
### get
```bash
# get
owl get -e 127.0.0.1:2379 /conf/test.yaml
# put
owl put -e 127.0.0.1:2379 /conf/test.yaml ./mock/test.yaml
# get keys
owl put -e 127.0.0.1:2379 /conf/
```