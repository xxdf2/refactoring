package old

type Rental struct {
	movie     *Movie
	dayRented int
}

func NewRental(movie *Movie, dayRented int) *Rental {
	return &Rental{
		movie:     movie,
		dayRented: dayRented,
	}
}

func (this *Rental) GetDayRented() int {
	return this.dayRented
}

func (this *Rental) GetMovie() *Movie {
	return this.movie
}
