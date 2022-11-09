package main

import "fmt"

func main() {

	user1 := &User{
		name: "Régis Perez",
		age:  35,
	}
	user2 := &User{
		name: "Test",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}