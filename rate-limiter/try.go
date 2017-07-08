// example of rate limiting from golang wiki 
// https://github.com/golang/go/wiki/RateLimiting
// for higher precision and general usuage of rate limiter use golang.org/x/time/rate.Limiter instead

package main

import (
	"time"
	"fmt"
)

func main() {

	var reqnum int32

	rate := time.Second / 1
	burstLimit := 10
	tick := time.NewTicker(rate)
	defer tick.Stop()
	throttle := make(chan time.Time, burstLimit)
	go func() {
		for t := range tick.C {
			select {
			case throttle <- t:
			default:
			}
		} // exits after tick.Stop()
	}()
	time.Sleep(time.Duration(11 * time.Second))
	for {
		<-throttle // rate limit our Service.Method RPCs
		fmt.Println("sending the client request: ", reqnum)
		reqnum++
		//go client.Call("Service.Method", req, ...)
	}
}

