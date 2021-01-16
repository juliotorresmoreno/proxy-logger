package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/juliotorresmoreno/proxy-logger/config"
)

//Client .
type Client struct {
	*redis.Client
}

//Set .
func (c *Client) Set(key, value string) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	c.Client.Set(context, key, value, 24*time.Hour)
}

//Get .
func (c *Client) Get(key string) string {
	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	value := c.Client.Get(context, key).String()
	return value
}

//NewClient .
func NewClient() (*Client, error) {
	config, err := config.GetConfig()
	if err != nil {
		return &Client{}, err
	}
	url, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		return &Client{}, err
	}
	rdb := redis.NewClient(url)
	return &Client{rdb}, nil
}
