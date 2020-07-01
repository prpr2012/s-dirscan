package main

import (
	"fmt"
	"net/http"
	"s-dirscan/utils"
	"strings"
	"sync"
	"flag"
	"github.com/wonderivan/logger"

)


var wg sync.WaitGroup
var (
	host string
	dic string
	thread int
	help string
)
var url string
func main()  {
	flag.StringVar(&host,"h","","target")
	flag.StringVar(&dic,"d","dic/dic.txt","dic")
	flag.IntVar(&thread,"t",10,"goroutine num")
	flag.Parse()
	if !utils.Check(url){
		fmt.Println("plase input right url...")
		return
	}
	paths,err:= utils.ReadLines(dic)
	if err == nil{
		logger.SetLogger("config/log.json")
		length := len(paths)
		COROUTNUM := thread
		groupLength := length/COROUTNUM
		wg.Add(10)
		for i:=0;i<COROUTNUM;i++{
			go getpath(paths[i*groupLength:(i+1)*groupLength])
		}
		go getpath(paths[COROUTNUM*groupLength:])
	}else {
		flag.PrintDefaults()
		return
	}
	wg.Wait()
	logger.Info("Done!")
	}


func getpath(paths []string){
	for _,v := range paths{
		defer func() {
			if err := recover();err !=nil{
				logger.Error(err)

			}
		}()
		//fmt.Println(v)
		if strings.Index(v,"/") != 0{
			v = "/"+v

		}
		resp,err:= http.Head(url+string(v))

		if err != nil{
			panic(err)
		}
		if resp.StatusCode == 200{
			logger.Info(host+string(v),"-->",resp.StatusCode)
		}
		resp.Body.Close()

	}
	wg.Done()
}
