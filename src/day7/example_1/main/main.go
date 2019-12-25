package main

import (
	"day7/example_1/balance"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port := rand.Intn(255)
		item := balance.NewInstance(host, port)
		insts = append(insts, item)
	}

	var balanceName = "hashBalance"
	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			return
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
