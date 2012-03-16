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
	"fmt"
)

func formatsMain() {
	defer log.Flush()
	defaultFormat()
	stdFormat()
	dateTimeFormat()
	dateTimeCustomFormat()
	logLevelTypesFormat()
	fileTypesFormat()
	funcFormat()
	xmlFormat()
}

func defaultFormat() {
	fmt.Println("Default format")
	
	testConfig := `
<seelog type="sync" />`

	logger, err := log.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		fmt.Println(err)
	}
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func stdFormat() {
	fmt.Println("Standard fast format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Ns [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func dateTimeFormat() {
	fmt.Println("Date time format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Date/%Time [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func dateTimeCustomFormat() {
	fmt.Println("Date time custom format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Date(2006 Jan 02/3:04:05.000000000 PM MST) [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func logLevelTypesFormat() {
	fmt.Println("Log level types format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Level %Lev %LEVEL %LEV %l %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func fileTypesFormat() {
	fmt.Println("File types format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%File %FullPath %RelFile %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func funcFormat() {
	fmt.Println("Func format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Func %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func xmlFormat() {
	fmt.Println("Xml format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="` +
		`&lt;log&gt;` +
			 `&lt;time&gt;%Ns&lt;/time&gt;` +
			 `&lt;lev&gt;%l&lt;/lev&gt;` +
			 `&lt;msg&gt;%Msg&lt;/msg&gt;` +
		 `&lt;/log&gt;"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))

	log.UseLogger(logger)
	
	log.Trace("Test message!")
}
