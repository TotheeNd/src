package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var ip = "10.251.26.114"   // 此处需要改造为支持命令行输入
	timeout := 1 * time.Second // 超时时间
	for i := 21; i <= 120; i++ {
		var address = fmt.Sprintf("%s:%d", ip, i)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			fmt.Println(address, "是关闭的")
			continue
		}
		conn.Close()
		fmt.Println(address, "打开")
	}
}
