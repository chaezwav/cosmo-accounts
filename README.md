
# COSMO Accounts

COSMO Accounts is a central account for websites or apps that wish to integrate it as an O'Auth method

## Prerequisites

* [Pkl v0.25.3](https://pkl-lang.org/)
* [Go v1.22.2](https://go.dev/learn/)
* [PostgreSQL v16.2](https://www.postgresql.org/download/)

## Run Locally

```bash
  git clone https://github.com/chaeyeonswav/cosmo-accounts.git
  cd cosmo-accounts

  go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest

  pkl-gen-go pkl/AppConfig.pkl --base-path github.com/chaeyeonswav/cosmo-accounts
```

Start the server

```bash
  go run *.go
```

## Configuration

Config files use the `pkl` library from Apple.

You can learn how to use `pkl` [here](https://pkl-lang.org/main/current/language.html)

The template is provided in the `pkl` directory with the base template being at `pkl/AppConfig.pkl`

You **MUST** provide a file in `pkl/dev` called `config.pkl`

```pkl
amends "../AppConfig.pkl"

postgres {
    auth {
        username = "[PostgreSQL Username]"
        password = "[PostgreSQL Password]"
    }
}
```

Assigning a new host or port is as simple as:

```pkl
ammends "../AppConfig.pkl"

host = "[Your host]"

port = [Your port as an int]
```

## Authors

* [@chaeyeonswav](https://www.github.com/chaeyeonswav)
