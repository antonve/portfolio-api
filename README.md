# Portfolio API

## Getting started

### 1. Setup environment

- Install [`direnv`](https://direnv.net/)
- Copy over the default environment: `$ cp .env{.sample,}`
- Go over the file and make sure the environment variables are correct for your env (eg. database url)
- Make sure to generate a JWT Secret (instructions are inside .env)
- Allow direnv `$ direnv allow`

## Commands

### Build project

```sh
$ make all
```

### Run project without tests
```sh
$ make run
```

### Generate files

```sh
$ make gen
```

### Run tests

```sh
$ make test
```
