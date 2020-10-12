package main

type BrokerConfig struct {
  BrokerAddress string
  OutputTopic string
  InputTopic string
}

type SerialMessage struct {
  Command string
  Value string
}

type JSONMessage struct {
  Command string `json:"command"`
  Value   string `json:"value"`
}
