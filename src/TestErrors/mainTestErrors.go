package main

import (
	"fmt"
	"time"
	"math/rand"
	"ErrorHelper"
	"ErrorHelper/Core"
)

var tsProgrammStart time.Time = time.Now()

func init() {
	fmt.Printf("main.init() - initializing: %v ...\n", tsProgrammStart  )
}


func main() {
	rand.Seed( time.Now().UnixNano() )
	
	fmt.Printf("%v - TestErrors running ...\n\n", time.Now() )
	
	ErrorHelper.PublicErrorHelper()
	
	
	Core.CoreTest()
}

