package utils

import (
	"net/http"
)

func Check(url string)(bool){
	_,err := http.Get(url)
	if err != nil{
		return false
	}
	return true
}
