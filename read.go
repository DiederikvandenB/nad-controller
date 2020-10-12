package main

import (
  "bufio"
  log "github.com/sirupsen/logrus"
  "strings"
)

// Delimit incoming messages from the serial port
func read(brokerConfig BrokerConfig, reader *bufio.Reader) {
  for {
    message, err := reader.ReadString('\r')
    if err != nil {
      log.Fatal(err)
    }

    // Convert the raw message to JSON and publish it
    command := onMessage(message)
    if command != "" {
      publishMessage(brokerConfig, command)
    }
  }
}

// Handler for new messages
func onMessage(message string) string {
  // Remove the trailing \r and skip if the message is empty
  message = strings.Replace(message, "\r", "", 1)
  if message == "" {
    return ""
  }

  log.WithField("message", message).Debug("raw msg received from device")

  command, err := parseMessage(message)
  if err != nil {
    // We do not support all commands, so it is expected that we see errors here
    return ""
  }

  return command
}

// Publish incoming messages to MQTT
func publishMessage(brokerConfig BrokerConfig, command string) {
  client := startMQTT(brokerConfig)
  topic := brokerConfig.OutputTopic

  token := client.Publish(topic, 0, false, command)
  token.Wait()

  client.Disconnect(250)
}
