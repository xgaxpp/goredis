package cmd

import (
	"github.com/go-redis/redis"
)

func resolvingCmd(args []string) {
	client := redisCli["cluster"]
	if clusterCli, ok := (client).(redis.Client); ok {

	}
}

var redisCli = map[string]interface{}{
	"cluster": *redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"123123", "123123"},
		Password: "",
	}),
	"single": *redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
	}),
	"sentinel": *redis.NewSentinelClient(&redis.Options{
		Addr:     "",
		Password: "",
	}),
}
