package main

import (
	"fmt"
	"sync"
	"time"

	models "github.com/mainflux/mqtt-test/models"
	mqttclient "github.com/mainflux/mqtt-test/mqttclient"
)

func main() {
	var wg = sync.WaitGroup{}
	clients := make(map[int]*mqttclient.MqttClient)
	startTime := time.Now()
	var config models.Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Can't load config yaml")
	}
	fmt.Println(cfg)
	// Map certain number of clients regarding to clients count parameter
	for i := 0; i < cfg.ClientsCount; i++ {
		c := &mqttclient.MqttClient{
			ID:       i,
			Username: cfg.MfxAccessUsername,
			Password: cfg.MfxAccessToken,
			Payload:  cfg.SenMLPayload,
		}
		clients[i] = c
		wg.Add(1)
		go mqttclient.MakeMqttClient(cfg, &wg, c)
	}

	wg.Wait()
	fmt.Println("Testing Done! In", time.Since(startTime))

}
