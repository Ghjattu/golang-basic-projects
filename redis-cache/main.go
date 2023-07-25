package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func AddSource(bytes []byte, source string) []byte {
	var obj map[string]interface{}
	json.Unmarshal(bytes, &obj)
	obj["source"] = source
	bytes, _ = json.Marshal(&obj)

	return bytes
}

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// check if we have a cached value
	val, err := rdb.Get(ctx, "characters").Result()
	// if we have a cache hit, return the cached value
	if err == nil {
		bytes := AddSource([]byte(val), "cache")
		fmt.Fprintln(w, string(bytes))
		return
	}

	// otherwise we are calling api for response
	res, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer res.Body.Close()

	bytes, _ := io.ReadAll(res.Body)
	rdb.Set(ctx, "characters", string(bytes), 10*time.Second)

	bytes = AddSource(bytes, "api")
	fmt.Fprintln(w, string(bytes))
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	http.HandleFunc("/characters", GetCharacters)
	http.ListenAndServe(":8080", nil)
}
