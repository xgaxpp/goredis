package cmd

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
)

func resolvingCmd(args []string) {
	client := redisCli["cluster"]
	object := reflect.ValueOf(client)

	fmt.Println(object.NumMethod())

	//v:= object.MethodByName("HGet")
	//in := v.Type().NumIn()
	//out := v.Type().NumOut()
	//fmt.Println(in, out)
	//var param []reflect.Value
	////param[0] = reflect.Value{"", "", ""}
	//fmt.Println(v.Call(param))
}

var redisCli = map[string]interface{}{
	"cluster": redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{},
	}),
	"single": redis.NewClient(&redis.Options{
		Addr: "",
	}),
	"sentinel": redis.NewSentinelClient(&redis.Options{
		Addr: "",
	}),
}
