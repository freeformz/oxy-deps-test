package main

import (
	"fmt"

	"github.com/vulcand/oxy/buffer"
	"github.com/vulcand/oxy/cbreaker"
	"github.com/vulcand/oxy/connlimit"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/memmetrics"
	"github.com/vulcand/oxy/ratelimit"
	"github.com/vulcand/oxy/roundrobin"
	"github.com/vulcand/oxy/roundrobin/stickycookie"
	"github.com/vulcand/oxy/stream"
	"github.com/vulcand/oxy/trace"
)

func main() {
	var cl connlimit.ConnLimiter
	fmt.Println(cl)

	var buf buffer.Buffer
	fmt.Println(buf)

	var cb cbreaker.CircuitBreaker
	fmt.Println(cb)

	var fwd forward.Forwarder
	fmt.Println(fwd)

	var mm memmetrics.RatioCounter
	fmt.Println(mm)

	var rl ratelimit.RateErrHandler
	fmt.Println(rl)

	var rr roundrobin.RoundRobin
	fmt.Println(rr)

	var sc stickycookie.CookieValue
	fmt.Println(sc)

	var st stream.Stream
	fmt.Println(st)

	var tr trace.Record
	fmt.Println(tr)
}
