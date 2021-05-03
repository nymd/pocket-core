package types

var ValidatorCacheSize int64 = 5
var SinfoCacheSize int64 = 8000

func InitConfig(validatorCacheSize int64) {
	ValidatorCacheSize = validatorCacheSize
}
