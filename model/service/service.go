package service

type Data struct {
	Data   []CryptoCurrencyMap  `json:"data"`
	Status CryptoCurrencyStatus `json:"status"`
}

type CryptoCurrencyMap struct {
	Id                  int                       `json:"id"`
	Rank                int                       `json:"rank"`
	Name                string                    `json:"name"`
	Symbol              string                    `json:"symbol"`
	Slug                string                    `json:"slug"`
	IsActive            int                       `json:"is_active"`
	Status              string                    `json:"status"`
	FirstHistoricalData string                    `json:"first_historical_data"`
	LastHistoricalData  string                    `json:"last_historical_data"`
	Platform            CryptoCurrencyMapPlatform `json:"platform"`
}

type CryptoCurrencyMapPlatform struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"token_address"`
}

type CryptoCurrencyStatus struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
}

type Quote struct {
	Price             float64 `json:"price"`
	Volume24h         float64 `json:"volume_24h"`
	VolumeChange24h   float64 `json:"volume_change_24h"`
	Volume24hReported float64 `json:"volume_24h_reported"`
	Volume7d          float64 `json:"volume_7d"`
	Volume7dReported  float64 `json:"volume_7d_reported"`
}

type QuotesLatest struct {
	Id                int                       `json:"id"`
	Name              string                    `json:"name"`
	Symbol            string                    `json:"symbol"`
	Slug              string                    `json:"slug"`
	NumMarketPairs    int                       `json:"num_market_pairs"`
	DateAdded         string                    `json:"date_added"`
	Tags              []string                  `json:"tags"`
	MaxSupply         float64                   `json:"max_supply"`
	CirculatingSupply float64                   `json:"circulating_supply"`
	TotalSupply       float64                   `json:"total_supply"`
	Platform          CryptoCurrencyMapPlatform `json:"platform"`
	CmcRank           int                       `json:"cmc_rank"`
	IsFiat            int                       `json:"is_fiat"`
	LastUpdated       string                    `json:"last_updated"`
	Quote             map[string]Quote          `json:"quote"`
}

type ListingLatest struct {
	Data   map[string]QuotesLatest `json:"data"`
	Status CryptoCurrencyStatus    `json:"status"`
}
