module github.com/itsmontoya/bringyour-provider

go 1.22.3

replace bringyour.com/connect v0.0.0 => ../connect/connect

replace bringyour.com/protocol v0.0.0 => ../connect/protocol/build/bringyour.com/protocol

require (
	bringyour.com/connect v0.0.0
	bringyour.com/protocol v0.0.0
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/hatchify/closer v0.4.81
	github.com/vroomy/vroomy v0.17.3
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/gdbu/atoms v1.0.1 // indirect
	github.com/gdbu/bst v0.3.1 // indirect
	github.com/gdbu/queue v0.4.81 // indirect
	github.com/gdbu/reflectio v0.1.3 // indirect
	github.com/gdbu/scribe v0.5.3 // indirect
	github.com/gdbu/stringset v0.2.0 // indirect
	github.com/golang/glog v1.2.1 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hatchify/atoms v0.4.79 // indirect
	github.com/hatchify/colors v0.4.79 // indirect
	github.com/hatchify/cron v0.4.82 // indirect
	github.com/hatchify/errors v0.4.82 // indirect
	github.com/oklog/ulid/v2 v2.1.0 // indirect
	github.com/vroomy/httpserve v0.10.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
