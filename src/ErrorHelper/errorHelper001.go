package ErrorHelper

import (
	"fmt"
	"time"
	"ErrorHelper/Core"
)

var tsStart time.Time = time.Now()


func init() {
	fmt.Printf("ErrorHelper.init() - initializing: %v ...\n", tsStart  )
}


func privateErrorHelper() {
	fmt.Printf("privateErrorHelper - after:%v ...\n", time.Since(tsStart)  )
}


func PublicErrorHelper() {
	
	
	// kann ich einen panic in CoreTest() auch heir abfangen (eine Ebene höher)?
	// yesssss ...
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("recovered from: %v ...\n", err )
		}
	}()

	fmt.Printf("PublicErrorHelper - executing, after:%v ...\n", time.Since(tsStart)  )
	
	Core.CoreTest()
	
	fmt.Printf("PublicErrorHelper - ... done.\n"  )

	
}
