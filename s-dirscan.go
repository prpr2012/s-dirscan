package main

import (
	"flag"
	"fmt"
	"github.com/wonderivan/logger"
	"github.com/schollz/progressbar/v3"
	"net/http"
	//"os"
	"s-dirscan/utils"
	//"strconv"
	"strings"
	"sync"
	//"time"
)

var wg sync.WaitGroup
var (
	host   string
	dic    string
	thread int
	sfile  string
)
var scanedNum int
var pathlen *int
var mutex sync.RWMutex
func main() {
	fmt.Print(utils.ReadFile("banner.txt"))
	flag.StringVar(&host, "h", "", "target")
	flag.StringVar(&dic, "d", "dic/dic.txt", "dic")
	flag.IntVar(&thread, "t", 10, "goroutine num")
	flag.StringVar(&sfile,"f","","read hosts from specific file")
	flag.Parse()

	paths,err2 := utils.ReadLines(dic)
	n := len(paths)
	pathlen = &n
	if err2 !=nil{
		fmt.Println("打开文件出错:",err2)
		return
	}
	bar := progressbar.Default(int64(*pathlen))
	go printProcess(bar)
	if sfile != ""{
		logger.SetLogger("config/log.json")
		hosts,err1 := utils.ReadLines(sfile)
		if err1 == nil{
			for _,host := range hosts{
				execTask(bar,host,paths)
			}
		}else {
			fmt.Println("打开文件出错:",err1)
		}

	}else if host !=""{
		execTask(bar,host,paths)
	} else {
		flag.PrintDefaults()
	}

}

func printProcess(bar *progressbar.ProgressBar)  {
	for {
		mutex.RLock()
		bar.Set(scanedNum)
		mutex.RUnlock()
	}

}

func execTask(bar *progressbar.ProgressBar,host string,paths []string){
	scanedNum = 0
	fmt.Printf("target:   %v  |  scanning... \n",host)
	if host!=""&&!utils.Check(host) {
		fmt.Println("plase input right url or target is dead...")
		return
	}
	length := len(paths)
	COROUTNUM := thread
	groupLength := length / COROUTNUM
	wg.Add(COROUTNUM)
	for i := 0; i < COROUTNUM; i++ {
		go getpath(bar,host,paths[i*groupLength : (i+1)*groupLength])
	}
	go getpath(bar,host,paths[COROUTNUM*groupLength:])
	wg.Wait()
	bar.Clear()
	logger.Info("Done!")

}

func getpath(bar *progressbar.ProgressBar,host string,paths []string) {
	for _, v := range paths {
		defer func() {
			if err := recover(); err != nil {
				bar.Clear()
				logger.Error(err)
			}
		}()

		if strings.Index(v, "/") != 0 {
			v = "/" + v
		}

		resp, err := http.Head(host + string(v))

		if err != nil {
			panic(err)
		}
		if resp.StatusCode == 200 {
			bar.Clear()
			fmt.Println(host+string(v), "-->", resp.StatusCode)
		}
		mutex.Lock()
		scanedNum += 1
		mutex.Unlock() // 写锁

		resp.Body.Close()
	}
	wg.Done()
}
