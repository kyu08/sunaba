package main

// Train はColleagueインターフェース
// 各列車がMediatorを通じて間接的に通信するための共通インターフェースを定義
type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}
