package gors

import "sync"

type ResponseHandlerFunc func(e *Error, resp any) any

var responseHandlerFunc ResponseHandlerFunc
var responseHandlerFuncMu sync.RWMutex

func RegisterResponseHandler(f ResponseHandlerFunc) {
	responseHandlerFuncMu.Lock()
	responseHandlerFunc = f
	responseHandlerFuncMu.Unlock()
}

func GetResponseHandler() ResponseHandlerFunc {
	var f ResponseHandlerFunc
	responseHandlerFuncMu.RLock()
	f = responseHandlerFunc
	responseHandlerFuncMu.RUnlock()
	return f
}
