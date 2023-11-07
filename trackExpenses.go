package main

import "fmt"

type Expense struct {
	Amount             float64
	Description        string
	ExpensesSharedWith []int
}

func (sw *Splitwise) AddExpense(userID int, description string, amount float64, sharedWith []int) {
	defer sw.mutex.Unlock()
	expense := Expense{
		Description:        description,
		Amount:             amount,
		ExpensesSharedWith: sharedWith,
	}
	sw.mutex.Lock()
	sw.Users[userID].Expenses = append(sw.Users[userID].Expenses, expense)
}

func (sw *Splitwise) GetBalance(userID int) {
	user := sw.Users[userID]
	balance := make(map[int]float64)

	for _, expense := range user.Expenses {
		for _, friendID := range expense.ExpensesSharedWith {
			if friendID == userID {
				continue
			}
			balance[friendID] += expense.Amount / float64(len(expense.ExpensesSharedWith)-1)
		}
	}

	for friendID, amount := range balance {
		if amount > 0 {
			fmt.Printf("%s owes you ₹%.2f\n", sw.Users[friendID].Name, amount)
		} else if amount < 0 {
			fmt.Printf("You owe %s ₹%.2f\n", sw.Users[friendID].Name, -amount)
		}
	}
}

func (sw *Splitwise) ViewExpenseHistory(userID int) {
	user := sw.Users[userID]

	fmt.Printf("Expense history for %s:\n", user.Name)
	for i, expense := range user.Expenses {
		sharedWithNames := make([]string, 0)
		for _, friendID := range expense.ExpensesSharedWith {
			sharedWithNames = append(sharedWithNames, sw.Users[friendID].Name)
		}

		fmt.Printf("%d. Description: %s, Amount: ₹%.2f, Shared with: %v\n", i+1, expense.Description, expense.Amount, sharedWithNames)
	}
}
