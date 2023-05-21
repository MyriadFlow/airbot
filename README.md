# Air-gateway

Bot and API Gateway to support AIR NFTs

## Requirements

A discord account with subscription to Midjourney. Add the account to your desired channel along with [midjourney-bot](https://docs.midjourney.com/docs/invite-the-bot)

## Installation

Configure .env file using the sample-env, then run through docker

```
docker-compose up --build -d
```

## Usage/Examples

#### To generate image with bot, use this command in the channel

```
/generate <prompt>
```

##

#### To upscale an image from given set of image :

```
!airbot upscale <index>
```

where <index> is the index of the image you want to upscale

##

#### To get variation of an image from the given images :

```
!airbot variation <index>
```

where <index> is the index of the image you want to get variations of
