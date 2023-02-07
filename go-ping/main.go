package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int64
	size    int
	count   int
	typ     uint8 = 8
	code    uint8 = 0
)

type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	ID          uint16
	SequenceNum uint16
}

func GetCmdArgs() {
	flag.Int64Var(&timeout, "w", 10000, "timeout")
	flag.IntVar(&size, "l", 32, "length of send data")
	flag.IntVar(&count, "n", 4, "times of requests")
	flag.Parse()
}

func GenCheckSum(data []byte) uint16 {
	length := len(data)
	index := 0
	var sum uint32
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		length -= 2
		index += 2
	}
	if length == 1 {
		sum += uint32(data[index])
	}

	hi := sum >> 16
	for hi != 0 {
		sum += hi + uint32(uint16(sum))
		hi = sum >> 16
	}

	return uint16(^sum)
}

func main() {
	GetCmdArgs()
	// fmt.Println(timeout, size, count)

	dstIP := os.Args[len(os.Args)-1]

	conn, err := net.DialTimeout("ip:icmp", dstIP, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	fmt.Printf("Pinging %s [%s] with %d bytes of data:\n", dstIP, remoteAddr, size)

	for i := 0; i < count; i++ {
		icmp := &ICMP{
			Type:        typ,
			Code:        code,
			CheckSum:    uint16(0),
			ID:          uint16(i),
			SequenceNum: uint16(i),
		}

		var buffer bytes.Buffer
		_ = binary.Write(&buffer, binary.BigEndian, icmp)
		data := make([]byte, size)
		_, _ = buffer.Write(data)
		data = buffer.Bytes()
		// fmt.Println(data)
		checkSum := GenCheckSum(data)
		data[2] = byte(checkSum >> 8)
		data[3] = byte(checkSum)

		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))

		startTime := time.Now()
		n, err := conn.Write(data)
		if err != nil {
			log.Println(err)
			break
		}

		buf := make([]byte, 1024)
		n, err = conn.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("Reply from %s: bytes=%d time=%dms TTL=%d\n", remoteAddr, n-28, time.Since(startTime).Milliseconds(), buf[8])
		time.Sleep(time.Second)
	}
}

// go run main.go -n 10 www.google.com
