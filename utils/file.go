package utils

import (
	"bufio"
	"github.com/axgle/mahonia"
	"os"
	"io/ioutil"
)


func ReadLines(path string)(lines []string,err error){
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	file ,err := os.Open(path)
	defer file.Close()
	if err != nil{
		return nil,err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines,enc.ConvertString(scanner.Text()))
	}
	return lines,err
}

func ReadFile(path string) ( s string){
	c,err := ioutil.ReadFile(path)
	if err == nil{
		return string(c)
	}else {
		return ""
	}

}
