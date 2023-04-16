package old

import "testing"

func TestNewCustomer(t *testing.T) {
	c := NewCustomer("eric")
	c.AddRental(NewRental(NewMovie("harry poter", NEW_RELEASE), 2))
	c.AddRental(NewRental(NewMovie("pig henry", CHILDREN), 3))
	c.AddRental(NewRental(NewMovie("day up", REGULAR), 30))
	s := c.Statement()
	println(s)
}
