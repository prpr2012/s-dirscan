package main

import (
	"fmt"
	"github.com/schollz/progressbar/v2"
	"time"
)

func main() {
	bar := progressbar.New(100)
	go out(bar)
	for{
		bar.Clear()
		fmt.Println("1")
		time.Sleep(time.Second)
	}



	//for i := 0; i < 50; i++ {
	//	time.Sleep(100 * time.Millisecond)
	//	h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
	//	fmt.Printf("\r%.0f%%[%s]", float64(i)/49*100, h)
	//	os.Stdout.Sync()
	//}
	//fmt.Println("\nAll System Go!")
	//for i:=0;i<10;i++{
	//	time.Sleep(time.Second)
	//
	//	io.WriteString(os.Stdout,"\r"+strconv.Itoa(i))
	//	os.Stdout.Sync()
	//	fmt.Println("\rtest")
	//
	//}
}
func out(bar *progressbar.ProgressBar)  {
	for i := 0; i < 100000000; i++ {
		bar.Add(1)
		time.Sleep(300*time.Millisecond)
	}

}
