package main

import (
	"context"
	"time"

	"log"

	"github.com/xswitch-cn/proto/go/proto/cman"
	"github.com/xswitch-cn/proto/go/proto/xctrl"
	"github.com/xswitch-cn/proto/xctrl/client"
	ulog "github.com/xswitch-cn/proto/xctrl/util/log"
)

type Logger struct {
	ulog.Logger
}

func (t *Logger) Log(level int, v ...interface{}) {
	log.Println(v...)
}

func (t *Logger) Logf(level int, format string, v ...interface{}) {
	log.Printf(format, v...)
}

// simple example
func main() {
	logLevel := ulog.LevelDebug

	ulog.SetLevel(logLevel)     // set ctrl log level
	ulog.SetLogger(new(Logger)) // tell ctrl to use our logger
	log.Print("Hello, world!")  // the world starts from here

	c := client.NewFakeClient()
	c.Init(client.Selector())
	service := xctrl.NewXNodeService("fake", c)

	response, err := service.NativeAPI(context.Background(), &xctrl.NativeAPIRequest{
		Cmd: "status",
	}, client.WithAddress("cn.xswitch.node"), client.WithRequestTimeout(1*time.Second))

	if err != nil {
		panic(err)
	}

	log.Printf("response: %v", response.Data)

	cListReq := &xctrl.ConferenceListRequest{
		CtrlUuid: "fakeUUID",
		Data: &xctrl.ConferenceListRequestData{
			Command: "conferenceInfo",
			Data: &xctrl.ConferenceListRequestDataData{
				Domain: "",
			},
		},
	}
	rsp, err := service.ConferenceList(context.Background(), cListReq,
		client.WithAddress("cn.xswitch.node"), client.WithRequestTimeout(1*time.Second))
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%v", rsp.Data)
		for _, c := range rsp.Data {
			log.Printf("conference %s %s", c.ConferenceName, c.Domain)
		}
	}

	cManService := cman.NewCManService("fake", c)

	res, err := cManService.GetConferenceList(context.Background(), &cman.GetConferenceListRequest{},
		client.WithRequestTimeout(1*time.Second))
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("conferences %v", res.Conferences)
	}
}
