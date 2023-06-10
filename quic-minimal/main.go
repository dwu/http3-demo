package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/logging"
	"github.com/quic-go/quic-go/qlog"
)

const addr = "localhost:4242"

var keyLog io.Writer
var qconf quic.Config

func main() {
	keyLogFile := flag.String("keylog", "", "key log file")
	enableQlog := flag.Bool("qlog", false, "output a qlog (in the same directory)")
	flag.Parse()

	if len(*keyLogFile) > 0 {
		f, err := os.Create(*keyLogFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		keyLog = f
	}

	if *enableQlog {
		qconf.Tracer = func(ctx context.Context, p logging.Perspective, connID quic.ConnectionID) logging.ConnectionTracer {
			filename := fmt.Sprintf("client_%x.qlog", connID)
			f, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Creating qlog file %s.\n", filename)
			return qlog.NewConnectionTracer(NewBufferedWriteCloser(bufio.NewWriter(f), f), p, connID)
		}
	}

	go func() { echoServer() }()
	fmt.Println("Waiting before sending message...")
	time.Sleep(1 * time.Second)

	err := clientMain()
	if err != nil {
		panic(err)
	}
}

func echoServer() {
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server: Waiting for request message...")
	conn, err := listener.Accept(context.Background())
	if err != nil {
		panic(err)
	}
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		request := readMessage(stream)
		fmt.Printf("Server: Got message with payload '%s'\n", request)
		response := fmt.Sprintf("response %s", request)
		writeMessage(stream, response)
	}
}

// Writes a string message with length prefix
func writeMessage(stream quic.Stream, msg string) {
	mbuflen := []byte{byte(len(msg))}
	stream.Write(mbuflen)
	stream.Write([]byte(msg))
}

// Reads a string message with length prefix
func readMessage(stream quic.Stream) string {
	mbuflen := []byte{0}
	stream.Read(mbuflen)
	mbuf := make([]byte, mbuflen[0])
	stream.Read(mbuf)
	return string(mbuf)
}

func clientMain() error {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
		KeyLogWriter:       keyLog,
	}
	conn, err := quic.DialAddr(context.Background(), addr, tlsConf, &qconf)
	if err != nil {
		return err
	}

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < 3; i += 1 {
		message := fmt.Sprintf("message %d", i)
		fmt.Printf("Client: Sending request message '%s'\n", message)
		writeMessage(stream, message)
		fmt.Println("Client: Wrote message, waiting for response...")

		response := readMessage(stream)
		fmt.Printf("Client: Got response message '%s'\n", response)
	}

	return nil
}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
