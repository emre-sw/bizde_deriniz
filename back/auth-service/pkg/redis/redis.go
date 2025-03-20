package redis

import (
	"auth/pkg/configs"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() (*RedisClient, error) {
	config := configs.GetConfig()

	client := redis.NewClient(&redis.Options{
		Addr:         config.RedisHost + ":" + config.RedisPort,
		Password:     config.RedisPassword,
		DialTimeout:  10 * time.Second, // connection timeout
		ReadTimeout:  10 * time.Second, // read timeout
		WriteTimeout: 10 * time.Second, // write timeout
		PoolSize:     10,               // pool size
		MinIdleConns: 2,                // minimum idle connections
		DB:           config.RedisDB,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("redis error: ", err)
		return nil, err
	}

	log.Println("redis connected")
	return &RedisClient{Client: client}, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Println("redis set error: ", err)
	}

	return nil
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	result, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // if key not found, return empty string
	} else if err != nil {
		log.Println("redis get error: ", err)
		return "", err
	}

	return result, nil
}

func (r *RedisClient) Delete(ctx context.Context, keys ...string) error {
	err := r.Client.Del(ctx, keys...).Err()
	if err != nil {
		log.Println("redis delete error: ", err)
		return err
	}

	return nil
}

func (r *RedisClient) Close() error {
	log.Println("closing redis client")
	return r.Client.Close()
}

func (r *RedisClient) SetJSON(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON for Redis: %v", err)
		return err
	}
	return r.Set(ctx, key, jsonData, expiration)
}

func (r *RedisClient) GetJSON(ctx context.Context, key string, dest interface{}) error {
	jsonData, err := r.Get(ctx, key)
	if err != nil {
		return err
	}

	if jsonData == "" {
		return nil
	}

	err = json.Unmarshal([]byte(jsonData), dest)
	if err != nil {
		log.Printf("Failed to unmarshal JSON from Redis: %v", err)
		return err
	}
	return nil
}

func (r *RedisClient) BlacklistToken(ctx context.Context, token string, expiration time.Duration) error {
	err := r.Client.Set(ctx, "blacklist_tokens:"+token, "true", expiration).Err()
	if err != nil {
		log.Println("redis blacklist token error: ", err)
		return err
	}
	return nil
}

func (r *RedisClient) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	exists, err := r.Client.Get(ctx, "blacklist_tokens:"+token).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		log.Println("redis is token blacklisted error: ", err)
		return false, err
	}
	return exists == "true", nil
}

func (r *RedisClient) IsIPBanned(ctx context.Context, ip string) (bool, error) {
	exists, err := r.Client.Get(ctx, "banned_ip:"+ip).Result()
	if err == redis.Nil {
		return false, nil // if ip not found, return false
	} else if err != nil {
		log.Println("redis is ip banned error: ", err)
		return false, err
	}
	return exists == "true", nil
}

func (r *RedisClient) IncrementLoginAttempts(ctx context.Context, ip string) error {
	log.Println(31)
	config, err := configs.LoadConfig()
	if err != nil {
		log.Println("redis increment login attempts error: ", err)
		return err
	}

	failedAttempts, err := r.Client.Incr(ctx, "failed_attempts:"+ip).Result()
	if err != nil {
		log.Println("redis increment login attempts error: ", err)
		return err
	}
	log.Println("failed attempts: ", failedAttempts)
	if failedAttempts == 3 {
		r.Client.Expire(ctx, "failed_attempts:"+ip, config.LoginAttemptsTime)
	}

	if failedAttempts >= config.MaxLoginAttempts {
		r.Client.Set(ctx, "banned_ip:"+ip, "true", config.LoginAttemptsTime)
		log.Println("ip banned: ", ip)
	}

	return nil
}

func (r *RedisClient) ResetLoginAttempts(ctx context.Context, ip string) error {
	err := r.Client.Del(ctx, "failed_attempts:"+ip).Err()
	if err != nil {
		log.Println("redis reset login attempts error: ", err)
		return err
	}
	return nil
}
