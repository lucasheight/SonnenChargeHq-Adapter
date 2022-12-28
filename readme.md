# Sonnen Batterie - ChargHq Integration

This integration allows smart charging of your EV with [Sonnen Solar Batterie](https://sonnen.com.au/) and [ChargeHq](https://chargehq.net/).

I have decided to use Microsoft .Net7 for this integration, but it would also be quite easy to port to other dev environments.

## Setup

If running just .Net7 updatate the [appsettings.json](./service/appsettings.json) with your relevant urls and keys and compile to your target
platform:

```json
  "Sonnen": {
    "BaseUrl": "http://localhost/"
  },
  "ChargeHq": {
    "BaseUrl": "https://api.chargehq.net/api/public/",
    "ApiKey": "YourApiKey"
  }

```

## Docker

If you want to run a docker container, they you can use the [publish.sh](./service/publish.sh) script to create a linux build/docker image.

### Exporting docker

To export the docker image, execute the [export.sh](./service/export.sh) script to add a tar archive in the [./service/dist](./service/dist/)
directory. I use this to run the container on my Synolgy NAS.

### Environment Settings

If using docker, add the following docker environment Variables for your configuration:

```docker
Sonnen__BaseUrl=http://your_local_sonnen_address/
ChargeHq__ApiKey=Your_Api_Key

```
