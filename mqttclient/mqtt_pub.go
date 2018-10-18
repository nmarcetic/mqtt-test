package mqttclient

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/eclipse/paho.mqtt.golang"
	models "github.com/mainflux/mqtt-test/models"
)

var (
	errReadCert      = errors.New("Error reading TLS cert")
	errParseRootCert = errors.New("Faild to parse root certificate")
)

// MqttClient is test client for mqtt broker
type MqttClient struct {
	ID       int
	Username string
	Password string
	Payload  string
}

// MakeMqttClient creates and connect mqtt client to MFX MQTT broker
func MakeMqttClient(cfg *models.Config, wg *sync.WaitGroup, c *MqttClient) mqtt.Client {
	defer wg.Done()
	caCert, err := ioutil.ReadFile(cfg.TLSCertPath)
	if err != nil {
		fmt.Println(errReadCert)
		os.Exit(1)
	}
	rootCa := x509.NewCertPool()
	loadCa := rootCa.AppendCertsFromPEM([]byte(caCert))
	if !loadCa {
		fmt.Println(errParseRootCert)
		os.Exit(1)
	}
	tlsConfig := &tls.Config{
		RootCAs: rootCa,
	}
	opts := mqtt.NewClientOptions()
	opts.SetTLSConfig(tlsConfig)
	opts.AddBroker(cfg.BrokerURL)
	opts.SetUsername(c.Username)
	opts.SetPassword(c.Password)
	opts.SetClientID("mqtt-test-clientID" + strconv.Itoa(c.ID))
	mc := mqtt.NewClient(opts)

	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	wg.Add(1)
	go sendMsg(mc, wg, cfg, c.ID, c.Payload)

	return mc
}

func sendMsg(mc mqtt.Client, wg *sync.WaitGroup, cfg *models.Config, clientID int, payload string) {
	defer wg.Done()
	for m := 0; m <= cfg.MsgPerClientCount; m++ {
		fmt.Println("clientID: ", clientID, "sending on channelID: ", cfg.ChannelID, "messageID: ", m)
		if token := mc.Publish("channels/"+strconv.Itoa(cfg.ChannelID)+"/messages", cfg.QosLevel, false, payload); token.Wait() && token.Error() != nil {
			fmt.Printf("%s %s \n", strconv.Itoa(m), token.Error())
		}
	}
}
