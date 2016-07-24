package goslim

import (
	"flag"
	"fmt"
	"goslim/slimexecutor"
	"goslim/slimlist"
	"goslim/slimprotocol"
	"goslim/slimsocket"
	"log"
	"os"
)

func init() {
	initLog()
}

var (
	logFileName = flag.String("log", "GoSlim.log", "Goslim log file name")
)

func initLog() {
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "Goslim start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func WAIT_AND_RUN() {
	slimSocket.Instance().Listen()

	slimSocket.Instance().SendMsg(slimprotocol.VERSION())

	for {
		if requestMsg, err := slimSocket.Instance().ReceiveMsg(); err != nil {
			break
		} else {

			if slimprotocol.BYE() == requestMsg {
				break
			}

			responseMsg := handleMessage(requestMsg)
			slimSocket.Instance().SendMsg(fmt.Sprintf("%06d:", len(responseMsg)))
			slimSocket.Instance().SendMsg(responseMsg)
		}
	}
}

func handleMessage(message string) string {
	log.Println("handleMessage request: ", message)
	instructions := slimlist.SlimList_Deserialize(message)
	if instructions == nil {
		return slimprotocol.EXCEPTION("message nok")
	}

	results := slimexecutor.Instance().Execute(instructions)

	response := results.Serialize()
	log.Println("handleMessage response: ", response)
	return response
}

func RegisterFixture(className string, constructor slimexecutor.Constructor) {
	slimexecutor.Instance().RegisterFixture(className, constructor)
}

func RegisterMethod(className string, aliasMethodName string, realMethodName string) {
	slimexecutor.Instance().RegisterMethod(className, aliasMethodName, realMethodName)
}
