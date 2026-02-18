package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

const (
	host    = "example.com"
	port    = 80
	bufSize = 1000
)

// example.comにGET requestを送る簡易的なhttp client
// (学習目的なのでrequest先はexample.com固定にしている)
func main() {
	// 1. DNS解決（net.LookupIP）
	//    ドメイン名からIPアドレスを取得する
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to lookup IP: %v\n", err)
		os.Exit(1)
	}

	// IPv4アドレスを探す
	var ipv4 net.IP
	for _, ip := range ips {
		if ip4 := ip.To4(); ip4 != nil {
			ipv4 = ip4
			break
		}
	}
	if ipv4 == nil {
		fmt.Fprintf(os.Stderr, "no IPv4 address found for %s\n", host)
		os.Exit(1)
	}

	// 2. ソケット作成（syscall.Socket）
	//    AF_INET, SOCK_STREAM, 0 でTCPソケット作成
	sfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create socket: %v\n", err)
		os.Exit(1)
	}
	defer syscall.Close(sfd)

	// 3. サーバーへ接続（syscall.Connect）
	//    SockaddrInet4 構造体を使用し、ポート80に接続
	addr := &syscall.SockaddrInet4{
		Port: port,
		Addr: [4]byte{ipv4[0], ipv4[1], ipv4[2], ipv4[3]},
	}
	if err := syscall.Connect(sfd, addr); err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect: %v\n", err)
		os.Exit(1)
	}

	// 4. HTTPリクエスト送信（syscall.Write）
	httpRequest := "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"
	n, err := syscall.Write(sfd, []byte(httpRequest))
	if err != nil || n != len(httpRequest) {
		fmt.Fprintf(os.Stderr, "failed to write: %v\n", err)
		os.Exit(1)
	}

	// 5. レスポンス受信（syscall.Read）
	buf := make([]byte, bufSize)
	nread, err := syscall.Read(sfd, buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("received %d bytes:\n%s\n", nread, string(buf[:nread]))

	// 6. ソケットクローズ（syscall.Close）
	//    defer で解放済み
}
