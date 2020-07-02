package main

import (
	"flag"
	"fmt"
	"github.com/wonderivan/logger"
	"net/http"
	"s-dirscan/utils"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var (
	host   string
	dic    string
	thread int
	sfile  string
)

func main() {
	fmt.Print(utils.ReadFile("banner.txt"))
	flag.StringVar(&host, "h", "", "target")
	flag.StringVar(&dic, "d", "dic/dic.txt", "dic")
	flag.IntVar(&thread, "t", 10, "goroutine num")
	flag.StringVar(&sfile,"f","","read hosts from specific file")
	flag.Parse()

	paths,err2 := utils.ReadLines(dic)
	if err2 !=nil{
		fmt.Println("打开文件出错:",err2)
		return
	}

	if sfile != ""{
		logger.SetLogger("config/log.json")
		hosts,err1 := utils.ReadLines(sfile)
		if err1 == nil{
			for _,host := range hosts{
				execTask(host,paths)
			}
		}else {
			fmt.Println("打开文件出错:",err1)
		}

	}else if host !=""{
		execTask(host,paths)
	} else {
		flag.PrintDefaults()
	}

}

func execTask(host string,paths []string){
	if host!=""&&!utils.Check(host) {
		fmt.Println("plase input right url or target is dead...")
		return
	}
	length := len(paths)
	COROUTNUM := thread
	groupLength := length / COROUTNUM
	wg.Add(10)
	for i := 0; i < COROUTNUM; i++ {
		go getpath(paths[i*groupLength : (i+1)*groupLength])
	}
	go getpath(paths[COROUTNUM*groupLength:])
	wg.Wait()
	logger.Info("Done!")
	logger.Info("================================")

}

func getpath(paths []string) {
	for _, v := range paths {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)

			}
		}()
		//fmt.Println(v)
		if strings.Index(v, "/") != 0 {
			v = "/" + v

		}
		resp, err := http.Head(host + string(v))

		if err != nil {
			panic(err)
		}
		if resp.StatusCode == 200 {
			logger.Info(host+string(v), "-->", resp.StatusCode)
		}
		resp.Body.Close()

	}
	wg.Done()
}
