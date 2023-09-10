package price

type coinoneTradeHistory struct {
	Trades []struct {
		Timestamp     uint64 `json:"timestamp"`
		Price         string `json:"price"`
		Volume        string `json:"volume"`
		IsSellerMaker bool   `json:"is_seller_maker"`
	}
}

type bithumbTransactionHistory struct {
	Data []struct {
		TransactionData string `json:"transaction_data"`
		Type            string `json:"type"`
		UnitsTraded     string `json:"units_traded"`
		Price           string `json:"price"`
		Total           string `json:"total"`
	}
}

type upbitMarket struct {
	Market             string  `json:"market"`
	TradeData          string  `json:"trade_data"`
	TradeTime          string  `json:"trade_time"`
	TradeDateKst       string  `json:"trade_date_kst"`
	TradeTimeKst       string  `json:"trade_time_kst"`
	TradeTimestamp     string  `json:"trade_timestamp"`
	OpeningPrice       float64 `json:"opening_price"`
	HighPrice          float64 `json:"high_price"`
	LowPrice           float64 `json:"low_price"`
	TradePrice         float64 `json:"trade_price"`
	PrevClosingPrice   float64 `json:"prev_closing_price"`
	Change             float64 `json:"change"`
	ChangePrice        float64 `json:"change_price"`
	ChangeRate         float64 `json:"change_rate"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	TradeVolume        float64 `json:"trade_volume"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp          int64   `json:"timestamp"`
}

type binancePrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
