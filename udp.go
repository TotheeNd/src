package main

import (
    "fmt"
    "net"
    "flag"
)
var (
	address    = flag.String("web.listen-address", ":18383", "the address to listen http request")
	// statusCode = http.StatusOK
)

func init() {
	flag.Parse()
}

func main() {
    // 监听地址和端口号
    // address := "0.0.0.0:8080"

    // 创建UDP连接
    conn, err := net.ListenPacket("udp", *address)
    if err != nil {
        fmt.Printf("Failed to listen on UDP: %s\n", err)
        return
    }
    defer conn.Close()

    fmt.Printf("UDP server is listening on %s\n", *address)

    // 接收数据包
    buffer := make([]byte, 1024)
    for {
        n, addr, err := conn.ReadFrom(buffer)
        if err != nil {
            fmt.Printf("Failed to read UDP packet: %s\n", err)
            continue
        }

        // 打印接收到的数据
        fmt.Printf("Received UDP packet from %s: %s\n", addr.String(), string(buffer[:n]))

        // 发送响应数据
        _, err = conn.WriteTo(buffer[:n], addr)
        if err != nil {
            fmt.Printf("Failed to send UDP response: %s\n", err)
        }
    }

//       // 构建UDP数据包
//       payload := []byte("Hello, UDP!") // 数据内容
//       destAddr := "localhost:8080"
  
//       // 创建UDP连接
//       conn, err := net.Dial("udp", destAddr)
//       if err != nil {
//           fmt.Printf("Failed to create UDP connection: %s\n", err)
//           return
//       }
//       defer conn.Close()
  
//       // 发送UDP数据包
//       _, err = conn.Write(payload)
//       if err != nil {
//           fmt.Printf("Failed to send UDP packet: %s\n", err)
//           return
//       }
  
//       fmt.Println("UDP packet sent successfully!")
}