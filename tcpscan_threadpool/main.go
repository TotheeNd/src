package main

//线程池方式
import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/loveleshsharma/gohive"
)

//wg
var wg sync.WaitGroup

//地址管道,100容量
var addressChan = make(chan string, 100)

//工人
func worker() {
	//函数结束释放连接
	defer wg.Done()
	for {
		address, ok := <-addressChan
		if !ok {
			break
		}
		//fmt.Println("address:", address)
		conn, err := net.DialTimeout("tcp", address, time.Second*3)
		if err != nil {
			//fmt.Println("close:", address, err)
			continue
		}
		conn.Close()
		fmt.Println("open:", address)
	}
}
func main() {
	var begin = time.Now()
	//ip
	var ip = "10.251.21.56"
	//线程池大小
	var pool_size = 30000                         //此处后续修改为可读取命令行参数版本
	var pool = gohive.NewFixedSizePool(pool_size) //此处还需修改

	//拼接ip:端口
	//启动一个线程,用于生成ip:port,并且存放到地址管道种
	go func() {
		for port := 1; port <= 65535; port++ {
			var address = fmt.Sprintf("%s:%d", ip, port)
			//将address添加到地址管道
			//fmt.Println("<-:",address)
			addressChan <- address
		}
		//发送完关闭 addressChan 管道
		close(addressChan)
	}()
	//启动pool_size工人,处理addressChan种的每个地址
	for work := 0; work < pool_size; work++ {
		wg.Add(1)
		pool.Submit(worker)
	}
	//等待结束
	wg.Wait()
	//计算时间
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时:", elapseTime)
}
