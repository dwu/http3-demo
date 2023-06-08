# HTTP/3 Demo

- [HTTP/3 Demo](#http3-demo)
  - [Visualization HTTP3 GET](#visualization-http3-get)
    - [Wireshark](#wireshark)
    - [qvis](#qvis)
  - [HTTP/3 Streams](#http3-streams)
  - [HTTP/3 frame layout](#http3-frame-layout)
  - [References](#references)
  - [Licenses](#licenses)
  - [Misc notes](#misc-notes)
    - [Create self-signed certificate](#create-self-signed-certificate)

## Visualization HTTP3 GET

### Wireshark

![](docs/http3-client-GET-wireshark.png)

### qvis

![](docs/http3-client-GET-qvis.png)

## HTTP/3 Streams

- Control Stream
- Request Stream
- Push Stream

## HTTP/3 frame layout

```
HTTP/3 Frame Format {
  Type (i),
  Length (i),
  Frame Payload (..),
}
```

```
DATA Frame {
  Type (i) = 0x00,
  Length (i),
  Data (..),
}
```

```
HEADERS Frame {
  Type (i) = 0x01,
  Length (i),
  Encoded Field Section (..),
}
```

## References

- HTTP/3 RFC, https://www.rfc-editor.org/rfc/rfc9114.html
- HTTP/3 explained, <https://http3-explained.haxx.se/>
- qlog visualization, <https://qvis.quictools.info/#/files>
- Curl with HTTP3 enabled, <https://github.com/yurymuski/curl-http3>
- Quiche HTTP/3 client library, <https://github.com/cloudflare/quiche>
- Generate Random Extensions And Sustain Extensibility (GREASE), <https://textslashplain.com/2020/05/18/a-bit-of-grease-keeps-the-web-moving/>
  - Application in TLS, <https://www.rfc-editor.org/rfc/rfc8701.html>

## Licenses

- http3-client is a slightly modified version of Cloudflare's [http3-client example](https://github.com/cloudflare/quiche/blob/master/quiche/examples/http3-client.c) with support for the `SSLKEYLOGFILE` environment variable.

## Misc notes

### Create self-signed certificate

```sh
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
  -keyout localhost.key -out localhost.crt -subj "/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,IP:127.0.0.1,172.17.0.1"
```
