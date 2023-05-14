package main

import (
	"booking-app/helper"
	"booking-app/model"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Package level variables
// Define outside all functions
// They can be accessed inside any of the functions and in all files, which are in the same package
// Best Practice: Define variable as "local" as possible, they can be accessed only inside that function or block of code. Create the variable where you need it

const conferenceTickets int = 50

// sugar syntax instead of-> var variableName = "value". Syntax only for var, not const
var conferenceName = "Go Conference"

// uint only positives int
var remainingTickets uint = 50

// var bookings []string // slice of string
// var bookings = make([]map[string]string, 0) // make an empty list of maps. make function for a List of Maps needs an initial value

// var bookings = make([]model.User, 0) or bellow
var bookings []model.User

var wg = sync.WaitGroup{}

// go run main.go // to run the program
// go run booking-app
// go run main.go helper.go
// or go run -gcflags "all=-N -l" main
// or go run .
// run the main function of the 'go run' command
func main() {
	// empty array
	// var bookings = [50]string{}
	// another way to declare an empty array type
	// var bookings [50]string

	// To use slices we need create an array without size specification
	// Another syntax for an empty array
	// var bookings = []string{}
	// Another syntax for an empty array
	//bookings := []string{}
	// Slices is an abstraction of an Array, like a list, and it can expand automatically
	//var bookings []string
	//fmt.Printf("-- Bookings type: %T\n", bookings)

	// fmt is a go builtin package
	// Printf and Println are functions
	//var test = greetUsers
	//fmt.Printf("Test type: %T", test)
	//test()

	greetUsers()

	city := whichCity()

	// Array has fixed size and cannot grow up, syntax: var variableName = [numberOfElements]type{theElements}
	//var bookings = [50]string{"John", "Maria", "Peter"} with some elements at declare time
	//fmt.Printf("Values for bookings: %v, bookings type: %T\n", bookings, bookings)

	// define a variable, (keyword)var variableName string(type)
	// go has only the 'for' loop, and the for loop has same types
	// infinite loop for true {} => Syntax: 'for' while the condition is true, doing something inside here {}
	// if the loop for is an infinite loop, we can omit the keyword 'true' and use this: for {}
	for {

		firstName, lastName, email, userTickets := getUserInputs()

		//remainingTickets = remainingTickets - userTickets
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			var newUser = bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1) // add the following one new threads to the main group
			// the go keyword makes the application concurrency, just creating a new thread for sendTicket
			go sendTicket(newUser)

			if remainingTickets == 0 { // bool type
				fmt.Println("We've done for today. We already had sold all the tickets. Have a nice day!")

				// the go builtin 'range' keywords is used for iterates through all the array elements and others types, like slices, strings and maps
				// other type of 'for' to iterate over and over the array elements
				//for index, value := range bookings {
				//	fmt.Printf("-- Index is [%v], and the value is: %v\n", index, value)
				//}
				firstNames := getFirstNames()
				fmt.Printf("\n>>> The first names of bookings are: %v <<<\n", firstNames)

				// at the end the program has finished
				break // go out of the infinite loop and the finish the program
			}
		} else { // else if condition can be here if necessary
			// instead of this for we can use continue to start the loop again
			/*for {
				fmt.Printf("%v %v, we have no more tickets to sell. We only have '%v' tickets. Please, try again!\n", firstName, lastName, remainingTickets)
				// ask user for their tickets
				fmt.Println("Enter number of tickets: ")
				_, err = fmt.Scan(&userTickets)
				if err != nil {
					panic(err)
				}
				if remainingTickets >= userTickets {
					break
				}
			}*/
			//fmt.Printf("%v %v, we have no more tickets to sell. We only have '%v' tickets. Please, try again!\n\n", firstName, lastName, remainingTickets)
			validationMessage := printInputValidationError(isValidName, isValidEmail, isValidTicketNumber)
			// in golang all string or has a value or is empty, nil is not accept for type string
			if validationMessage != "" {
				fmt.Println(validationMessage)
			}
			//continue
		}
	}
	printSelectedCity(city)

	fmt.Println("Waiting until all processing be done!...")
	wg.Wait() // wait until all threads are removed from the wait group
	fmt.Println("All processing was done!")
}

// Function is only executed, when "called"! with ()
// This is greeting the users with a great message
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// This is the function documentation.
// whichCity is called for user to choose the conference city.
// You can choose by the number or the city's name and the value will be returned.
func whichCity() string {
	var city string
	initialMessage := "Which city will be the conference?" +
		"\nType the number or city's name to choose one!" +
		"\n1. New York" +
		"\n2. Singapore" +
		"\n3. London" +
		"\n4. Berlin" +
		"\n5. Mexico City" +
		"\n6. Hong Kong"
	fmt.Println(initialMessage)
	_, _ = fmt.Scan(&city) // _ _ just scape to use them, the return value or the error
	return city
}

// getFirstNames just format the first names of the bookings slice and return them as a slice of first names
func getFirstNames() []string {
	var firstNames []string
	// keyword range to iterate through all map elements
	for _, element := range bookings {
		//var names = strings.Fields(element) // split by white spaces into a slice
		//firstNames = append(firstNames, names[0])
		//var firstName = element["firstName"] // return "" empty string when the key not found because this is a string behavior
		//firstNames = append(firstNames, firstName)
		firstNames = append(firstNames, element.FirstName)
	}
	return firstNames
}

// getUserInputs get the first name, last name, email address, ticket's number and return them in the same mentioned order
func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	_, err := fmt.Scan(&firstName)
	// syntax := is to create the variables on the left hand side
	// blank identifier '_' is discard by compiler and ignore a variable you don't want to use. But it needs there because the signature
	// &firstName the keyword & is to access the memory pointer of the variable to store the value
	if err != nil {
		panic(err)
	}

	// Example of the variable value and the address memory of pointer
	//fmt.Printf("-- Variable value: %v, pointer value: %v, variable type: %T\n", firstName, &firstName, firstName)

	// ask user for their last name
	fmt.Println("Enter your last name: ")
	_, err = fmt.Scan(&lastName)
	// syntax = is only to assigned values to the already exists variables
	if err != nil {
		panic(err)
	}

	// ask user for their email
	fmt.Println("Enter you email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		panic(err)
	}

	// ask user for their tickets
	fmt.Println("Enter number of tickets: ")
	_, err = fmt.Scan(&userTickets)
	if err != nil {
		panic(err)
	}

	return firstName, lastName, email, userTickets
}

// prints the confirmation user booking and remaining tickets
func bookTicket(userTickets uint, firstName string, lastName string, email string) model.User {
	remainingTickets -= userTickets

	// make creates an empty map of key string and value string
	// map syntax: map[keyType]valueType
	// map in goland is { key: value }
	//var userData = make(map[string]string)
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // converts uint to string
	// append the map to the list of maps
	//bookings = append(bookings, userData)

	var newUser = model.User{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}
	fmt.Printf("The full name of user is '%v'\n", newUser.GetFullName())
	bookings = append(bookings, newUser)
	fmt.Printf("List of bookings is %v\n", bookings)

	// append is a go builtin function that appends new elements to an array and return a new array with the elements
	//bookings = append(bookings, firstName+" "+lastName)

	/*
		fmt.Printf("-- The whole array: %v\n", bookings)
		fmt.Printf("-- The first value: %v\n", bookings[0])
		fmt.Printf("-- Array type: %T\n", bookings)
		fmt.Printf("-- Array length: %v\n", len(bookings))
	*/
	fmt.Printf("Thank you '%v %v' for booking '%v' tickets. You will receive a confirmation email at '%v'\n", firstName, lastName, userTickets, email)
	fmt.Printf("'%v' tickets remaining for '%v'\n", remainingTickets, conferenceName)

	return newUser
}

// Returns the messages validation error for the user inputs if any occur
func printInputValidationError(isValidName bool, isValidEmail bool, isValidTicketNumber bool) string {
	var validationMessages []string
	if !isValidName {
		validationMessages = append(validationMessages, "'first name or last name is too short'")
	}
	if !isValidEmail {
		validationMessages = append(validationMessages, "'email address doesn't contains '@' sign'")
	}
	if !isValidTicketNumber {
		validationMessages = append(validationMessages, "'number of tickets is invalid'")
	}
	return fmt.Sprintf("Your input data is invalid.\nWhat's wrong: %v.\nPlease, try again!", strings.Join(validationMessages, ", "))
}

// prints the selected city by the user for the conference
func printSelectedCity(city string) {
	// switch statement
	message := "Great! You've selected the '%v' as the conference city. Good job!\n"
	switch city {
	case "1", "New York":
		fmt.Printf(message, "New York")
	case "2", "Singapore":
		fmt.Printf(message, "Singapore")
	case "3", "London":
		fmt.Printf(message, "London")
	case "4", "Berlin":
		fmt.Printf(message, "Berlin")
	case "5", "Mexico City":
		fmt.Printf(message, "Mexico City")
	case "6", "Hong Kong":
		fmt.Printf(message, "Hong Kong")
	default:
		fmt.Println("Sorry, but you didn't choose a valid city for the conference. We will discard all the bookings!")
	}
}

func sendTicket(user model.User) {
	// simulating the creation and sending ticket for email with delay
	time.Sleep(50 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", user.NumberOfTickets, user.FirstName, user.LastName)
	fmt.Println("################################################")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, user.Email)
	fmt.Println("################################################")

	wg.Done() // remove from the wait group this thread because it's done
}
