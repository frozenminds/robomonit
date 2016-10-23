# RoboMonit

Let robots monitor your apps and files.

## What's that?!?

Pipe your data into `RoboMonit` and let it control your hardware when defined patterns match.

`RoboMonit` is powered by [GoBot](https://gobot.io/), so you'll have support for a lot of [platforms](https://gobot.io/documentation/platforms/) like Raspberry Pi, Arduino, Beaglebone, Spark, Edison and why not drones. Blink LED's, move your Sphero or sound an alarm.

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
pins:
  green: "15"
  yellow: "16"
  blue: "17"
  red: "18"

patterns:
  green: (GET|HEAD)
  yellow: POST
  blue: (PATCH|PUT)
  red: DELETE
```

Then `pipe` the logs into `RoboMonit`:

```
ssh domain.com tail -f /var/log/nginx/access.log | robomonit pipe
```
