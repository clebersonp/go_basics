package helper

import (
	"strings"
)

// Scope: Package leve
// Variables and Functions defined outside any function, can be accessed in all other files within the same package

// ValidateUserInput validate the user inputs, such as: first name, last name, email address, user number tickets and number of remaining tickets.
//
// All the inputs need to be valid.
//
// - For the first name and last name inputs, they must be greater or igual 2.
//
// - Email address must have the @ sign.
//
// - The number of user tickets must have greater than 0 Zero and the number of remaining tickets must be tested as greater or equal to the user tickets.
//
// returns all results for each validation. Cool golang feature!
//
// name validation, email validation and ticket number validation respectively
//
// To export a function, just capitalize first letter of function name
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // and condition &&, or condition ||
	return isValidName, isValidEmail, isValidTicketNumber
}
