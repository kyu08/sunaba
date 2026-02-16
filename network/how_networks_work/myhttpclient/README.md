# 簡易的なhttp clientの実装

```sh
$ make run
clang main.c -o main && ./main
received 842 bytes:
HTTP/1.1 200 OK
Date: Mon, 16 Feb 2026 14:31:02 GMT
Content-Type: text/html
Transfer-Encoding: chunked
Connection: keep-alive
CF-RAY: 9cedbf11495dddfe-NRT
Last-Modified: Thu, 12 Feb 2026 14:00:39 GMT
Allow: GET, HEAD
Accept-Ranges: bytes
Age: 8651
cf-cache-status: HIT
Server: cloudflare

210
<!doctype html><html lang="en"><head><title>Example Domain</title><meta name="viewport" content="width=device-width, initial-scale=1"><style>body{background:#eee;width:60vw;mar
gin:15vh auto;font-family:system-ui,sans-serif}h1{font-size:1.5em}div{opacity:0.8}a:link,a:visited{color:#348}</style></head><body><div><h1>Example Domain</h1><p>This domain is
 for use in documentation examples without needing permission. Avoid use in operations.</p><p><a href="https://iana.org/domains/example">Learn more</a></p></div></body></html>

0
```
