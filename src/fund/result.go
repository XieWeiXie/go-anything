package fund

type (
	OneFoundResult struct {
		Code              string  `json:"code"`
		Name              string  `json:"name"`
		Type              string  `json:"type"`
		BuyRate           string  `json:"buy_rate"`
		Scale             string  `json:"scale"`
		Date              string  `json:"date"`
		Company           string  `json:"company"`
		Price             float64 `json:"price"`
		GuessPrice        float64 `json:"guess_price"`
		Rate              string  `json:"rate"`
		GuessRate         string  `json:"guess_rate"`
		Stocks            []Stock `json:"stocks"`
		YearForEarn       string  `json:"year_for_earn"`
		SixMonthForEarn   string  `json:"six_month_for_earn"`
		ThreeMonthForEarn string  `json:"three_month_for_earn"`
		MonthForEarn      string  `json:"month_for_earn"`
	}
	Stock struct {
		Code  string `json:"code"`
		Name  string `json:"name"`
		Rate  string `json:"rate"`  // 占比
		Ratio string `json:"ratio"` // 本身的涨幅
		Price string `json:"price"` // 当前价位
	}
	Manager struct {
		Id       string       `json:"id"`
		Name     string       `json:"name"`
		Star     int          `json:"star"`
		WorkTime string       `json:"work_time"`
		Size     string       `json:"size"`
		Power    ManagerPower `json:"power"`
	}
	ManagerPower struct {
		Description []OnePowerSeries `json:"description"`
	}
	OnePowerSeries struct {
		Category string  `json:"category"`
		Value    float64 `json:"value"`
	}
	ManagerProfit struct {
		Description []OneProfitSeries `json:"description"`
	}
	OneProfitSeries struct {
		Value float64 `json:"value"`
	}
)
