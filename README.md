# backpack

A wrapper for pack which handle **crowley pack**'s environment variables in order to change directory, perform an user switch and then, execute the given command line...

Also, it helps reduce boilerplate code and configuration for developers and/or sysadmins.

> **NOTE**: You can execute multiple command using: `bash -c "cmd1 && cmd2"`

## Usage

```console
Usage: crowley-backpack command [args]
       crowley-backpack --version

User management and command invoker for crowley-pack build system.

Arguments:
  command      Command to execute
  args         Command's arguments

Options:
  -h, --help       Print usage and quits
  -v, --version    Print version information and quits
```

**Example:**

`crowley-backpack make foo`

## Testing

You can run tests for backpack by using this command:

```console
make test
```

Note that test cases are run in a **Docker container**, so you need to install `docker` first.

## Linting

backpack use [gometalinter](https://github.com/alecthomas/gometalinter).

### Install

```bash
go get -u github.com/alecthomas/gometalinter
gometalinter --install --update
```

### Usage

```console
make lint
```

## Inspirations

* [`gosu`](https://github.com/tianon/gosu)

## Dependencies

Name                           | License
-------------------------------|----------
github.com/opencontainers/runc | `Apache 2.0`

## License

This is Free Software, released under the terms of the [`GPL v3`](LICENSE).
