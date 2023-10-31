package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func connectDB() (*redis.Client, error) {
	addr, exists := os.LookupEnv("REDIS_ADDRESS")
	if !exists {
		return nil, errors.New("invalid redis address")
	}
	password, exists := os.LookupEnv("REDIS_ADDRESS")
	dbStr, exists := os.LookupEnv("REDIS_DB")
	db := 0
	if exists {
		dbInt, err := strconv.Atoi(dbStr)
		if err != nil {
			return nil, errors.New("invalid redis db")
		}
		db = dbInt

	}
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return client, nil
}
