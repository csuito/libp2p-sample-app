# Sample libp2p app

To run locally: ```go run main.go -p [port] [options]```

## Options

| Flag   | Type | Desc                         |
| ------ | ---- | ---------------------------- |
| -p     | int  | port                         |
| -d     | addr | target peer to dial          |
| -secio | bool | activate secure data streams |
| -seed  | int  | set random seed for id gen   |

Based on: [Code a simple p2p blockchain](https://medium.com/@mycoralhealth/code-a-simple-p2p-blockchain-in-go-46662601f417)