module github.com/sql668/go-toolkit/contrib/storage/cache

go 1.22.0

require (
	github.com/redis/go-redis/v9 v9.5.1
	github.com/sql668/go-toolkit v0.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/spf13/cast v1.6.0 // indirect
)

replace github.com/sql668/go-toolkit => ../../../
