package metric

import (
	"go.uber.org/zap"

	price "github.com/HCJ222/Pricice_1E/price"
)

var (
	metricData metric

	GaugesNamespaceList = [...]string{"usd_krw_dunamu",

					"btc_krw_upbit",
					"btc_usdt_upbit",
					"btc_usdt_binance",
					"btc_usdt_houbiGlobal",

					"atom_krw_coinone",
					"atom_krw_upbit",
					"atom_usdt_binance",
					"atom_usdt_houbiGlobal",
					"atom_btc_binance",
					"atom_btc_houbiGlobal",

					"luna_krw_coinone",
					"luna_krw_bithumb",
					"luna_btc_upbit",

					"iris_usdt_houbiGlobal",
					"iris_btc_houbiGlobal",

					"kava_usdt_binance",
					"kava_btc_binance",
					"kava_krw_coinone",

					"band_usdt_binance",

					"sol_busd_binance",
                                        "sol_btc_binance",

				}
)

type metric struct {

	USD struct {
		KRW struct {
			Dunamu	float64
		}
	}

	BTC struct {
		KRW struct {
			Upbit	float64
		}
		USDT struct {
			Upbit       float64
                        Binance     float64
                        HuobiGlobal float64
                }
	}

	ATOM struct {
		KRW struct {
			Coinone	float64
			Upbit	float64
		}
		BTC struct {
			Upbit		float64
			Binance		float64
			HuobiGlobal	float64
		}
		USDT struct {
			Binance         float64
                        HuobiGlobal float64
                }
	}

	LUNA struct {
                KRW struct {
                        Coinone	float64
                        Bithumb	float64
		}
		BTC struct {
			Upbit   float64
		}
        }

	IRIS struct {
		BTC struct {
                        HuobiGlobal float64
                }
                USDT struct {
                        HuobiGlobal float64
		}
	}

	KAVA struct {
                BTC struct {
                        Binance	float64
                }
		USDT struct {
			Binance	float64
		}
		KRW struct {
                        Coinone float64
                }
        }

	BAND struct {
                USDT struct {
                        Binance float64
                }
        }

	SOL struct {
                BTC struct {
                        Binance float64
                }
                BUSD struct {
                        Binance float64
                }
        }
}



func SetMetric(log *zap.Logger, ps *price.PriceService) {

	// USD to KRW
	metricData.USD.KRW.Dunamu = ps.GetPrice("usd/krw/dunamu")

	// BTC
	metricData.BTC.KRW.Upbit = ps.GetPrice("btc/krw/upbit")
	metricData.BTC.USDT.Upbit = ps.GetPrice("btc/usdt/upbit")
	metricData.BTC.USDT.Binance = ps.GetPrice("btc/usdt/binance")
	metricData.BTC.USDT.HuobiGlobal = ps.GetPrice("btc/usdt/huobiGlobal")

	// ATOM
	metricData.ATOM.KRW.Coinone = ps.GetPrice("atom/krw/coinone")
	metricData.ATOM.KRW.Upbit = ps.GetPrice("atom/krw/upbit")
	metricData.ATOM.BTC.Upbit = ps.GetPrice("atom/btc/upbit")
	metricData.ATOM.BTC.Binance = ps.GetPrice("atom/btc/binance")
	metricData.ATOM.BTC.HuobiGlobal = ps.GetPrice("atom/btc/huobiGlobal")
	metricData.ATOM.USDT.HuobiGlobal = ps.GetPrice("atom/usdt/huobiGlobal")
	metricData.ATOM.USDT.Binance = ps.GetPrice("atom/usdt/binance")

	// IRIS
	metricData.IRIS.BTC.HuobiGlobal = ps.GetPrice("iris/btc/huobiGlobal")
	metricData.IRIS.USDT.HuobiGlobal = ps.GetPrice("iris/usdt/huobiGlobal")

	// Kava
	metricData.KAVA.BTC.Binance = ps.GetPrice("kava/btc/binance")
	metricData.KAVA.USDT.Binance = ps.GetPrice("kava/usdt/binance")
	metricData.KAVA.KRW.Coinone = ps.GetPrice("kava/krw/coinone")

}

func GetMetric() *metric {

	return &metricData
}
