package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {

	var begin = time.Now()
	//wg
	var wg sync.WaitGroup
	//ip
	var ip = "10.251.21.56" //修改点1：需要修改为从命令行读取参数
	//循环
	for j := 21; j <= 65535; j++ {
		//添加wg
		wg.Add(1)
		go func(i int) {
			//释放wg
			defer wg.Done()
			var address = fmt.Sprintf("%s:%d", ip, i)
			conn, err := net.DialTimeout("tcp", address, time.Second*3)
			if err != nil {
				//fmt.Println(address, "是关闭的", err)
				return
			}
			conn.Close()
			fmt.Println(address, "打开状态") // 修改点2：这个结果需要输出到文件里面；
		}(j)
	}
	//等待wg
	wg.Wait()
	var elapseTime = time.Since(begin)
	fmt.Println("耗时:", elapseTime)
}
