module github.com/zjedens/http_server_chirpy

go 1.26.7

replace github.com/zjedens/api v0.0.0 => ./api

require github.com/zjedens/api v0.0.0

require github.com/lib/pq v1.6.0

require (
	github.com/alexbrainman/sspi v0.0.0-20180613141037-e580b900e9f5 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/goidentity/v6 v6.0.1 // indirect
	github.com/jcmturner/gokrb5/v8 v8.2.0 // indirect
	github.com/jcmturner/rpc/v2 v2.0.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4 // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
)
