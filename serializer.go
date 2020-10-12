package main

import (
  "encoding/json"
  "errors"
  "strings"
)

func serializeMessage(json string) (string, error) {
  // First, we transform the JSON string into a JSONMessage struct
  fromJSON, err := unmarshallJSON(json)

  if err != nil {
    return "", errors.New("could not unmarshall JSON")
  }

  // Then we fill a new SerialMessage based on the JSONMessage.Command
  var message SerialMessage
  switch fromJSON.Command {
  case "power":
    message = writePower(fromJSON)
  case "volume":
    message = writeVolume(fromJSON)
  case "mute":
    message = writeMute(fromJSON)
  case "source":
    message = writeSource(fromJSON)
  }

  if message.Command == "" {
    return "", errors.New("unsupported write command")
  }

  // Then we convert the SerialMessage struct to a string
  return strings.Join([]string{message.Command, message.Value}, "="), nil
}

func writePower(message JSONMessage) SerialMessage {
  return SerialMessage{
    Command: "Power",
    Value:   serializeBool(message.Value),
  }
}

func writeVolume(message JSONMessage) SerialMessage {
  decibel := getDecibels(message.Value)

  return SerialMessage{
    Command: "Volume",
    Value:   decibel,
  }
}

func writeMute(message JSONMessage) SerialMessage {
  return SerialMessage{
    Command: "Mute",
    Value:   serializeBool(message.Value),
  }
}

func writeSource(message JSONMessage) SerialMessage {
  return SerialMessage{
    Command: "Source",
    Value:   message.Value,
  }
}

func unmarshallJSON(str string) (JSONMessage, error) {
  var message JSONMessage

  err := json.Unmarshal([]byte(str), &message)

  if err != nil {
    return message, err
  }

  return message, nil
}

func serializeBool(value string) string {
  if value == "1" {
    return "On"
  }

  return "Off"
}
