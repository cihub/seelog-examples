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
	"fmt"
	log "github.com/cihub/seelog"
	"strings"
	"time"
)

var longMessage = strings.Repeat("A", 1024*100)

func main() {
	defer log.Flush()
	syncLogger()
	fmt.Println()
	asyncLoopLogger()
	fmt.Println()
	asyncTimerLogger()
}

func syncLogger() {
	fmt.Println("Sync test")

	testConfig := `
<seelog type="sync">
	<outputs>
		<filter levels="trace">
			<file path="log.log"/>
		</filter>
		<filter levels="debug">
			<console />
		</filter>
	</outputs>
</seelog>
`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)

	doTest()
}

func asyncLoopLogger() {
	fmt.Println("Async loop test")

	testConfig := `
<seelog>
	<outputs>
		<filter levels="trace">
			<file path="log.log"/>
		</filter>
		<filter levels="debug">
			<console />
		</filter>
	</outputs>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)

	doTest()

	time.Sleep(1e9)
}

func asyncTimerLogger() {
	fmt.Println("Async timer test")

	testConfig := `
<seelog type="asynctimer" asyncinterval="500">
	<outputs>
		<filter levels="trace">
			<file path="log.log"/>
		</filter>
		<filter levels="debug">
			<console />
		</filter>
	</outputs>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)

	doTest()

	time.Sleep(1e9)
}

func doTest() {
	start := time.Now()
	for i := 0; i < 30; i += 2 {
		fmt.Printf("%d\n", i)
		log.Trace(longMessage)
		log.Debug("%d", i+1)
	}
	end := time.Now()
	dur := end.Sub(start)
	fmt.Printf("Test took %d ns\n", dur)
}
