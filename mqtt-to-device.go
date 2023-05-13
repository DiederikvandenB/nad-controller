package main

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"go.bug.st/serial"
	"strings"
	"sync"
	"time"
)

func mqttToDevice(serialPort serial.Port, address string, topic string) {

	clientOptions := mqtt.NewClientOptions().AddBroker(address)
	clientOptions.SetPingTimeout(30 * time.Second)
	clientOptions.SetKeepAlive(30 * time.Second)
	clientOptions.SetAutoReconnect(true)
	clientOptions.SetMaxReconnectInterval(10 * time.Second)
	clientOptions.SetOnConnectHandler(func(client mqtt.Client) {
		client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			jsonString := string(msg.Payload())
			serialMessage := jsonToSerialMessage(jsonString)

			log.WithFields(log.Fields{
				"json":  jsonString,
				"topic": topic,
			}).Trace("message received from mqtt")

		if serialMessage.Command == "" || serialMessage.Value == "" {
			log.WithField("json", jsonString).Debug("tried to send unsupported command to device")
			return
		}

			// And finally send the command to the device
			var mutex sync.Mutex
			sendCommand(serialPort, serialMessage, &mutex)
		})
	})
	client := mqtt.NewClient(clientOptions)
	MqttChannel <- client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.WithField("broker", address).Info("connected to mqtt broker")

	log.WithField("topic", topic).Info("listening for commands to send to device")

}

func jsonToSerialMessage(jsonString string) SerialMessage {
	var jsonMessage JSONMessage
	var serialMessage SerialMessage

	err := json.Unmarshal([]byte(jsonString), &jsonMessage)

	if err != nil {
		log.WithField("json", jsonString).Error("unable to unmarshal json string")
		return SerialMessage{}
	}

	switch jsonMessage.Command {
	case "power":
		serialMessage = writePower(jsonMessage)
	case "volume":
		serialMessage = writeVolume(jsonMessage)
	case "mute":
		serialMessage = writeMute(jsonMessage)
	case "source":
		serialMessage = writeSource(jsonMessage)
	}

	return serialMessage
}

func sendCommand(port serial.Port, message SerialMessage, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()

	// Format the command for the serial port
	command := strings.Join([]string{message.Command, message.Value}, "=")
	command = strings.Join([]string{"\r", "Main.", command, "\r\n"}, "")

	log.WithField("command", command).Debug("sending command to device")

	_, err := port.Write([]byte(command))
	if err != nil {
		log.WithFields(log.Fields{
			"command": command,
			"error":   err,
		}).Error("unable to send command to device")
	}

	// Wait before sending the next command
	// It takes a bit longer for sources to switch, so we take that into account
	if message.Command == "Source" {
		time.Sleep(2200 * time.Millisecond)
		return
	}

	time.Sleep(50 * time.Millisecond)
}
