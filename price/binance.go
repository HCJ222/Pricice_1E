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

func (ps *PriceService) Binance(log *zap.Logger, paymentCurrency string, currency string) {


        for {
                func() {

                        defer func() {
                                if r := recover(); r != nil {
                                        // log
                                }

                              time.Sleep(cfg.Config.Options.Interval * time.Second)
//                                time.Sleep(10 * time.Second)
                        }()


                        changeCurrency := strings.ToUpper(currency) +strings.ToUpper(paymentCurrency)

//                        resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" +changeCurrency)
                        resp, err := http.Get(cfg.Config.APIs.Binance +changeCurrency)
			// log
                        if err != nil {
                                // handle error
                                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to fetch from Binance_" +fmt.Sprint(err)))
                        }

                        defer func() {
                                resp.Body.Close()
                        }()


                        body, err := ioutil.ReadAll(resp.Body)
			// log
		        if err != nil {
		                // handle error
		                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to read body from Binance_" +fmt.Sprint(err)))
		        }

                        bp := binancePrice{}
                        json.Unmarshal(body, &bp)

                        price := bp.Price
			log.Info("Price", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Binance_" +strings.ToUpper(currency) +" to " +strings.ToUpper(paymentCurrency), price))

                        ps.SetPrice(currency +"/" +paymentCurrency +"/binance", utils.StringToFloat64(price))
                }()
        }
}
