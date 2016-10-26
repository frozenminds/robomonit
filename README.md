# RoboMonit

Let robots monitor your apps and files.

## What's that?!?

Pipe your data into `RoboMonit` and let it control your hardware when defined patterns match.

`RoboMonit` is powered by [GoBot](https://gobot.io/), so you'll have support for a lot of [platforms](https://gobot.io/documentation/platforms/) like *Raspberry Pi*, *Arduino*, *Beaglebone*, *Edison* and why not drones. Blink LED's, move your *Sphero* or sound an alarm.

### Requirements

- Golang for ARM: https://github.com/golang/go/wiki/GoArm
- A supported platform: https://gobot.io/documentation/platforms/

### Install

```
go install github.com/frozenminds/robomonit
```

### Configure

Create a new YAML configuration file named `robomonit.yaml` and place it under `$HOME` or `/etc`.

Sample configuration to match HTTP request methods (`GET` or `HEAD`, `POST`, `PATCH` or `PUT`, `DELETE`) and light up an LED:

```
platforms:
  raspi:
    direct-pin:
      green: "15"
      yellow: "16"
      blue: "17"
      red: "18"
patterns:
  raspi.direct-pin.green: (GET|HEAD)
  raspi.direct-pin.yellow: POST
  raspi.direct-pin.blue: (PATCH|PUT)
  raspi.direct-pin.red: DELETE
```

Then `pipe` the logs into `RoboMonit`:

```
ssh domain.com tail -f /var/log/nginx/access.log | robomonit monitor
```

There's convenience machine `MachineFactory` built in that configurable via the configuration file. It supports *Arduino*, *Beaglebone*, *C.H.I.P*, *Edison* and *Raspberry Pi* with *direct PIN* and *LED* drivers.

For more control, safety, drivers and so on, create your own `struct`. The sole purpose of `MachineFactory` is to easily configure and use  `RobotMonit`.

### Special Thanks

* [Golang](https://golang.org/) - The Go Programming Language
* [Viper](https://github.com/spf13/viper) - Go configuration with fangs!
* [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
* [Gobot](https://gobot.io/) - Golang Powered Robotics
* [RaspberryPi](https://www.raspberrypi.org/) - Credit card-sized single-board computers

