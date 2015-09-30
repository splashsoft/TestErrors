package Core

import (
	"fmt"
	"time"
	"os"
	"log"
	"strings"
	"math/rand"
	
)

type MyLog struct {
	coreLogger *log.Logger
	currMinute int
	currFile string
	fhLogFile *os.File	
}


func (l *MyLog)Printf( strFmt string, args ...interface{}) error {
	
	// TODO: check, if day changed and open file accordingly
	t := time.Now()
	if t.Minute() != l.currMinute {
		if nil != l.fhLogFile {
			l.fhLogFile.Close()
		}
		
	fmt.Printf("MyLog::Printf - change log file - currMinute=%d - minute=%d ...\n", l.currMinute, t.Minute() )
	
	l.currFile = "" 	
	l.currMinute = t.Minute()
	l.coreLogger = nil
	} 

	var err error
	// do i have to (re)open the file 	
	if nil == l.coreLogger {
		s := fmt.Sprintf(".\\CoreLog%d%02d%02d_%02d.log", t.Year(), t.Month(), t.Day(), t.Minute() )
		
		
		fmt.Printf("MyLog::Printf - need to open log file '%s' ...\n", l.currFile )

		l.fhLogFile,err = os.OpenFile( s, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666);
		if err == nil {
			l.currFile = s
			l.coreLogger = log.New(l.fhLogFile,  "", log.Ltime | log.Lmicroseconds )
			} else {
				err = fmt.Errorf("MyLog: failed to open log file '%s' - %v!", s, err )
			}
	}
	
	if nil != l.coreLogger {
		err = nil
		if !strings.HasSuffix(strFmt, "\n")  {
			strFmt += "\n"	
			}
		l.coreLogger.Printf( strFmt, args )
	}
	
	return err
}

var coreLogging = MyLog{currMinute:-1}


func init() {
	// Kenne ich hier mein übergeordnetes Package?
	fmt.Printf("Core.init() - '%s' - initializing: %v ...\n", os.Getenv("USERNAME"), time.Now())
	log.SetOutput( os.Stdout)
	log.SetFlags( log.Ldate | log.Ltime | log.Lmicroseconds ) // Mikrosekunden Auflösung
	
	// log.SetPrefix("ErrorHelper/Core: ") - doof, das kommt noch vor  Datum/Uhrzeit
	
	// wann wird eine defered-Funktion der init-Funktion aufgerufen? // direkt nach Ende init(), also nicht Ende Programm (wäre irgendwie cool)
	
	// hier brauche ich eine Datei
	err := coreLogging.Printf("Core.init() - '%s' - initializing: %v ...", os.Getenv("USERNAME"), time.Now())
	if nil != err {
		fmt.Printf("logging failed - '%v' ...\n", err)
	}
	
	defer func() {
		fmt.Printf("das ist deferred Core.init() ...\n")
	}()
}


func CoreTest() {
	
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Printf("recovered from: %v ...\n", err )
//		}
//	}()
	
	fmt.Printf("ErrorHelper/CoreTest executing ... \n")
	nMax := 30
	for i:=0; i<nMax; i++ {
		fmt.Printf("ErrorHelper/CoreTest - working - %d of %d ...\n", i+1, nMax)
		
		coreLogging.Printf("ErrorHelper/CoreTest - run:%d ...", i )
		time.Sleep( (1 +time.Duration(rand.Intn(8)))*time.Second )
	}
	
	// log.Fatalf("ErrorHelper/CoreTest - unrecoverable error: %v ...\n", fmt.Errorf("This is a fake error") ) // Fatal kann nicht recovered werden 
	log.Panic(fmt.Errorf("This is a fake error") )
	//panic( fmt.Errorf("This is a fake error") )
	
	fmt.Printf("ErrorHelper/CoreTest ... done.\n")
	
}