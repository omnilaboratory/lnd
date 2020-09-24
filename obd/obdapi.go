// APIs that interact with OBD

package obd

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lightningnetwork/lnd/obd/bean"
	"log"
	"net/url"
	"time"
)

func init()  {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

var (
	conn *websocket.Conn
)
// ConnectToOBD connect to an OmniBOL Daemon (OBD).
func ConnectToOBD(nodeAddress string) (string, error) {

	if conn == nil {
		// connection with websocket
		fmt.Println("NodeAddress is = ", nodeAddress)
		nodeAddress = "62.234.216.108:60020"
		// nodeAddress = "127.0.0.1:60020"

		u := url.URL{Scheme: "ws", Host: nodeAddress, Path: "/wstest"}

		fmt.Println("Connect to OBD node -> ", u.String())

		connection, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			fmt.Println("Connection failed : ", err)
			return "", err
		}
		conn = connection;
		go goroutine()
	}
	returnMsg := "Connection Succeeded"
	return returnMsg, nil
}

func goroutine() {
	sendHeartBeat()
	defer conn.Close()

	// read message
	for {
		if conn == nil {
			return
		}
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("socket to obd get err:", err)
			return
		}
		replyMessage := bean.ReplyMessage{}
		err = json.Unmarshal(message, &replyMessage)
		if err == nil {
			log.Println(replyMessage)
			switch replyMessage.Type {
			case -102001:
				log.Println("login in success")
			}
		}
	}
}

func sendHeartBeat()  {
	ticker := time.NewTicker(time.Minute * 2)
	defer ticker.Stop()

	defer func(ticker *time.Ticker) {
		if r := recover(); r != nil {
			log.Println("tracker goroutine recover")
			ticker.Stop()
			conn = nil
		}
	}(ticker)
	// heartbeat
	go func() {
		for {
			select {
			case t := <-ticker.C:
				if conn != nil {
					info := make(map[string]interface{})
					info["type"] = -102007
					info["data"] = t.String()
					bytes, err := json.Marshal(info)
					err = conn.WriteMessage(websocket.TextMessage, bytes)
					if err != nil {
						log.Println("HeartBeat:", err)
						return
					}
				} else {
					return
				}
			}
		}
	}()
}

// ObdLogin login to an OmniBOL Daemon (OBD).
func ObdLogin(mnemonicWords string) (string, error) {

	if conn == nil {
		return "Connection is nil", nil
	}

	fmt.Println("Mnemonic Words is = ", mnemonicWords)
	mnemonicWords = "opera muffin option float guess bracket arrest snake correct business captain brass"

	// send message to obd
	mnemonic := bean.Mnemonic{Mnemonic: mnemonicWords}
	msg := bean.Message{Type: -102001, Data: mnemonic}
	msgBytes, err := json.Marshal(msg)

	log.Println(string(msgBytes))

	err = conn.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		return "", err
	}
	return "", nil
}

// ObdOpenChannel launch a request to create a channel with someone else(Bob).
func ObdOpenChannel(nodeId, userId string, info bean.OpenChannelInfo) (string, error) {

	if conn == nil {
		return "Connection is nil", nil
	}

	// send message to obd
	mnemonic := bean.Mnemonic{Mnemonic: mnemonicWords}
	msg := bean.Message{Type: -102001, Data: mnemonic}
	msgBytes, err := json.Marshal(msg)

	log.Println("Sent Msg --> ", string(msgBytes))

	err = conn.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		return "", err
	}
	return "", nil
}

// ObdAcceptChannel counterparty replies to OpenChannel.
func ObdAcceptChannel(nodeId, userId string, info bean.AcceptChannelInfo) (string, error) {

	if conn == nil {
		return "Connection is nil", nil
	}

	// send message to obd
	mnemonic := bean.Mnemonic{Mnemonic: mnemonicWords}
	msg := bean.Message{Type: -102001, Data: mnemonic}
	msgBytes, err := json.Marshal(msg)

	log.Println("Sent Msg --> ", string(msgBytes))

	err = conn.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		return "", err
	}
	return "", nil
}
