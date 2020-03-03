package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	res, ok := ctx.Value("trace_id").(string)
	if !ok {
		res = "this is default value"
	}
	fmt.Printf("ret:%s\n", res)
	s, _ := ctx.Value("session").(string)
	fmt.Printf("session:%s\n", s)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", "req_12312312312")
	ctx = context.WithValue(ctx, "session", "fasdfdsfafadbm")
	process(ctx)
}
