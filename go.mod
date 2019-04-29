module release_manager

require (
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/jmoiron/sqlx v1.2.0
	github.com/yeejlan/maru v0.0.0
)

replace github.com/yeejlan/maru => ../maru

go 1.12
