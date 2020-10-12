package main

import (
  "encoding/json"
  mqtt "github.com/eclipse/paho.mqtt.golang"
  log "github.com/sirupsen/logrus"
  "go.bug.st/serial"
  "strings"
)

func mqttToDevice(serialPort serial.Port, client mqtt.Client, topic string) {
  log.WithField("topic", topic).Info("listening for commands to send to device")

  client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
    jsonString := string(msg.Payload())
    serialMessage := jsonToSerialMessage(jsonString)

    if serialMessage.Command == "" {
      log.WithField("json", jsonString).Debug("tried to send unsupported command to device")
      return
    }

    // Format the command for the serial port
    serialCommand := strings.Join([]string{serialMessage.Command, serialMessage.Value}, "=")
    serialCommand = strings.Join([]string{"\r", "Main.", serialCommand, "\r"}, "")

    // And finally send the command to the device
    _, err := serialPort.Write([]byte(serialCommand))
    if err != nil {
      log.WithFields(log.Fields{
        "serialCommand": serialCommand,
        "error": err,
      }).Error("unable to send command to device")
    }
  })
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
