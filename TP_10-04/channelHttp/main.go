package main

import (
	"fmt"
	myFunc "channelHttp/callServer"
)

func main() {
	ch0 := make(chan myFunc.Response)
	ch1 := make(chan myFunc.Response)

	go myFunc.CallServer("http://localhost:8000/handle?id=id0", ch0)
	go myFunc.CallServer("http://localhost:8000/handle?id=id1", ch1)

	select {
		case res := <- ch0:
            fmt.Println("channel 0:", res.RespText)
        case res := <- ch1: //faster endpoint
            fmt.Println("channel 1:", res.RespText)
	}
}