package exporter

import (
	"time"
	"go.uber.org/zap"

	price "github.com/HCJ222/Pricice_1E/price"
	metric "github.com/HCJ222/Pricice_1E/exporter/metric"

	"github.com/prometheus/client_golang/prometheus"
)

var (

)

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
					//Error Log
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
					metricData.BTC.USDT.HuobiGlobal,

					metricData.ATOM.KRW.Coinone,
					metricData.ATOM.KRW.Upbit,
					metricData.ATOM.USDT.Binance,
					metricData.ATOM.USDT.HuobiGlobal,
					metricData.ATOM.BTC.Binance,
					metricData.ATOM.BTC.HuobiGlobal,

					metricData.IRIS.USDT.HuobiGlobal,
					metricData.IRIS.BTC.HuobiGlobal,

					metricData.KAVA.USDT.Binance,
					metricData.KAVA.BTC.Binance,
					metricData.KAVA.KRW.Coinone,

				}

				for i := 0; i < len(gaugesNamespaceList); i++ {
					gauges[i].Set(gaugesValue[i])
				}
		}()
	}
}
