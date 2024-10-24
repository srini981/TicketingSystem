package pkg

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"intelXlabs/dbService/internals"
	"log"
)

// dbclient to handle all db request
type client struct {
	PostgresClient *gorm.DB
	RedisClient    *redis.Client
}

var dbClient = InitDB()

// init function for postgres and redis clients
func InitDB() client {
	dbclient := client{}

	dsn := "host=localhost user=pg password=pass dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect postgres database")
		return client{}
	}

	db.AutoMigrate(internals.Ticket{})
	db.AutoMigrate(internals.Role{})
	db.AutoMigrate(internals.User{})
	postgresdb, err := db.DB()

	if err != nil {
		log.Fatal("failed to create postgres db object")
		return client{}
	}

	err = postgresdb.Ping()

	if err != nil {
		log.Fatal("failed to Ping postgres database")

		return client{}
	}

	dbclient.PostgresClient = db

	redisclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisclient.Ping().Result()

	if err != nil {
		log.Fatal("failed to Ping redis database")
		return client{}
	}
	
	dbclient.RedisClient = redisclient
	return dbclient

}
