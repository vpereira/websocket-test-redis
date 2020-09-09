package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func NewRedisClient() {
	log.Printf("Connecting to Redis..")
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Can't connect to Redis, %s", err)
	}
	log.Println(pong, err)
	log.Printf("Connected to Redis!")
}

func InsertKV(key string, value string) error {
	return client.Set(context.Background(), key, value, 0).Err()
}

func UpdateKV(key string, value string) error {
	return InsertKV(key, value)
}

func FlushAllKeys() error {
	return client.FlushAll(context.Background()).Err()
}

func DeleteKV(key string) error {
	return client.Del(context.Background(), key).Err()
}

func GetKeys(pattern string) ([]string, error) {
	r := client.Keys(context.Background(), pattern)
	return r.Result()
}

func GetValue(key string) (string, error) {
	res := client.Get(context.Background(), key)
	return res.Result()
}

func GetValues(pattern string) ([]string, error) {
	r, _ := GetKeys(pattern)
	var result []string
	for _, element := range r {
		val, err := GetValue(element)
		if err != nil {
			return result, nil
		}
		result = append(result, val)
	}
	return result, nil
}

func GetAllKeysValues() (map[string]string, error) {
	mapKeysValues := make(map[string]string)

	allKeys, err := GetKeys("*")

	if err != nil {
		return mapKeysValues, err
	}

	for _, key := range allKeys {
		val, _ := GetValue(key)
		mapKeysValues[key] = val
	}
	return mapKeysValues, nil
}

func Subscribe(pattern string, stop <-chan struct{}) (<-chan *redis.Message, error) {
	ctx := context.Background()
	pubsub := client.PSubscribe(ctx, pattern)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return nil, err
	}
	go func() {
		<-stop
		pubsub.Close()
	}()
	return pubsub.Channel(), nil
}

func Publish(channel string, message interface{}) error {
	return client.Publish(context.Background(), channel, message).Err()
}

func IsNotFound(err error) bool {
	return err == redis.Nil
}
