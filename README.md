# Pricing Engine

Flight pricing engine developed for SnappTrip Summer Bootcamp.

## Requirements
To run this app you need `docker` and `docker-compose`.

## Installation
Just clone the repository and use the following command to run it:

```
docker-compose up
```

It will Listen and serve HTTP on localhost:8080

## Endpoints
| Method | Endpoint | Description |
| --- | --- | --- |
| POST | /createRule | It takes an array of rules and inserts them into the database. |
| GET | /changePrice | It takes an array of tickets and returns the price with the highest profit for each ticket. |

There are some [examples](https://github.com/arrshad/pricing-engine/tree/master/examples) for each endpoint.


## Built with
- [Gin](https://github.com/gin-gonic/gin) - I used it to handle requests and their responses.
- [Gorm](https://github.com/go-gorm/gorm) - It made working and querying the database easier.
- [Go Redis](https://github.com/go-redis/redis) - To communicate with Redis.

## License

This software is licensed under the [MIT](https://github.com/arrshad/pricing-engine/tree/master/LICENSE)
