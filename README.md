# NAD Controller

This software allows you to control your NAD amplifier over the RS232 port. Since I do not own all NAD amplifiers (obviously), I am unable to test which models work. Currently, only the NAD C368 is supported, but it is possible that similar models work too.

Upon starting, it opens a serial port and MQTT connection. Messages sent by the device are published to the MQTT broker for further usage in other apps. Controlling the device is possible by sending messages in JSON format to the specified MQTT topic.

**Note**: it is required to set your amplifier to display the volume as a percentage rather than as decibels. This can be done in the menu of the device.

## Usage
Run the bin with flag `-h` to see the possible options.

``` shell script
$ ./nad-controller -h

    Usage of nad-controller:
      -broker-address string
            the mqtt broker address (default "tcp://127.0.0.1:1883")
      -debug
            enables debug logs
      -input-topic string
            topic to subscribe to for controlling the device
      -output-topic string
            topic to send device updates to
      -serial-port-address string
            the serial port address (default "/dev/ttyUSB0")
      -trace
            enables trace logs
```

**Note**: `-input-topic` and `-output-topic` are required flags.

## Supported commands
The device outputs a bunch of commands on the serial port we do not use. Currently, only these commands are supported:

1. `{"command": "power", "value": "1"}`
2. `{"command": "volume", "value": "0.5"}`
3. `{"command": "mute", "value": "1"}`
4. `{"command": "source", "value": "3"}`

**Note**: values are always stringified (for now, see todo in main.go).

## Calculating the volume
Even though the volume is shown as a percentage on the device's display, the serial port still receives the volume in decibels. Since these are logarithmic, volume.go is a simple way to map these two values.

Surely there's a simple formula that we can use to do this conversion. Please let me know if you know how to do this.

## Disclaimer
This is a heavily opinionated piece of software currently, since it's tailor-made for my setup. It is of course possible to generalize the code so it works for more people. Please do open an issue or a PR. 
