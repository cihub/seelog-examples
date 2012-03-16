// Copyright (c) 2012 - Cloud Instruments Co. Ltd.
// 
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	log "github.com/cihub/seelog"
	"github.com/cihub/seelog/test"
	"crypto/rand"
	"path/filepath"
	"math/big"
	"sync"
	"fmt"
	"os"
	"time"
)

const (
	LogDir = "log"
	goroutinesCount = 1000
	logsPerGoroutineCount = 100
	LogFile = "log.log"
)

var loggerReplacements int

var counterMutex *sync.Mutex
var waitGroup *sync.WaitGroup

var counter int64


var fileConfig = `
<seelog type="asyncloop">
	<outputs>
		<file path="` + filepath.Join(LogDir, LogFile) + `" formatid="testFormat"/>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileAsyncLoopConfig = `
<seelog type="asyncloop">
	<outputs>
		<file path="` + filepath.Join(LogDir, LogFile) + `" formatid="testFormat"/>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileAsyncTimer100Config = `
<seelog type="sync">
	<outputs>
		<file path="` + filepath.Join(LogDir, LogFile) + `" formatid="testFormat"/>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileAsyncTimer1000Config = `
<seelog type="asynctimer" asyncinterval="1000">
	<outputs>
		<file path="` + filepath.Join(LogDir, LogFile) + `" formatid="testFormat"/>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileAsyncTimer10000Config = `
<seelog type="asynctimer" asyncinterval="10000">
	<outputs>
		<file path="` + filepath.Join(LogDir, LogFile) + `" formatid="testFormat"/>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`



var fileBufferedConfig = `
<seelog type="sync">
	<outputs>
		<buffered size="100" formatid="testFormat">
			<file path="` + filepath.Join(LogDir, LogFile) + `"/>
		</buffered>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileBufferedAsyncLoopConfig = `
<seelog type="asyncloop">
	<outputs>
		<buffered size="100" formatid="testFormat">
			<file path="` + filepath.Join(LogDir, LogFile) + `"/>
		</buffered>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileBufferedAsyncTimer100Config = `
<seelog type="asynctimer" asyncinterval="100">
	<outputs>
		<buffered size="100" formatid="testFormat">
			<file path="` + filepath.Join(LogDir, LogFile) + `"/>
		</buffered>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileBufferedAsyncTimer1000Config = `
<seelog type="asynctimer" asyncinterval="1000">
	<outputs>
		<buffered size="100" formatid="testFormat">
			<file path="` + filepath.Join(LogDir, LogFile) + `"/>
		</buffered>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`

var fileBufferedAsyncTimer10000Config = `
<seelog type="asynctimer" asyncinterval="10000">
	<outputs>
		<buffered size="100"  formatid="testFormat">
			<file path="` + filepath.Join(LogDir, LogFile) + `"/>
		</buffered>
	</outputs>
	<formats>
	    <format id="testFormat" format="%Msg%n"/>
	</formats>
</seelog>`


var configPool = []string {
	fileConfig,
	fileAsyncLoopConfig,
	fileAsyncTimer100Config,
	fileAsyncTimer1000Config,
	fileAsyncTimer10000Config,
	fileBufferedConfig,
	fileBufferedAsyncLoopConfig,
	fileBufferedAsyncTimer100Config,
	fileBufferedAsyncTimer1000Config,
	fileBufferedAsyncTimer10000Config,
}

func switchToRandomConfigFromPool() {
	
	configIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(configPool))))
	
	if err != nil {
		panic(fmt.Sprintf("Error during random index generation: %s", err.Error()))
	}
	
	randomCfg := configPool[int(configIndex.Int64())]
	
	logger, err := log.LoggerFromConfigAsBytes([]byte(randomCfg))

	if err != nil {
		panic(fmt.Sprintf("Error during config creation: %s", err.Error()))
	}

	log.ReplaceLogger(logger)
	loggerReplacements++
}

func logRoutine(ind int) {
	for i := 0; i < logsPerGoroutineCount; i++ {
		counterMutex.Lock()
		log.Debug("%d", counter)
		//fmt.Printf("log #%v from #%v\n", i, ind)
		counter++
		switchToRandomConfigFromPool()
		counterMutex.Unlock()
	}
	
	waitGroup.Done()
}



func main() {
	os.Remove(filepath.Join(LogDir, LogFile))
	switchToRandomConfigFromPool()
	
	timeStart := time.Now()

	counterMutex = new(sync.Mutex)
	waitGroup = new(sync.WaitGroup)
	
	waitGroup.Add(goroutinesCount)
	
	for i := 0; i < goroutinesCount; i++ {
		go logRoutine(i)
	}
	
	waitGroup.Wait()
	log.Flush()
	
	timeEnd := time.Now()
	duration := timeEnd.Sub(timeStart)
	averageLoggerReplaceFrequency := float32(loggerReplacements) / (float32(duration.Nanoseconds()) / 1e9)

	gotCount, err := test.CountSequencedRowsInFile(filepath.Join(LogDir, LogFile))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Logger replaced %d times. Average replacement frequency: %f times / second. Output log is consistent: no log messages are missing or come in incorrect order.\n", loggerReplacements, averageLoggerReplaceFrequency)
	
	if counter == gotCount {
		fmt.Println("PASS! Output is valid")
	} else {
		fmt.Println("ERROR! Output not valid")
	}
}
