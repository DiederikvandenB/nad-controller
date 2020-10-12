package main

import (
  "bufio"
  "flag"
  mqtt "github.com/eclipse/paho.mqtt.golang"
  filename "github.com/keepeye/logrus-filename"
  log "github.com/sirupsen/logrus"
  "go.bug.st/serial"
)

/**
 * TODO:
 *  1. Currently, all values in JSON messages are stringified. Lets create individual structs for the commands
 *  2. We now start two different MQTT clients, which shouldn't be necessary. Let's ensure we open just one connection
 *  3. Some of the filenames / function names are not ideally chosen. Let's pick better ones
 *     i.e. we can rename read & write to mqtt-to-device and device-to-mqtt
 *     then we can rename serializer to mqtt-parser and parser to device-parser
 *  4. Add tests
 */

func main() {
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
    filenameHook := filename.NewHook()
    filenameHook.Field = "file"
    log.AddHook(filenameHook)
    log.SetLevel(log.DebugLevel)
  }

  brokerConfig := BrokerConfig{
    BrokerAddress: *brokerAddress,
    OutputTopic:   *outputTopic,
    InputTopic:    *inputTopic,
  }

  serialPort := openPort(*serialPortAddress)

  go write(brokerConfig, serialPort)
  go read(brokerConfig, bufio.NewReader(serialPort))
  select {}
}

func openPort(serialPort string) serial.Port {
  port, err := serial.Open(serialPort, &serial.Mode{
    BaudRate: 115200,
  })

  if err != nil {
    log.Fatal(err)
  }

  log.WithField("serial-port", serialPort).Info("opened serial connection")

  return port
}

func startMQTT(config BrokerConfig) mqtt.Client {
  client := mqtt.NewClient(mqtt.NewClientOptions().AddBroker(config.BrokerAddress))

  if token := client.Connect(); token.Wait() && token.Error() != nil {
    log.Fatal(token.Error())
  }

  log.WithField("broker", config.BrokerAddress).Info("connected to mqtt broker")

  return client
}
