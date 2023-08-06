package price

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"go.uber.org/zap"

	utils "github.com/HCJ222/Pricice_1E/utils"
	cfg "github.com/HCJ222/Pricice_1E/config"
)

func (ps *PriceService) Dunamu(log *zap.Logger) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					//
				}

				time.Sleep(cfg.Config.Options.Interval * time.Second)
			}()

			resp, err := http.Get(cfg.Config.APIs.Dunamu)
			// log
                        if err != nil {
                                // handle error
                                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to fetch from Dunamu_" +fmt.Sprint(err)))
                        }

			defer func() {
				resp.Body.Close()
			}()
			body, err := ioutil.ReadAll(resp.Body)
			// log
                        if err != nil {
                                // handle error
                                log.Fatal("Price", zap.Bool("Success", false), zap.String("err", "Fail to read body from Dunamu_" +fmt.Sprint(err)))
                        }

			re, _ := regexp.Compile("\"basePrice\":[0-9.]+")
			str := re.FindString(string(body))
			re, _ = regexp.Compile("[0-9.]+")

			price := re.FindString(str)
			log.Info("Price", zap.Bool("Success", true), zap.String("err", "nil"), zap.String("Dunamu_USD to KRW", price))

			ps.SetPrice("usd/krw/dunamu", utils.StringToFloat64(price))
		}()
	}
}
