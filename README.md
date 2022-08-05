# sentry-sample

## setup sentry container

#### Copy .env-sample to .env

```bash
$ cp .env-sample .env
```

#### Generate secret key

```bash
$ docker-compose run --rm sentry config generate-secret-key
```

* 生成されたsecret-keyを.envのSENTRY_SECRET_KEYにセットする

#### Initialize database and create Account

```bash
$ docker-compose run --rm sentry upgrade
```

* 実行すると、以下DBマイグレーションが行われる（5分程がかかる）

  ```bash
  :
  > sentry:0470_org_saved_search
  > sentry:0471_global_saved_search_types
  > sentry:0472_auto__add_field_sentryapp_author
  ```

* 終わると、アカウントを作成されるか聞かれるので、作成を行う<br>
（Emailは適当なアドレスでOK、ブラウザ上からログインする時に利用する）

  ```bash
  Would you like to create a user account now? [Y/n]:
  ```

#### Start up

```bash
$ docker-compose up -d
```

#### Open Browser

* http://localhost:9000
