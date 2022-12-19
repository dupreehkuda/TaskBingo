# Task Bingo

> ***Disclaimer!*** This project was made for educational purposes. I would not use this architecture in production.

> ***Disclaimer#2!*** Version 1.0 is not released yet. 
> 
> Some things may change before the release. Some things are still in development.

## Idea

This is not an average bingo game. It's task bingo.

The concept is to think of 16 tasks you would like to accomplish every day to grow your productivity, and then
at the end of the day, you will fill the table with task numbers you've done to see if you earned any bingo.

## The Story

A friend of mine came up with this idea in early June'20 when covid lockdowns happened, and our productivity decreased. 
We wanted to solve this problem, so we made a list of 16 tasks we wanted to accomplish every day, and he made an app 
using p5.js and ngrok. It looked kinda quirky: 

![old bingo](./desk/oldBingo.png)

At the end of the lockdown, I lost the game, but we figured out how to work at home, stay productive and not go insane.

A lot has changed since then, and I still want to beat him but not in the game, but in app creation, as you already noticed. 
I wanted to do it in late May'22 with React.js [(you can check it here)](https://github.com/dupreehkuda/ReactiveBingo), 
but I understood it was not the right time for revenge.

But here it is, a SvelteKit app with a microservice backend written in Go and storage in Postgres and Redis 
live at [taskbingo.com](https://taskbingo.com)

## Architecture

As the disclaimer said, this project was made for educational purposes, so there is 2 databases and three microservices.
I've already learned a lot, but that's not the end, and I would not use this architecture in production.

![arch](./desk/TaskBingoArchitecture.jpg)

## Frontend

As for the frontend, I chose SvelteKit. I tried React but wanted to try Svelte as it is making a lot of noise nowadays.
And as I see, I was not mistaken about my choice.

Svelte components are neat and comprehensive. Working with them is definitely an interesting experience, but still, 
I'm not a frontend developer, and I need to work with it a bit more. To be continued...

## Backend

Backend consists of three microservices: 
 - **Game service** is the main service, the 'heart' of the game, 
 - **User service** is the service that is responsible for processing and storing user data
 - **Task service** is the service that is responsible for processing and storing tasks packs

### Game service

Game service is a 'heart' of the game

I have implemented a three layer architecture: handler -> processor -> taskClient/userClient.
Processor is a layer of business-logic, so it manages requests to task service and user service.

This architecture allows us to change each part as we want, and it will work while each part implements specific
[interface](./game-service/internal/interfaces/interfaces.go).

taskClient and userClient is layers of gRPC connection to task and user services

This service has a middleware to check correctness of JWT-tokens that authorize users.

### User service

I have implemented a three layer architecture: handler -> processor -> storage. 
(processor is a layer of business-logic)

This architecture allows us to change each part as we want, and it will work while each part implements specific 
[interface](./user-data-service/internal/interfaces/interfaces.go).

Handlers is a layer where all gRPC handlers live.

Storage uses [jackc/pgx](https://github.com/jackc/pgx) as a driver and as an interface for accessing Postgres.
For now ERD looks like this:
![erd](./desk/user_erd.png)

### Task service

Just like **user service** I have implemented a three layer architecture: handler -> processor -> storage. 
(processor is a layer of business-logic)

This architecture allows us to change each part as we want, and it will work while each part implements specific 
[interface](./task-data-service/internal/interfaces/interfaces.go).

Handlers is a layer where all gRPC handlers live.

Storage uses [go-redis](https://github.com/go-redis/redis) and [go-rejson](https://github.com/nitishm/go-rejson) libraries, 
because Redis storage has a RedisJSON module, so it stores JSON values.

## Deploy

All this is deployed on VDS and available at [taskbingo.com](https://taskbingo.com)

For microservices and databases I used [docker-compose](docker-compose.yml). 
When version 1.0 will be released, databases will host on other managed servers. 
Of course database containers are temporary and unsafe.

For the frontend I just use PM2 manager for Node.js app managing.

Nginx is used to route request properly. 
 - `/` route leads to Svelte app
 - `/api` route leads to game service
