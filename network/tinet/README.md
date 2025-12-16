体験しながら学ぶ ネットワーク技術入門の学習用repo


# ref
セットアップスクリプトはこちらからダウンロードした。
https://www.sbcr.jp/support/4815617794/

# Chapter 1 検証環境を作ろう

```sh
# UBUNTU環境の作成
multipass launch 20.04 --cpus 2 --name UBUNTU --mount /Users/kyu08/code/tinet/:/mnt/c/tinet
```

# Chapter2 ネットワーク2プロトコルを知ろう
`tcpdump -i net0 -w /tmp/tinet/ethernet.pcap ether host xxx`を実行したところ`tcpdump: Couldn't change ownership of savefile`と表示されたが、以下の記事に従ってパスを変更したところ解決した。
https://qiita.com/riv_infra/items/945478c49714dee44329
