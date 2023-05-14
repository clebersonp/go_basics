package model

// custom type like a class in java
// type called User that has the following struct(Structure)
// golang has no hierarchy

type User struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

// GetFullName is like a method
func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}
