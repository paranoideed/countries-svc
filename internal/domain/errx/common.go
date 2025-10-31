package errx

import (
	"github.com/chains-lab/ape"
)

var ErrorInternal = ape.DeclareError("INTERNAL_ERROR")

var ErrorInvalidCityStatus = ape.DeclareError("INVALID_CITY_STATUS")

var ErrorInvalidSlug = ape.DeclareError("INVALID_SLUG")

var ErrorInvalidTimeZone = ape.DeclareError("INVALID_TIME_ZONE")
