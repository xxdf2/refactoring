package old

type PriceCode int

const CHILDREN PriceCode = 2
const REGULAR PriceCode = 0
const NEW_RELEASE PriceCode = 1

type Movie struct {
	title     string
	priceCode PriceCode
}

func NewMovie(title string, priceCode PriceCode) *Movie {
	return &Movie{
		title:     title,
		priceCode: priceCode,
	}
}

func (this *Movie) GetPriceCode() PriceCode {
	return this.priceCode
}

func (this *Movie) SetPriceCode(priceCode PriceCode) {
	this.priceCode = priceCode
}

func (this *Movie) GetTitle() string {
	return this.title
}
