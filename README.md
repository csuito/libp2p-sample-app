# Sample libp2p app

To run locally: ```go run main.go -p [port] [options]```

## Options

| Flag   | Type | Desc                         |
| ------ | ---- | ---------------------------- |
| -p     | int  | port                         |
| -d     | addr | target peer to dial          |
| -secio | bool | activate secure data streams |
| -seed  | int  | set random seed for id gen   |
