package services

import (
	"encoding/json"
	"github.com/cryptowat-go/server/config"
	"github.com/cryptowat-go/server/models"
	"github.com/cryptowat-go/server/utils"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type ethService struct {
	cfg config.Config
	db  *gorm.DB
}

func (e ethService) GetCurrentPrice() float64 {
	resp, err := http.Get("https://api.cryptowat.ch/markets/kraken/ethusd/price")
	if err != nil {
		log.Fatal("Error getting price data from CryptoWat.ch API")
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Error parsing price data from CryptoWat.ch API")
	}

	if price, ok := data["result"].(map[string]interface{})["price"].(float64); ok {
		return price
	}

	log.Fatal("Error getting price from CryptoWat.ch API response")
	return 0
}

func (e ethService) UpSertETHPosition(position models.Position) error {
	err := e.db.Where("id=?", position.ID).FirstOrCreate(&position).Error
	if err != nil {
		return err
	}
	return nil
}

func (e ethService) GetPositionByTime() {
	//TODO implement me
	panic("implement me")
}

func (e ethService) GetAllPositions() ([]models.Position, error) {
	var positions []models.Position
	err = e.db.Find(&positions).Error
	if err != nil {
		return positions, err
	}
	return positions, err
}

func NewEthService(cfg config.Config, db *gorm.DB) ETHServices {
	return &ethService{cfg, db}
}

func (eth *ethService) InitWebSocket() {
	// Connect to CryptoWat.ch WebSocket API
	cryptoSocket, _, err = websocket.DefaultDialer.Dial("wss://stream.cryptowat.ch/connect?apikey="+eth.cfg.CryptoWatchApiKey, nil)
	if err != nil {
		log.Fatal("Error connecting to CryptoWat.ch WebSocket API")
	}
	// Listen for WebSocket events
	// Read first message, which should be an authentication response
	_, message, err := cryptoSocket.ReadMessage()
	var authResult struct {
		AuthenticationResult struct {
			Status string `json:"status"`
		} `json:"authenticationResult"`
	}
	err = json.Unmarshal(message, &authResult)
	if err != nil {
		panic(err)
	}

	// Send a JSON payload to subscribe to a list of resources
	// Read more about resources here: https://docs.cryptowat.ch/websocket-api/data-subscriptions#resources
	resources := []string{
		"markets:96:trades",
	}
	subMessage := struct {
		Subscribe SubscribeRequest `json:"subscribe"`
	}{}
	// No map function in golang :-(
	for _, resource := range resources {
		subMessage.Subscribe.Subscriptions = append(subMessage.Subscribe.Subscriptions, Subscription{StreamSubscription: StreamSubscription{Resource: resource}})
	}
	msg, err := json.Marshal(subMessage)
	err = cryptoSocket.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		panic(err)
	}
	go eth.loopToUpdateData()
}

func (eth *ethService) loopToUpdateData() {
	for {
		_, message, err := cryptoSocket.ReadMessage()
		if err != nil {
			log.Fatalf("Error reading from CryptoWat.ch WebSocket API %v\n", err)
		}

		var data map[string]interface{}
		err = json.Unmarshal(message, &data)
		if err != nil {
			log.Fatal("Error parsing WebSocket message from CryptoWat.ch")
		}

		if value, ok := data["marketUpdate"].(map[string]interface{}); ok {
			if value1, ok := value["tradesUpdate"].(map[string]interface{}); ok {
				if value2, ok := value1["trades"].([]interface{}); ok {
					for _, v := range value2 {

						mu.Lock()

						if len(positions) == 0 {
							var position models.Position
							position.Value = utils.ParseFloat(v.(map[string]interface{})["priceStr"].(string)) * utils.ParseFloat(v.(map[string]interface{})["amountUSDStr"].(string))
							position.UpdatedAt = time.Now()
							position.Currency = "eth"
							err = eth.UpSertETHPosition(position)
							if err != nil {
								log.Println("Error updating position in database:", err)
							}
						}

						mu.Unlock()
					}
				}
			}
		}
	}

}
