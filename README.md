# sentry-sample

## setup sentry container

Generate secret key

```bash
$ docker-compose run --rm sentry config generate-secret-key
```

Initialize database and create Account

```bash
$ docker-compose run --rm sentry upgrade
```

Start up

```bash
$ docker-compose up -d
```

Open Browser

* http://localhost:9000
