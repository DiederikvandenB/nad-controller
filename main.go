package main

import (
  "bufio"
  "flag"
  mqtt "github.com/eclipse/paho.mqtt.golang"
  filename "github.com/keepeye/logrus-filename"
  log "github.com/sirupsen/logrus"
  "go.bug.st/serial"
)

type SerialMessage struct {
  Command string
  Value string
}

type JSONMessage struct {
  Command string `json:"command"`
  Value   string `json:"value"`
}

/**
 * TODO:
 *  1. Currently, all values in JSON messages are stringified. Lets create individual structs for the commands
 *  2. Add tests
 */

func main() {
  trace := flag.Bool("trace", false, "enables trace logs")
  debug := flag.Bool("debug", false, "enables debug logs")
  serialPortAddress := flag.String("serial-port-address", "/dev/ttyUSB0", "the serial port address")
  brokerAddress := flag.String("broker-address", "tcp://127.0.0.1:1883", "the mqtt broker address")
  outputTopic := flag.String("output-topic", "", "topic to send device updates to")
  inputTopic := flag.String("input-topic", "", "topic to subscribe to for controlling the device")
  flag.Parse()

  if *outputTopic == "" || *inputTopic == "" {
    log.Fatal("Please provide an output topic and an input topic for the mqtt broker")
  }

  log.SetFormatter(&log.TextFormatter{})
  if *debug == true {
    log.SetLevel(log.DebugLevel)
  }

  if *trace == true {
    filenameHook := filename.NewHook()
    filenameHook.Field = "file"
    log.AddHook(filenameHook)
    log.SetLevel(log.TraceLevel)
  }

  serialPort := openSerialPort(*serialPortAddress)
  mqttClient := startMQTT(*brokerAddress)

  go mqttToDevice(serialPort, mqttClient, *inputTopic)
  go deviceToMqtt(bufio.NewReader(serialPort), mqttClient, *outputTopic)
  select {}
}

func openSerialPort(serialPort string) serial.Port {
  port, err := serial.Open(serialPort, &serial.Mode{
    BaudRate: 115200,
  })

  if err != nil {
    log.Fatal(err)
  }

  log.WithField("serial-port", serialPort).Info("opened serial connection")

  return port
}

func startMQTT(address string) mqtt.Client {
  client := mqtt.NewClient(mqtt.NewClientOptions().AddBroker(address))

  if token := client.Connect(); token.Wait() && token.Error() != nil {
    log.Fatal(token.Error())
  }

  log.WithField("broker", address).Info("connected to mqtt broker")

  return client
}
