# backpack

A wrapper for pack which handle Packer's environment variables in order to change directory, perform an user switch and then, execute all required command lines...

Also, it helps reduce boilerplate code and configuration for developers and/or sysadmins by using only a `yaml` file.

## Usage

```console
backpack [-c packer.yml]
```

## Configuration

```yaml
backpack:
  pre-hooks:
    - apt-get update
    - apt-get install foo
  execute:
    - make
  post-hooks:
    - make install
```

## Inspirations

* [`gosu`](https://github.com/tianon/gosu)

## Dependencies

Name                           | License
-------------------------------|----------
github.com/jawher/mow.cli      | `MIT`
github.com/opencontainers/runc | `Apache 2.0`
gopkg.in/yaml.v2               | `LGPL v3`

## License

This is Free Software, released under the terms of the [`GPL v3`](LICENSE).
