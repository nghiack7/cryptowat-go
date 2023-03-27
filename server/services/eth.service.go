package services

import (
	"github.com/cryptowat-go/server/models"
	"github.com/gorilla/websocket"
	"sync"
)

type ETHServices interface {
	GetCurrentPrice() float64
	UpSertETHPosition(position models.Position) error
	GetPositionByTime()
	GetAllPositions() ([]models.Position, error)
	InitWebSocket()
}

type Subscription struct {
	StreamSubscription `json:"streamSubscription"`
}

type StreamSubscription struct {
	Resource string `json:"resource"`
}

type SubscribeRequest struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Update struct {
	MarketUpdate struct {
		Market struct {
			MarketId int `json:"marketId,string"`
		} `json:"market"`

		TradesUpdate struct {
			Trades []Trade `json:"trades"`
		} `json:"tradesUpdate"`
	} `json:"marketUpdate"`
}

type Trade struct {
	Timestamp     int `json:"timestamp,string"`
	TimestampNano int `json:"timestampNano,string"`

	Price  string `json:"priceStr"`
	Amount string `json:"amountStr"`
}
type Price struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

var (
	upgrader     = websocket.Upgrader{}
	cryptoSocket *websocket.Conn
	mu           sync.Mutex
	positions    []models.Position
	err          error
)
