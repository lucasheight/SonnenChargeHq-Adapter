# Sonnen Batterie - ChargeHq Integration

This integration allows smart charging of your EV with [Sonnen Solar Batterie](https://sonnen.com.au/) and [ChargeHq](https://chargehq.net/).

There are implementations in both [Golang](./goService/) and [.NET core](./service/) with dockerfiles for both.

Also note, this integration assumes that Sonnen API v2 is used.

## Setup

### .NET

If running as .Net7 service update the [appsettings.json](./service/appsettings.json) with your relevant urls and keys and compile to your target
platform:

```json
  "Sonnen": {
    "BaseUrl": "http://localhost"
  },
  "ChargeHq": {
    "BaseUrl": "https://api.chargehq.net",
    "ApiKey": "YourApiKey",
    "RefreshMs":"120000" //optional- defaults to 2 mins
  }

```

### GO

Create a .env file in the goService folder with the following settings:

```env
 SONNEN__BASEURL=http://localhost
 CHARGEHQ__APIKEY=your_api_key
 CHARGEHQ__REFRESHMS=120000 (optional)

```

### Getting your Sonnen local address

- Open [https://find-my.sonnen-batterie.com/](https://find-my.sonnen-batterie.com/) on the local LAN. This page shows your local Sonnen IP address:

![image](./assets/sonnenDash.png)

- Use this as your BaseUrl in the Sonnen configuration setting.

### Getting your ChargeHq API key

- Log into [https://app.chargehq.net/](https://app.chargehq.net/).
- Click settings
- Select "My Equipment" -> Solar battery equipment.
- Select the "Push Api" item.
- Copy your API key into the ChargeHq ApiKey setting.

More details can be found here. [https://chargehq.net/kb/push-api](https://chargehq.net/kb/push-api).

## Docker

If you want to run a docker container, then, you can use the [publish.sh](./service/publish.sh) scripts to create a linux build/docker images.

### Exporting the docker image

To export the docker image, execute the [export.sh](./service/export.sh) script to add a tar archive in the [./service/dist](./service/dist/)
directory. I use this to run the container on my Synology NAS.

### Environment Settings

If using docker, add the following docker environment Variables for your configuration:

```docker
SONNEN__BASEURL=http://your_local_sonnen_address
CHARGEHQ__APIKEY=Your_Api_Key

```
