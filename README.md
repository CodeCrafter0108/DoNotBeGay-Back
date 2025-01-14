# DoNotBeGay backend

## Gateway

Documentation - [Here](https://134.209.194.104/v1/docs/index.html)

- API for upload tokens images
- API for getting tokens images

## Randomizer

- Generate random UINT256 number
- Cron call for randomize

## Example for migrations

migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users

## See Makefile for more commands

```bash
make help
```
