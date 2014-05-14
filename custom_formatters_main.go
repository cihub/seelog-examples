// Copyright (c) 2014 - Cloud Instruments Co. Ltd.
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
	"os"
	"strconv"
	"time"
)

var hostName string
var appName = "test"
var pid int

func init() {
	hostName, _ = os.Hostname()
	pid = os.Getpid()
}

var myLevelToString = map[log.LogLevel]string{
	log.TraceLvl:    "MyTrace",
	log.DebugLvl:    "MyDebug",
	log.InfoLvl:     "MyInfo",
	log.WarnLvl:     "MyWarn",
	log.ErrorLvl:    "MyError",
	log.CriticalLvl: "MyCritical",
	log.Off:         "MyOff",
}

var levelToSyslogSeverity = map[log.LogLevel]int{
	// Mapping to RFC 5424 where possible
	log.TraceLvl:    7,
	log.DebugLvl:    7,
	log.InfoLvl:     6,
	log.WarnLvl:     5,
	log.ErrorLvl:    3,
	log.CriticalLvl: 2,
	log.Off:         7,
}

func createSyslogHeaderFormatter(params string) log.FormatterFunc {
	facility := 20
	i, err := strconv.Atoi(params)
	if err == nil && i >= 0 && i <= 23 {
		facility = i
	}

	return func(message string, level log.LogLevel, context log.LogContextInterface) interface{} {
		return fmt.Sprintf("<%d>1 %s %s %s %d - -", facility*8+levelToSyslogSeverity[level],
			time.Now().Format("2006-01-02T15:04:05Z07:00"),
			hostName, appName, pid)
	}
}

func createMyLevelFormatter(params string) log.FormatterFunc {
	return func(message string, level log.LogLevel, context log.LogContextInterface) interface{} {
		levelStr, ok := myLevelToString[level]
		if !ok {
			return "Broken level!"
		}
		return levelStr
	}
}

func init() {
	// Errors are omitted here for simplicity. Do not omit them in production.
	_ = log.RegisterCustomFormatter("CustomSyslogHeader", createSyslogHeaderFormatter)
	_ = log.RegisterCustomFormatter("MyLevel", createMyLevelFormatter)
}

func customFormattersMain() {
	levelsFormatter()
	syslogCustomFormatter()
}

func levelsFormatter() {
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%MyLevel %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}

func syslogCustomFormatter() {
	defer log.Flush()

	// NOTE: in real usecases a conn writer should be used instead of console.
	// Example: <conn formatid="syslogfmt" net="tcp4" addr="server.address:5514" tls="true" insecureskipverify="true" />
	testConfig := `
<seelog type="sync">
	<outputs formatid="syslogfmt">
		<console/>
	</outputs>
	<formats>
		<format id="syslogfmt" format="%CustomSyslogHeader(20) %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}
