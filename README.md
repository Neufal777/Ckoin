## C koin

Get price of different cryptocurrencies [Bitstamp]


## Usage

```python
package main

import (
	app "github.com/ckoin/app"
)

func main() {

	crypt := []string{
		"btcusd",
		"btceur",
	}
	app.CheckPrices(crypt)
}

```

## Output

```bash
BTC - USD 

{
  "ask": "6869.40",
  "bid": "6863.23",
  "count": 1,
  "high": "7323.70",
  "last": "6869.40",
  "low": "6815.22",
  "open": "7299.12",
  "timestamp": "1586531068",
  "volume": "9792.54811812",
  "vwap": "7051.56"
}


BTC - EUR 

{
  "ask": "6283.82",
  "bid": "6277.09",
  "count": 1,
  "high": "6715.57",
  "last": "6277.09",
  "low": "6250.00",
  "open": "6691.63",
  "timestamp": "1586531069",
  "volume": "2573.19308995",
  "vwap": "6445.69"
}

etc..
```
