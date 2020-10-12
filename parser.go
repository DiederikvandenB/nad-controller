package main

import (
  "encoding/json"
  "fmt"
  "regexp"
  "strconv"
  "strings"
)

// Parses incoming messages
func parseMessage(msg string) (string, error) {
  // The device also send other messages that are not caught by this regex
  // (like Main.IR.LearningDevice=Amp), which is done on purpose. We don't need those responses
  pattern := regexp.MustCompile(`Main.([A-Za-z]+)=(.*)`)
  matches := pattern.FindStringSubmatch(msg)

  if matches == nil {
    return "", fmt.Errorf("received unknown message from device: %s", msg)
  }

  message := SerialMessage{
    Command: matches[1],
    Value: matches[2],
  }

  var command string
  switch matches[1] {
  case "Power":
    command = readPower(message)
  case "Volume":
    command = readVolume(message)
  case "Mute":
    command = readMute(message)
  case "Source":
    command = readSource(message)
  default:
    return "", fmt.Errorf("unknown command received from device: %s=%s", message.Command, message.Value)
  }

  return command, nil
}

func readPower(message SerialMessage) string {
  state := parseBool(message.Value)

  return toJSON("power", state)
}

func readVolume(message SerialMessage) string {
  // For some reason, when the amp starts up, it sends the volume as a percentage
  // Probably only when `Main.VolumeDisplayMode=Percent`
  if strings.Contains(message.Value, "%") {
    str := strings.Replace(message.Value, "%", "", 1)
    toFloat, _ := strconv.ParseFloat(str, 32)
    vol := fmt.Sprintf("%.2f", toFloat / 100)

    return toJSON("volume", vol)
  }

  decibel := message.Value
  percentage := getPercentages(decibel)

  return toJSON("volume", percentage)
}

func readMute(message SerialMessage) string {
  state := parseBool(message.Value)

  return toJSON("mute", state)
}

func readSource(message SerialMessage) string {
  state := message.Value

  return toJSON("source", state)
}

func toJSON(command string, value string) string {
  output := JSONMessage{
    Command: command,
    Value: value,
  }

  jsonData, _ := json.Marshal(output)

  return string(jsonData)
}

func parseBool(value string) string {
  if value == "On" {
    return "1"
  }

  return "0"
}
