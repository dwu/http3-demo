package main

import (
	"context"
	"crypto/tls"
	"encoding/gob"
	"flag"
	"fmt"
	"os"

	"github.com/quic-go/quic-go"

	"example.com/quic-bench/internal/messages"
)

const addr = "localhost:4242"

func main() {
	keyFile := flag.String("keyfile", "", "server key file")
	crtFile := flag.String("crtfile", "", "server cert file")
	flag.Parse()

	if len(*keyFile) == 0 || len(*crtFile) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	cert, err := tls.LoadX509KeyPair(*crtFile, *keyFile)
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"quic-bench"},
	}

	listener, err := quic.ListenAddr(addr, tlsConfig, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server: Waiting for connection...")
	conn, err := listener.Accept(context.Background())
	if err != nil {
		panic(err)
	}

	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(stream)

	var filenameMessage messages.MessageFilename
	dec.Decode(&filenameMessage)
	fmt.Printf("Got filename: %s\n", filenameMessage.Filename)

	var messageSizeMessage messages.MessageSize
	dec.Decode(&messageSizeMessage)
	fmt.Printf("Got message size: %d\n", messageSizeMessage.Size)

	f, err := os.Create(filenameMessage.Filename)
	defer f.Close()

	var dataMessage messages.MessageData
	for {
		dec.Decode(&dataMessage)
		//fmt.Printf("Got data: size=%d, payload='%s'\n", dataMessage.DataLen, dataMessage.Data)
		_, err := f.Write(dataMessage.Data)
		if err != nil {
			panic(err)
		}
		if dataMessage.DataLen < messageSizeMessage.Size {
			fmt.Println("Done.")
			os.Exit(0)
		}
	}
}
