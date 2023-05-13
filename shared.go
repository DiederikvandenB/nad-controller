package main

import mqtt "github.com/eclipse/paho.mqtt.golang"

// Declare a shared channel
var MqttChannel = make(chan mqtt.Client)
