package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alice": {
		AuthToken: "123ABC",
		Username:  "alice",
	},
	"bob": {
		AuthToken: "456DEF",
		Username:  "bob",
	},
	"charle": {
		AuthToken: "789GHI",
		Username:  "charle",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alice": {
		Coins:    100,
		Username: "alice",
	},
	"bob": {
		Coins:    200,
		Username: "bob",
	},
	"charle": {
		Coins:    75,
		Username: "charle",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var coinData = CoinDetails{}
	coinData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &coinData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
