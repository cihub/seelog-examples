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
	"time"
	log "github.com/cihub/seelog"
)

func adaptiveMain() {
	defer log.Flush()
	loadAdaptiveConfig()
	testMsgIntensity(1)
	testMsgIntensity(5)
	testMsgIntensity(10)
}

func testMsgIntensity(intensity int) {
	log.Default.Infof("Intensity test: %d", intensity)
	
	for j := 0; j < 4; j++ {
		for i := 0; i < intensity; i++ {
			log.Tracef("%d", i)
			<-time.After(time.Second / time.Duration(intensity))
		}
	}
	
	log.Default.Info("Messages sent")
	
	<-time.After(time.Second * time.Duration(intensity))
}

func loadAdaptiveConfig() {
	testConfig := `<seelog type="adaptive" mininterval="200000000" maxinterval="1000000000" critmsgcount="5">
	<outputs formatid="msg">
		<console/>
	</outputs>
	<formats>
		<format id="msg" format="%Time: %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))

	log.ReplaceLogger(logger)
}
