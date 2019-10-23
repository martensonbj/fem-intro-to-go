// // Uncomment the entire file

package main

// import "fmt"

// func main() {

// 	var userEmails map[int]string

// 	userEmails[1] = "user1@gmail.com"
// 	userEmails[2] = "user2@gmail.com"

// 	fmt.Println(userEmails)

// 	// ****************************

// 	var userEmails map[int]string = make(map[int]string)
// 	// userEmails := make(map[int]string)

// 	userEmails[1] = "user1@gmail.com"
// 	userEmails[2] = "user2@gmail.com"

// 	fmt.Println(userEmails)

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	fmt.Println(userEmails)

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	fmt.Println(userEmails)

// 	fmt.Println(userEmails[1])

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	fmt.Println(userEmails)

// 	userEmails[1] = "newUser1@gmail.com"

// 	fmt.Println(userEmails)

// 	fmt.Println(userEmails[3])

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	email1, ok := userEmails[1]
// 	fmt.Println("Email:", email1, "Present?", ok)

// 	// email3, ok := userEmails[3]
// 	// fmt.Println("Email", email3, "Present?", ok)

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	if email, ok := userEmails[1]; ok {
// 		fmt.Println(email)
// 	} else {
// 		fmt.Println("I don't know what you want from me")
// 	}

// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	delete(userEmails, 2)

// 	fmt.Println(userEmails)
// 	// ****************************

// 	userEmails := map[int]string{
// 		1: "user1@gmail.com",
// 		2: "user2@gmail.com",
// 	}

// 	for k, v := range userEmails {
// 		fmt.Printf("%s has an ID of %d.\n", v, k)
// 	}
// 	// ****************************
// }
