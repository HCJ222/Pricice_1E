package price

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

	cfg "github.com/HCJ222/Pricice_1E/config"
)

func (ps *PriceService) Upbit(log *zap.Logger, paymentCurrency string, currency string) {

	for {
		func() {

			defer func() {
				if r := recover(); r != nil {
					// log
				}

				time.Sleep(cfg.Config.Options.Interval * time.Second)
			}()

			changeCurrency := strings.ToUpper(paymentCurrency) + "-" + strings.ToUpper(currency)

			resp, err := http.Get(cfg.Config.APIs.Upbit + changeCurrency)
			// log
			if err != nil {
				// handle error
				log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to fetch from Upbit_"+fmt.Sprint(err)))
			}

			defer func() {
				resp.Body.Close()
			}()

			body, err := io.ReadAll(resp.Body)
			// log
			if err != nil {
				// handle error
				log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to read body from Upbit_"+fmt.Sprint(err)))
			}

			th := []upbitMarket{}
			json.Unmarshal(body, &th)

			price := th[0].TradePrice
			priceString := strconv.FormatFloat(price, 'f', -1, 64)

			log.Info("Price", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Bithumb_"+strings.ToUpper(currency)+" to "+strings.ToUpper(paymentCurrency), priceString))

			ps.SetPrice(currency+"/"+paymentCurrency+"/upbit", price)
		}()
	}
}
