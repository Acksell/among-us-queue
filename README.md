# Among Us Queue

A discord bot for you and your friends to easily express their intent to play Among Us and organize games. The bot will notify you whenever enough players are available! [Click here to add it to your server](https://discordapp.com/api/oauth2/authorize?client_id=754424813321322566&permissions=21195856&scope=bot).

Example use:

```
!amonq queue

!amonq leave

!amonq view
```


## Running locally

To run your own version of this bot, just add your own `BOT_TOKEN` in the .env file and run the following docker commands:

`$ docker build -t among-us-queue .` 

`$ docker run --env-file .env among-us-queue`

Alternatively run: 

`$ go install`

`$ among-us-queue`

