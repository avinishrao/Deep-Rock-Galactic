package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// Player represents a player's details
type Player struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Looks string `json:"looks"`
}

func main() {
	// Connect to Redis server
	rdb := connectToRedis()
	defer rdb.Close()

	// Prompt the user to input player details
	player := getPlayerDetails()

	// Store player data in Redis
	err := storePlayerData(rdb, player)
	if err != nil {
		fmt.Println("Error storing player data in Redis:", err)
		return
	}
	fmt.Println("Player details stored successfully in Redis")
}

func connectToRedis() *redis.Client {
	// Connect to Redis server
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Use default DB
	})

	// Ping the Redis server to check connectivity
	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil
	}
	fmt.Println("Connected to Redis")

	return rdb
}

func getPlayerDetails() Player {
	//Get player details
	var player Player
	fmt.Println("Enter player details:")
	fmt.Print("Name: ")
	fmt.Scanln(&player.Name)
	fmt.Print("Class(Engineer, Gunner, Scout, Driller): ")
	fmt.Scanln(&player.Class)
	fmt.Print("Looks(eg. Curly Hair & thin): ")
	fmt.Scanln(&player.Looks)
	return player
}

func storePlayerData(rdb *redis.Client, player Player) error {
	// Store player data in Redis
	playerJSON, err := json.Marshal(player)
	if err != nil {
		return err
	}
	err = rdb.Set(rdb.Context(), "player", string(playerJSON), 0).Err()
	if err != nil {
		return err
	}
	return nil
}
