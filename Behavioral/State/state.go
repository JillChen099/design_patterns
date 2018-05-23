/*
Created on 2018/5/18 9:20

author: ChenJinLong

Content:
在状态模式中，我们将对象在每一个状态下的行为和状态转移语句封装在一个个状态类中，
通过这些状态类来分散冗长的条件转移语句，让系统具有更好的灵活性和可扩展性。
*/
package State

import "fmt"

// 账户银行，环境类
type Account struct {
	state     AccountState
	owner     string
	balance   float64
}

func NewAccount(owner string,init float64) *Account {
	a := new(Account)
	a.owner = owner
	a.balance = 0
	a.state = &NormalState{a}
	fmt.Printf("%s开户,初始金额为%v\n",owner,init)
	fmt.Println("---------------------------------------------")
	return a
}

func (a *Account) getBalance() float64 {
	return a.balance
}

func (a *Account) setBalance(balance float64) {
	a.balance = balance
}

func (a *Account) setState(state AccountState) {
	a.state = state
}
//存款
func (a *Account) Deposit(amount float64) {
	fmt.Printf("%s存款%v\n",a.owner,amount)
	a.state.deposit(amount)
	fmt.Printf("现在余额为%v\n",a.balance)
	fmt.Printf("现在账户状态为%v\n",a.state.getStateName())
	fmt.Println("---------------------------------------------")
}

//取款
func (a *Account) Withdraw(amount float64) {
	fmt.Printf("%s取款款%v\n",a.owner,amount)
	a.state.withdraw(amount)
	fmt.Printf("现在余额为%v\n",a.balance)
	fmt.Printf("现在账户状态为%v\n",a.state.getStateName())
	fmt.Println("---------------------------------------------")
}

func (a *Account) ComputeInterest() {
	a.state.computeInterest()
}

// 抽象状态接口
type AccountState interface {
	deposit(amount float64)
	withdraw(amount float64)
	computeInterest()
	getStateName() string
	stateCheck()
}

//正常状态类
type NormalState struct {
	acc *Account
}

func (n *NormalState) getStateName() string{
	return "正常状态"
}
//存钱
func (n *NormalState) deposit(amount float64) {
	n.acc.setBalance(n.acc.getBalance()+amount)
	n.stateCheck()
}
//取钱
func (n *NormalState) withdraw(amount float64) {
	n.acc.setBalance(n.acc.getBalance()-amount)
	n.stateCheck()
}

func (n *NormalState) computeInterest() {
	fmt.Println("正常状态，无需支付利息！")
}
func( n *NormalState) stateCheck() {
	balance := n.acc.getBalance()
	if balance >0 {
		fmt.Println("维持正常状态")
	}else if balance > -2000 {
		n.acc.setState(&OverdraftState{n.acc})
	}else {
		n.acc.setState(&RestrictedState{n.acc})
	}
}



type OverdraftState struct {
	acc *Account
}

func (o *OverdraftState) getStateName() string{
	return "透支状态"
}
//存钱
func (o *OverdraftState) deposit(amount float64) {
	o.acc.setBalance(o.acc.getBalance()+amount)
	o.stateCheck()
}

//取钱
func (o *OverdraftState) withdraw(amount float64) {
	o.acc.setBalance(o.acc.getBalance()-amount)
	o.stateCheck()
}

func (o *OverdraftState) computeInterest() {
	fmt.Println("计算利息！")
}
func( o *OverdraftState) stateCheck() {
	balance := o.acc.getBalance()
	if balance >0 {
		o.acc.setState(&NormalState{o.acc})
	}else if balance > -2000 {
		fmt.Println("维持透支状态")
	}else {
		o.acc.setState(&RestrictedState{o.acc})
	}
}




type RestrictedState struct {
	acc *Account
}

func (r *RestrictedState) getStateName() string{
	return "受限状态"
}
//存钱
func (r *RestrictedState) deposit(amount float64) {
	r.acc.setBalance(r.acc.getBalance()+amount)
	r.stateCheck()
}

//取钱
func (r *RestrictedState) withdraw(amount float64) {
	fmt.Println("账户受限，取款失败")
}

func (r *RestrictedState) computeInterest() {
	fmt.Println("计算利息！")
}

func( r *RestrictedState) stateCheck() {
	balance := r.acc.getBalance()
	if balance >0 {
		r.acc.setState(&NormalState{r.acc})
	}else if balance > -2000 {
		r.acc.setState(&OverdraftState{r.acc})
	}else {
		fmt.Println("维持受限状态")
	}
}








