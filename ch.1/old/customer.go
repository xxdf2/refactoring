package old

import (
	"fmt"
	"strconv"
	"strings"
)

type Customer struct {
	name    string
	rentals []*Rental
}

func NewCustomer(name string) *Customer {
	return &Customer{
		name: name,
	}
}

func (this *Customer) AddRental(rental *Rental) {
	this.rentals = append(this.rentals, rental)
}

func (this *Customer) GetName() string {
	return this.name
}

func (this *Customer) Statement() string {
	var totalAmount float64
	var frequentRenterPoints int
	var sb strings.Builder
	sb.WriteString("Rental Record for " + this.GetName() + "\n")
	for _, rental := range this.rentals {
		var thisAmount float64
		switch rental.GetMovie().GetPriceCode() {
		case REGULAR:
			thisAmount += 2
			if rental.GetDayRented() > 2 {
				thisAmount += float64(rental.GetDayRented()-2) * 1.5
			}
		case NEW_RELEASE:
			thisAmount += float64(rental.dayRented) * 3
		case CHILDREN:
			thisAmount += 1.5
			if rental.dayRented > 3 {
				thisAmount += float64(rental.dayRented-3) * 1.5
			}
		}

		// 常客加积分
		frequentRenterPoints++

		// 新影片额外奖励
		if rental.GetMovie().GetPriceCode() == NEW_RELEASE && rental.GetDayRented() > 1 {
			frequentRenterPoints++
		}
		sb.WriteString("\t")
		sb.WriteString(rental.GetMovie().GetTitle())
		sb.WriteString("\t")
		sb.WriteString(strconv.FormatFloat(thisAmount, 'f', 2, 64))
		totalAmount += thisAmount
	}

	sb.WriteString("\n Amount owed is ")
	sb.WriteString(fmt.Sprintf("%.2f", totalAmount))
	sb.WriteString("\n")
	sb.WriteString("You earned ")
	sb.WriteString(strconv.Itoa(frequentRenterPoints))
	sb.WriteString(" frequent renter points")
	return sb.String()
}
