// APIs that interact with OBD

package obd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	conn *websocket.Conn
)
// ConnectToOBD connect to an OmniBOL Daemon (OBD).
func ConnectToOBD(nodeAddress string) (string, error) {

	// connection with websocket
	fmt.Println("NodeAddress is = ", nodeAddress)
	nodeAddress = "62.234.216.108:60020"

	u := url.URL{Scheme: "ws", Host: nodeAddress, Path: "/wstest"}

	fmt.Println("Connect to OBD node -> ", u.String())

	connection, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("Connection failed : ", err)
		return "", err
	}

	conn = connection;

	fmt.Println("Response: ", resp.Status)

	// Get response from OBD API
	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Failed to read message: ", err)
		return "", err
	}
	
	fmt.Println("Message is: " + string(message))

	// return message
	returnMsg := "Connection Succeeded"

	return returnMsg, nil
}

// ObdLogin login to an OmniBOL Daemon (OBD).
func ObdLogin(mnemonicWords string) (string, error) {
	
	if conn == nil {
		return "Connection is nil", nil
	}

	fmt.Println("Mnemonic Words is = ", mnemonicWords)
	mnemonicWords = "opera muffin option float guess bracket arrest snake correct business captain brass"

	// Message struct
	type Message struct {
		Type    int32  		 `json:"type"`
		Data    interface{}  `json:"data"`
	 }
	 
	type Mnemonic struct {
		Mnemonic string `json:"mnemonic"`
	}
	 
	// send message to obd
	mnemonic := Mnemonic{Mnemonic: mnemonicWords}
	msg := Message{Type: -102001, Data: mnemonic}
	msgBytes, err := json.Marshal(msg)

	log.Println(string(msgBytes))

	err = conn.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		return "", err
	}

	// Get response from OBD API
	_, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Failed to read message: ", err)
		return "", err
	}
	
	fmt.Println("Message is: " + string(message))

	// return message
	// returnMsg := "Connection Succeeded"

	return string(message), nil
}