package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	CITY     = "cities"
	AIRLINE  = "airlines"
	AGENCY   = "agencies"
	SUPPLIER = "suppliers"
)

type Database struct {
	gorm  *gorm.DB
	redis *redis.Client
}

func getDatabaseUri() string {
	var dbUser = os.Getenv("POSTGRES_USER")
	var dbPassword = os.Getenv("POSTGRES_PASSWORD")
	var db = os.Getenv("POSTGRES_DB")
	var dbHost = os.Getenv("POSTGRES_HOST")
	var dbPort = os.Getenv("POSTGRES_PORT")
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, dbPort, dbUser, db, dbPassword)
}

func getCacheUri() string {
	var cacheHost = os.Getenv("REDIS_HOST")
	var cachePort = os.Getenv("REDIS_PORT")
	return fmt.Sprintf("%s:%s", cacheHost, cachePort)
}

// New returns a connection to PostgreSQL as database and Redis for caching.
func New() (*Database, error) {

	var db Database
	var err error

	// Open PostgreSQL connection with gorm
	db.gorm, err = gorm.Open(postgres.Open(getDatabaseUri()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Make a redis client
	db.redis = redis.NewClient(&redis.Options{
		Addr:     getCacheUri(),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	if _, redis_err := db.redis.Ping().Result(); redis_err != nil {
		return nil, errors.New("unable to connect to redis")
	}

	err = db.initCache()
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized")

	return &db, nil
}

// getValidID returns id of the given name if exists
func (db *Database) getValidID(table string, name string) (int, error) {
	name = strings.ToUpper(name)
	id := db.redis.ZScore(table, name)

	if err := id.Err(); err == redis.Nil {
		return 0, fmt.Errorf("name %s not found", name)
	} else if err != nil {
		log.Println(err)
		return 0, errors.New("something went wrong")
	}
	return int(id.Val()), nil
}

// getID returns string id of the given name
func (db *Database) getID(table string, name string) string {
	id, _ := db.getValidID(table, name)
	return strconv.Itoa(id)
}
