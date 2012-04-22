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
	"time"
)

func outputsMain() {
	defer log.Flush()
	runExample(consoleWriter)
	runExample(fileWriter)
	runExample(rollingFileWriter)
	runExample(rollingFileWriterManyRolls)
	runExample(bufferedWriter)
	runExample(bufferedWriterWithFlushPeriod)
	runExample(bufferedWriterWithOverflow)
	runExample(splitDispatcher)
	runExample(filterDispatcher)
	//runExample(smtpWriter)
}

func runExample(exampleFunc func()) {
	exampleFunc()
	fmt.Println()
}

func consoleWriter() {
	testConfig := `
<seelog>
	<outputs>
		<console />
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Console writer")
	
	doLog()
}

func fileWriter() {
	
	testConfig := `
<seelog>
	<outputs>
		<file path="./log/log.log"/>
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("File writer")
	
	doLog()
}

func rollingFileWriter() {
	testConfig := `
<seelog>
	<outputs>
		<rollingfile type="size" filename="./log/roll.log" maxsize="100" maxrolls="5" />
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Rolling file writer")
	
	doLog()
}

func rollingFileWriterManyRolls() {
	testConfig := `
<seelog>
	<outputs>
		<rollingfile type="size" filename="./log/manyrolls.log" maxsize="100" maxrolls="4" />
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Rolling file writer. Many rolls")
	
	doLogBig()
}

func bufferedWriter() {
	testConfig := `
<seelog>
	<outputs>
		<buffered size="10000">
			<file path="./log/bufFile.log"/>
		</buffered>
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Buffered file writer. NOTE: file modification time not changed until next test (buffered)")
	time.Sleep(3e9)
	for i := 0; i < 3; i++ {
		doLog()	
		time.Sleep(5e9)
	}
	
	time.Sleep(2e9)
}

func bufferedWriterWithFlushPeriod() {
	testConfig := `
<seelog>
	<outputs>
		<buffered size="10000" flushperiod="1000">
			<file path="./log/bufFileFlush.log"/>
		</buffered>
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Buffered file writer with flush period. NOTE: file modification time changed after each 'doLog' because of small flush period.")
	time.Sleep(3e9)
	for i := 0; i < 3; i++ {
		doLog()	
		time.Sleep(5e9)
	}
	
	time.Sleep(2e9)
}

func bufferedWriterWithOverflow() {
	testConfig := `
<seelog>
	<outputs>
		<buffered size="20">
			<file path="./log/bufOverflow.log"/>
		</buffered>
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Buffered file writer with overflow. NOTE: file modification time changes after each 'doLog' because of overflow")
	time.Sleep(3e9)
	for i := 0; i < 3; i++ {
		doLog()	
		time.Sleep(5e9)
	}
	
	time.Sleep(1e9)
}


func splitDispatcher() {
	testConfig := `
<seelog>
	<outputs>
		<file path="./log/split.log"/>
		<console />
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Split dispatcher")
	
	doLog()	
}

func filterDispatcher() {
	testConfig := `
<seelog>
	<outputs>
		<filter levels="trace">
			<file path="./log/filter.log"/>
		</filter>
		<console />
	</outputs>
</seelog>
`
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	fmt.Println("Filter dispatcher")
	
	for i:=0; i < 5; i++ {
		log.Trace("This message on console and in file")
		log.Debug("This message only on console")
	}
}

func smtpWriter() {
 testConfig := `
 <seelog>
  <outputs>
   <smtp senderaddress="noreply-notification-service@none.org" sendername="Automatic notification service" hostname="mail.none.org" hostport="587" username="nns" password="123">
    <recipient address="john-smith@none.com"/>
   </smtp>
  </outputs>
 </seelog>
 `
 logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
 log.ReplaceLogger(logger)
 fmt.Println("SMTP writer is now sending emails to the specified recipients")
 doLog()
}

func doLog() {
	for i:=0; i < 5; i++ {
		log.Trace("%d", i)
	}
}

func doLogBig() {
	for i:=0; i < 50; i++ {
		log.Trace("%d", i)
	}
}
