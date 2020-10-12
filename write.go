package main

import (
  mqtt "github.com/eclipse/paho.mqtt.golang"
  log "github.com/sirupsen/logrus"
  "go.bug.st/serial"
  "strings"
)

func write(brokerConfig BrokerConfig, port serial.Port) {
  client := startMQTT(brokerConfig)
  topic := brokerConfig.InputTopic

  log.WithField("topic", topic).Info("started listening mqtt ")

  client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
    json := string(msg.Payload())
    command, err := serializeMessage(json)

    if err != nil {
      log.WithField("json", json).Warn(err)
      return
    }

    // Add leading and trailing \r
    command = strings.Join([]string{"\r", "Main.", command, "\r"}, "")

    _, err = port.Write([]byte(command))
    if err != nil {
      log.Fatal(err)
    }
  })
}
