package cryptocurrency

// CryptoCurrency is a cryptocurrency
type CryptoCurrency string

func (c CryptoCurrency) String() string {
	return string(c)
}

const (
	// ETH is Ethereum
	ETH CryptoCurrency = "ETH"
	// BTC is Bitcoin
	BTC CryptoCurrency = "BTC"
	// XMR is Monero
	XMR CryptoCurrency = "XMR"
)
