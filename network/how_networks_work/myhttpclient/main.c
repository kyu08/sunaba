#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#define BUF_SIZE 1000

// example.comにGET requestを送る簡易的なhttp client
// (学習目的なのでrequest先はexample.com固定にしている)
int main() {
  // ドメイン名からIPアドレスを取得する
  struct addrinfo hints, *result;
  memset(&hints, 0, sizeof(hints));
  hints.ai_family = AF_UNSPEC;
  hints.ai_socktype = SOCK_STREAM;

  // getaddrinfoを呼び出してリクエスト先のIPアドレスを特定する
  // https://man7.org/linux/man-pages/man3/getaddrinfo.3.html
  int status = getaddrinfo("example.com", "80", &hints, &result);
  if (status != 0) {
    fprintf(stderr, "getaddrinfo failed. status: %d\n", status);
    return -1;
  }

  // sfdにsocketのファイルディスクリプタを格納する
  // NOTE: resultはlinked_listとして返却される。
  // rpは現在の処理対象を格納するための一時変数。
  int sfd;
  struct addrinfo *rp;
  for (rp = result; rp != NULL; rp = rp->ai_next) {
    sfd = socket(rp->ai_family, rp->ai_socktype, rp->ai_protocol);
    if (sfd == -1) {
      continue;
    }

    // サーバーへの接続を確立する
    if (connect(sfd, rp->ai_addr, rp->ai_addrlen) != -1) {
      // 接続成功のケース
      break;
    }

    // 接続に失敗したケース
    close(sfd);
  }

  freeaddrinfo(result);

  if (rp == NULL) {
    fprintf(stderr, "failed to connect\n");
    return -1;
  }

  // HTTPリクエストを送信する
  char *http_request_message = "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n";
  size_t req_size = strlen(http_request_message);

  // https://man7.org/linux/man-pages/man2/write.2.html
  if (write(sfd, http_request_message, req_size) != req_size) {
    fprintf(stderr, "failed to write\n");
    return -1;
  }

  // responseを読み込む
  char buf[BUF_SIZE];
  ssize_t nread = read(sfd, buf, BUF_SIZE - 1);
  if (nread == -1) {
    fprintf(stderr, "failed to read\n");
    return -1;
  }
  buf[nread] = '\0';
  printf("received %zd bytes:\n%s\n", nread, buf);

  // ソケットをcloseする
  close(sfd);

  return 0;
}
