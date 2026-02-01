package main

// Mediator は仲介者インターフェース
// Colleagueオブジェクト間の通信を調整するためのインターフェースを定義
type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}
