
# COSMO Accounts

COSMO Accounts is a central account for websites or apps that wish to integrate it as an O'Auth method

## Dependencies

* [Pkl v0.25.3](https://pkl-lang.org/)
* [Go v1.22.2](https://go.dev/learn/)
* [PostgreSQL v16.2](https://www.postgresql.org/download/)

## Run Locally

```bash
  git clone https://github.com/chaeyeonswav/cosmoAccounts.git
  cd cosmoAccounts

  go get
  go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest

  pkl-gen-go pkl/AppConfig.pkl --base-path github.com/chaeyeonswav/cosmoAccounts
```

Start the server

```bash
  go run *.go
```

## Authors

* [@chaeyeonswav](https://www.github.com/chaeyeonswav)
