type Bank struct {
    accs []int64
}


func Constructor(balance []int64) Bank {
    return Bank{
        accs: balance,
    }
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if len(this.accs) < account1 || len(this.accs) < account2 {
        return false
    }
    if this.accs[account1 - 1] >= money {
        this.accs[account1 - 1] -= money
        this.accs[account2 - 1] += money
        return true
    }
    return false
}


func (this *Bank) Deposit(account int, money int64) bool {
    if len(this.accs) >= account {
        this.accs[account - 1] += money
        return true
    }
    return false
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if len(this.accs) >= account && this.accs[account - 1] >= money {
        this.accs[account - 1] -= money
        return true
    }
    return false
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */