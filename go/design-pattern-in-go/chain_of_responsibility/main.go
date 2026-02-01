package main

func main() {
	// 連鎖を構築（終端から逆順で設定）
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	// 患者を作成
	patient := &Patient{name: "Tanaka"}

	// 連鎖の先頭（受付）にリクエストを送信
	reception.execute(patient)
}
