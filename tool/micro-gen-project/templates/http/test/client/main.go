package main

import (
	"context"
	"fmt"
	"github.com/gs02048/api/demo"
	"github.com/gs02048/micro/pkg/naming/etcd"
	"github.com/gs02048/micro/pkg/net/rpc/warden/resolver"
	"go.etcd.io/etcd/clientv3"
)

func init(){
	cfg := &clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
	}
	resolver.Register(etcd.Builder(cfg))
}
func main(){
	client,err := demo.NewClient(nil)
	if err != nil{
		panic(err)
	}
	resp,err := client.SayHelloURL(context.Background(),&demo.HelloReq{Name:"huanghailin"})
	if err != nil{
		panic(err)
	}
	fmt.Println(resp.Content)
}
