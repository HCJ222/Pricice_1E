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

		"atom_krw_coinone",
		"atom_krw_upbit",
		"atom_usdt_binance",
		"atom_btc_binance",

		"kava_krw_upbit",
		"kava_usdt_binance",
		"kava_btc_binance",

		"ctk_krw_bithumb",
		"ctk_usdt_binance",
		"ctk_btc_binance",

		"iris_usdt_binance",
		"iris_btc_binance",

		"orc_krw_bithumb",

		"sol_krw_bithumb",
		"sol_krw_coinone",
		"sol_usdt_binance",
		"sol_btc_binance",
	}
)

type metric struct {
	USD struct {
		KRW struct {
			Dunamu float64
		}
	}

	BTC struct {
		KRW struct {
			Upbit float64
		}
		USDT struct {
			Upbit   float64
			Binance float64
		}
	}

	ATOM struct {
		KRW struct {
			Coinone float64
			Upbit   float64
		}
		BTC struct {
			Upbit   float64
			Binance float64
		}
		USDT struct {
			Binance float64
		}
	}

	KAVA struct {
		KRW struct {
			Upbit float64
		}
		USDT struct {
			Binance float64
		}
		BTC struct {
			Binance float64
		}
	}

	CTK struct {
		KRW struct {
			Bithumb float64
		}
		USDT struct {
			Binance float64
		}
		BTC struct {
			Binance float64
		}
	}

	IRIS struct {
		USDT struct {
			Binance float64
		}
		BTC struct {
			Binance float64
		}
	}

	ORC struct {
		KRW struct {
			Bithumb float64
		}
	}

	SOL struct {
		KRW struct {
			Bithumb float64
			Coinone float64
		}
		USDT struct {
			Binance float64
		}
		BTC struct {
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

	// ATOM
	metricData.ATOM.KRW.Coinone = ps.GetPrice("atom/krw/coinone")
	metricData.ATOM.KRW.Upbit = ps.GetPrice("atom/krw/upbit")
	metricData.ATOM.BTC.Upbit = ps.GetPrice("atom/btc/upbit")
	metricData.ATOM.BTC.Binance = ps.GetPrice("atom/btc/binance")
	metricData.ATOM.USDT.Binance = ps.GetPrice("atom/usdt/binance")

	// Kava
	metricData.KAVA.KRW.Upbit = ps.GetPrice("kava/krw/upbit")
	metricData.KAVA.USDT.Binance = ps.GetPrice("kava/usdt/binance")
	metricData.KAVA.BTC.Binance = ps.GetPrice("kava/btc/binance")

	// CTK
	metricData.CTK.KRW.Bithumb = ps.GetPrice("ctk/krw/bithumb")
	metricData.CTK.USDT.Binance = ps.GetPrice("ctk/usdt/binance")
	metricData.CTK.BTC.Binance = ps.GetPrice("ctk/btc/binance")

	// IRIS
	metricData.IRIS.USDT.Binance = ps.GetPrice("iris/usdt/binance")
	metricData.IRIS.BTC.Binance = ps.GetPrice("iris/btc/binance")

	// ORC
	metricData.ORC.KRW.Bithumb = ps.GetPrice("orc/krw/bithumb")

	// SOL
	metricData.SOL.KRW.Bithumb = ps.GetPrice("sol/krw/bithumb")
	metricData.SOL.KRW.Coinone = ps.GetPrice("sol/krw/coinone")
	metricData.SOL.USDT.Binance = ps.GetPrice("sol/usdt/binance")
	metricData.SOL.BTC.Binance = ps.GetPrice("sol/btc/binance")

}

func GetMetric() *metric {

	return &metricData
}
