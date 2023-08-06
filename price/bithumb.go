package price

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
	"strings"

	"go.uber.org/zap"

	utils "github.com/HCJ222/Pricice_1E/utils"
	cfg "github.com/HCJ222/Pricice_1E/config"
)

func (ps *PriceService) Bithumb(log *zap.Logger, currency string) {


	for {
		func() {

			defer func() {
				if r := recover(); r != nil {
					// log
				}

				time.Sleep(cfg.Config.Options.Interval * time.Second)
			}()


			changeCurrency := strings.ToUpper(currency) +"_KRW"

//			resp, err := http.Get("https://api.bithumb.com/public/transaction_history/" +changeCurrency)
			resp, err := http.Get(cfg.Config.APIs.Bithumb +changeCurrency)
			// log
			if err != nil {
                                // handle error
                                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to fetch from Bithumb_" +fmt.Sprint(err)))
                        }

			defer func() {
				resp.Body.Close()
			}()

			body, err := ioutil.ReadAll(resp.Body)
			// log
                        if err != nil {
                                // handle error
                                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to read body from Bithumb_" +fmt.Sprint(err)))
                        }

			th := bithumbTransactionHistory{}
			json.Unmarshal(body, &th)


			price := th.Data[len(th.Data)-1].Price

			log.Info("Price", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Bithumb_" +strings.ToUpper(currency) +" to KRW", price))

			ps.SetPrice(currency +"/krw/bithumb", utils.StringToFloat64(price))
		}()
	}
}

