package bank

var deposits = make(chan int) // 入金額を送信する
var balances = make(chan int) // 残高を受信する
var withdrawal = make(chan int)
var withdrawRes = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdrawal <- amount
	return <-withdrawRes
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdrawal:
			if amount > balance {
				withdrawRes <- false
			} else {
				balance -= amount
				withdrawRes <- true
			}
		}
	}
}

func init() {
	go teller()
}
