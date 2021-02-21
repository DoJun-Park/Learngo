package accounts

import (
	"errors"
)

type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
// go에서 Constructor는 func을 만들어서 object를 return
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// go언어에서는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 정의된다.

// 함수명 앞에 있는 a를 receiver라고 부른다. 리시버가 있는 함수는 메소드가 된다.
//리시버는 struct명의 첫글자를 따서 소문자로 지어야한다.
// Deposit x amout on your account
func (a *Account) Deposit(amount int) {
	// fmt.Println("Gonna deposit", amount)
	a.balance += amount // 여기서 a는 main에서 호출한 account(실제)의 a가 아닌 실제 account의 복사본이기 때문에 값을 변경하고 main에서 account의 balance를 출력해도 값이 안바뀜
	// 때문에 만약 실제 값까지 바꾸고 싶다면 pointer receiver 바꿔줘야함
}

// balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor ")
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// go가 자동적으로 호출해주는 메서드 String
func (a Account) String() string {
	return "whatever you want"
}
