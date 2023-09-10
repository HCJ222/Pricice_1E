package exporter

import (
	"time"

	"go.uber.org/zap"

	metric "github.com/HCJ222/Pricice_1E/exporter/metric"
	price "github.com/HCJ222/Pricice_1E/price"

	"github.com/prometheus/client_golang/prometheus"
)

var ()

func Start(log *zap.Logger) {

	gaugesNamespaceList := metric.GaugesNamespaceList

	var gauges []prometheus.Gauge = make([]prometheus.Gauge, len(gaugesNamespaceList))

	// nomal guages
	for i := 0; i < len(gaugesNamespaceList); i++ {
		gauges[i] = metric.NewGauge("price", gaugesNamespaceList[i], "")
		prometheus.MustRegister(gauges[i])
	}

	ps := price.NewPriceService()
	ps.OnStart(log)

	time.Sleep(1 * time.Second)

	for {
		func() {

			defer func() {

				if r := recover(); r != nil {
					log.Error("Recovered from panic", zap.Any("panic", r))
				}

				time.Sleep(1000 * time.Millisecond)
			}()

			metric.SetMetric(log, ps)
			metricData := metric.GetMetric()

			gaugesValue := [...]float64{

				metricData.USD.KRW.Dunamu,

				metricData.BTC.KRW.Upbit,
				metricData.BTC.USDT.Upbit,
				metricData.BTC.USDT.Binance,

				metricData.ATOM.KRW.Coinone,
				metricData.ATOM.KRW.Upbit,
				metricData.ATOM.USDT.Binance,
				metricData.ATOM.BTC.Binance,

				metricData.KAVA.KRW.Upbit,
				metricData.KAVA.USDT.Binance,
				metricData.KAVA.BTC.Binance,

				metricData.CTK.KRW.Bithumb,
				metricData.CTK.USDT.Binance,
				metricData.CTK.BTC.Binance,

				metricData.IRIS.USDT.Binance,
				metricData.IRIS.BTC.Binance,

				metricData.ORC.KRW.Bithumb,

				metricData.SOL.KRW.Bithumb,
				metricData.SOL.KRW.Coinone,
				metricData.SOL.USDT.Binance,
				metricData.SOL.BTC.Binance,

				metricData.CERE.USDT.HuobiGlobal,

				metricData.ETH.KRW.Bithumb,
			}

			for i := 0; i < len(gaugesNamespaceList); i++ {
				gauges[i].Set(gaugesValue[i])
			}
		}()
	}
}
