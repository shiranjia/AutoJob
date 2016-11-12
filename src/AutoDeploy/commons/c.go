package commons

import "log"

/**
模拟try catch 代码块
 */
func Try(fun func(),errHandler func(interface{}))  {
	defer func(){
		if err := recover(); err != nil{
			log.Println("panic a err:",err)
			errHandler(err)
		}
	}()
	fun()
}
