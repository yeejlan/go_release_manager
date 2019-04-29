package dal

import (
	"github.com/yeejlan/maru"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

var (
	DB *DBHelper
	DBReleaseManager *DBHelper

	Redis *redis.Client
)

func InitSharedVars() {
	db := maru.Registry.Get("db.release_manager").(*sqlx.DB)
	DB = NewDBHelper(db)
	DBReleaseManager = NewDBHelper(db)

	redis := maru.Registry.Get("redis.default").(*redis.Client)
	Redis = redis
}