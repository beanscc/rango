package internal

// CurrencyCode currency code
// 数据来源：https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml
// https://www.iso.org/iso-4217-currency-codes.html
// https://www.six-group.com/en/products-services/financial-information/data-standards.html
// #Current Currency# List One
type CurrencyCode struct {
	// the english short name of currency
	Name string
	// three letters code.
	// The first two letters of the ISO 4217 three-letter code are the same as the code for the country name, and,
	// where possible, the third letter corresponds to the first letter of the currency name.
	Code string

	// The three-digit numeric code is useful when currency codes need to be understood in countries that do not use
	// Latin scripts and for computerized systems. Where possible, the three-digit numeric code is the same as the
	// numeric country code.
	Numeric string

	// For currencies having minor units, ISO 4217:2015 also shows the relationship between the minor unit and the
	// currency itself (i.e. whether it divides into 100 or 1000).
	MinorUnits int
}
