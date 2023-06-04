## References

- qlog visualization, <https://qvis.quictools.info/#/files>
- HTTP/3 explained, <https://http3-explained.haxx.se/>
- Curl with HTTP3 enabled, <https://github.com/yurymuski/curl-http3>

## Licenses

- http3-client is a slightly modified version of Cloudflare's [http3-client example](https://github.com/cloudflare/quiche/blob/master/quiche/examples/http3-client.c) with support for the `SSLKEYLOGFILE` environment variable.

## Misc notes

### Create self-signed certificate

```sh
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
  -keyout localhost.key -out localhost.crt -subj "/CN=localhost" \
  -addext "subjectAltName=DNS:localhost,IP:127.0.0.1,172.17.0.1"
```