package main

import "fmt"

type User struct {
	UserId   int
	Name     string
	Expenses []Expense
}

type Group struct {
	Name    string
	Members []int
}

func (sw *Splitwise) AddUser(userID int, name string) {
	user := &User{UserId: userID, Name: name}
	sw.Users[userID] = user
}

func (sw *Splitwise) CreateGroup(groupName string, memberIDs []int) {
	group := &Group{
		Name:    groupName,
		Members: memberIDs,
	}
	sw.Groups[groupName] = group
}

func (sw *Splitwise) SimplifyGroupDebt(groupName string) {
	group, ok := sw.Groups[groupName]
	if !ok {
		fmt.Printf("Group %s does not exist\n", groupName)
		return
	}

	groupBalance := make(map[int]float64)

	for _, memberID := range group.Members {
		sw.mutex.Lock()
		user := sw.Users[memberID]
		sw.mutex.Unlock()
		for _, expense := range user.Expenses {
			for _, friendID := range expense.ExpensesSharedWith {
				if friendID != memberID && contains(group.Members, friendID) {
					groupBalance[memberID] += expense.Amount / float64(len(expense.ExpensesSharedWith)-1)
					groupBalance[friendID] -= expense.Amount / float64(len(expense.ExpensesSharedWith)-1)
				}
			}
		}
	}

	for _, memberID := range group.Members {
		for _, friendID := range group.Members {
			if memberID != friendID {
				amount := groupBalance[memberID] - groupBalance[friendID]
				if amount > 0 {
					fmt.Printf("%s owes %s ₹%.2f in group %s\n", sw.Users[friendID].Name, sw.Users[memberID].Name, amount, groupName)
				} else if amount < 0 {
					fmt.Printf("%s owes %s ₹%.2f in group %s\n", sw.Users[memberID].Name, sw.Users[friendID].Name, -amount, groupName)
				}
			}
		}
	}
}

func contains(arr []int, item int) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}
