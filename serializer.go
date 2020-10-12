package main

import (
  "fmt"
  "strconv"
  "strings"
)

func writePower(message JSONMessage) SerialMessage {
  return SerialMessage{
    Command: "Power",
    Value:   serializeBool(message.Value),
  }
}

func readPower(message SerialMessage) JSONMessage {
  return JSONMessage{
    Command: "power",
    Value:   parseBool(message.Value),
  }
}

func writeVolume(message JSONMessage) SerialMessage {
  decibel := getDecibels(message.Value)

  return SerialMessage{
    Command: "Volume",
    Value:   decibel,
  }
}

func readVolume(message SerialMessage) JSONMessage {
  // For some reason, when the amp starts up, it sends the volume as a percentage
  // Probably only when `Main.VolumeDisplayMode=Percent`
  if strings.Contains(message.Value, "%") {
    str := strings.Replace(message.Value, "%", "", 1)
    toFloat, _ := strconv.ParseFloat(str, 32)
    percentage := fmt.Sprintf("%.2f", toFloat/100)

    return JSONMessage{
      Command: "volume",
      Value:   percentage,
    }
  }

  decibel := message.Value
  percentage := getPercentages(decibel)

  return JSONMessage{
    Command: "volume",
    Value:   percentage,
  }
}

func writeMute(message JSONMessage) SerialMessage {
  return SerialMessage{
    Command: "Mute",
    Value:   serializeBool(message.Value),
  }
}

func readMute(message SerialMessage) JSONMessage {
  return JSONMessage{
    Command: "mute",
    Value:   parseBool(message.Value),
  }
}

func writeSource(message JSONMessage) SerialMessage {
  return SerialMessage{
    Command: "Source",
    Value:   message.Value,
  }
}

// It seems like the device does not publish a source change over the serial port when the change is initiated by the
// infrared remote.
func readSource(message SerialMessage) JSONMessage {
  return JSONMessage{
    Command: "source",
    Value:   message.Value,
  }
}

func parseBool(value string) string {
  if value == "On" {
    return "1"
  }

  return "0"
}

func serializeBool(value string) string {
  if value == "1" {
    return "On"
  }

  return "Off"
}
