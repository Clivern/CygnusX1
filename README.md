<p align="center">
    <img src="/static/logo.svg?v=0.2.0" width="200" />
    <h3 align="center">Peacock</h3>
    <p align="center">Modern Command Line Tool for Apache Kafka.</p>
    <p align="center">
        <a href="https://github.com/clivern/peacock/actions/workflows/cli.yml">
            <img src="https://github.com/clivern/peacock/actions/workflows/cli.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Peacock/actions/workflows/release.yml">
            <img src="https://github.com/Clivern/Peacock/actions/workflows/release.yml/badge.svg">
        </a>
        <a href="https://github.com/clivern/peacock/releases">
            <img src="https://img.shields.io/badge/Version-0.2.0-blue.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/clivern/peacock">
            <img src="https://goreportcard.com/badge/github.com/clivern/peacock?v=0.2.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/peacock">
            <img src="https://godoc.org/github.com/clivern/peacock?status.svg">
        </a>
        <a href="https://github.com/clivern/peacock/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-blue.svg">
        </a>
    </p>
</p>
<br/>

Peacock is a Modern Command Line Tool for Apache Kafka. It is easy to install and use. Here is some of the features:

- Consume messages on specific partitions between specific offsets.
- Display topic information (e.g., with partition offset and leader info).
- Modify consumer group offsets (e.g., resetting or manually setting offsets per topic and per partition).
- JSON output for easy consumption with tools like kp or jq.
- JSON input to facilitate automation via tools like jsonify.
- Support for TLS authentication.
- Cluster administration functions: Create & delete topics.

## Documentation

### Linux Installation

Download [the latest peacock binary](https://github.com/clivern/peacock/releases). Make it executable from everywhere.

```zsh
$ export LATEST_VERSION=$(curl --silent "https://api.github.com/repos/clivern/peacock/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/clivern/peacock/releases/download/v{$LATEST_VERSION}/peacock_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

### Mac Installation


### Usage

To run a local kafka cluster with `docker` & `docker-compose` for testing, you can use the following command:

```zsh
$ peacock cluster run [name] [port]

$ peacock cluster run local_clus1 3000
$ peacock cluster run local_clus2 3001
```

Please note that the above command requires both `docker` and `docker-compose`.

To destroy local clusters

```zsh
$ peacock cluster destroy [name]

$ peacock cluster destroy local_clus1
$ peacock cluster destroy local_clus2
```

To list all configured clusters

```zsh
$ peacock cluster list
```

To show cluster info

```zsh
$ peacock cluster show [name]

$ peacock cluster show local_clus1
```

To add a new remote cluster

```zsh
$ peacock cluster add [name]

$ peacock cluster add remote_cluster
```

To remove a configured cluster

```zsh
$ peacock cluster show [name]

$ peacock cluster remove remote_cluster
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Peacock is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/peacock/releases) for changelogs for each release version of Helmet. It contains summaries of the most noteworthy changes made in each release.

## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/peacock/issues

## Security Issues

If you discover a security vulnerability within Peacock, please send an email to [hello@clivern.com](mailto:hello@clivern.com)

## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.

## License

© 2022, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Peacock** is authored and maintained by [@Clivern](http://github.com/clivern).
