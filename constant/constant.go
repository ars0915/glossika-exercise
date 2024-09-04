package constant

import "time"

const (
	ServiceName        = "glossika-exercise"
	ResponseCodePrefix = 1

	CacheRecommendProductTTL = 10 * time.Minute

	CacheKeyPrefixRecommend = "recommend-"
)
