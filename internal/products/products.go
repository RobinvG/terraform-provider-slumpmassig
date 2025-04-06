package products

import (
	"math/rand"
	"unicode"
)

type ProductsParams struct {
	Length     int64
}

type ProductsResult struct {
	Result string
	Spongecase string
	L33t string
	Diacritics string
}

func convertToSpongecase(in string) string {
	out := make([]rune, len(in))
	var newChar rune

	for pos, char := range []rune(in) {
		toUpper := rand.Float32() > 0.5
		if toUpper {
			newChar = unicode.ToUpper(char)
		} else {
			newChar = unicode.ToLower(char)
		}
		out[pos] = newChar
	}

	return string(out)
}

// ReturnProduct generates a random product name based on the parameters
func (p *ProductsParams) ReturnProduct() (ProductsResult, error) {
	var err error
	var product_results ProductsResult

	product_results.Result = products[rand.Intn(len(products))]
	product_results.Diacritics = removeDiacritics(product_results.Result)
	product_results.Spongecase = convertToSpongecase(product_results.Result)
	product_results.L33t = convertToLeet(product_results.Diacritics)

	return product_results, err
}
