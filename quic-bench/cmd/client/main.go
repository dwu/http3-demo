package main

import (
	"context"
	"crypto/tls"
	"encoding/gob"
	"flag"
	"os"
	"path/filepath"
	"time"

	"example.com/quic-bench/internal/messages"
	"github.com/quic-go/quic-go"
)

const addr = "localhost:4242"

func main() {
	keyFile := flag.String("keyfile", "", "server key file")
	crtFile := flag.String("crtfile", "", "server cert file")
	inFile := flag.String("infile", "", "input file name")
	messageSize := flag.Uint("size", 128, "message size")
	flag.Parse()

	if len(*keyFile) == 0 || len(*crtFile) == 0 || len(*inFile) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	cert, err := tls.LoadX509KeyPair(*crtFile, *keyFile)
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		NextProtos:         []string{"quic-bench"},
	}

	conn, err := quic.DialAddr(context.Background(), addr, tlsConfig, nil)
	if err != nil {
		panic(err)
	}

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		panic(err)
	}
	enc := gob.NewEncoder(stream)

	_, file := filepath.Split(filepath.FromSlash(*inFile))

	enc.Encode(messages.MessageFilename{
		Filename: file,
	})

	enc.Encode(messages.MessageSize{
		Size: uint32(*messageSize),
	})

	f, err := os.Open(*inFile)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, *messageSize)
	for {
		n, err := f.Read(buf)
		if err != nil {
			break
		}
		if n > 0 {
			//fmt.Println(string(buf))
			enc.Encode(messages.MessageData{
				DataLen: uint32(n),
				Data:    buf[:n],
			})
		}
	}
	time.Sleep(1 * time.Second)

}
