package main

import (
	"sync"
)

type Splitwise struct {
	Users  map[int]*User
	Groups map[string]*Group
	mutex  sync.Mutex
}

func NewSplitwise() *Splitwise {
	return &Splitwise{
		Users:  make(map[int]*User),
		Groups: make(map[string]*Group),
	}
}

func main() {
	splitwise := NewSplitwise()

	splitwise.AddUser(1, "Sai")
	splitwise.AddUser(2, "Kumar")
	splitwise.AddUser(3, "Neelam")

	splitwise.AddExpense(1, "Dinner", 30.0, []int{1, 2})
	splitwise.AddExpense(2, "Waterworld", 200.0, []int{1, 2, 3})

	splitwise.GetBalance(1)
	splitwise.GetBalance(2)

	splitwise.ViewExpenseHistory(1)
	splitwise.ViewExpenseHistory(2)

	splitwise.CreateGroup("Colleagues", []int{1, 2, 3})
	splitwise.SimplifyGroupDebt("Colleagues")
}
