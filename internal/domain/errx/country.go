package errx

import (
	"github.com/chains-lab/ape"
)

var ErrorCountryNotFound = ape.DeclareError("COUNTRY_NOT_FOUND")

var ErrorCountryAlreadyExistsWithThisID = ape.DeclareError("COUNTRY_ALREADY_EXISTS_WITH_THIS_ID")

var ErrorInvalidCountryISO3ID = ape.DeclareError("INVALID_COUNTRY_ISO3_ID")
