<p align="center">
    <img src="/static/logo.svg?v=0.2.0" width="250" />
    <h3 align="center">Copper</h3>
    <p align="center">Modern Command Line Tool for Apache Kafka.</p>
    <p align="center">
        <a href="https://github.com/clivern/copper/actions/workflows/cli.yml">
            <img src="https://github.com/clivern/copper/actions/workflows/cli.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Copper/actions/workflows/release.yml">
            <img src="https://github.com/Clivern/Copper/actions/workflows/release.yml/badge.svg">
        </a>
        <a href="https://github.com/clivern/copper/releases">
            <img src="https://img.shields.io/badge/Version-0.2.0-blue.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/clivern/copper">
            <img src="https://goreportcard.com/badge/github.com/clivern/copper?v=0.2.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/copper">
            <img src="https://godoc.org/github.com/clivern/copper?status.svg">
        </a>
        <a href="https://github.com/clivern/copper/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-blue.svg">
        </a>
    </p>
</p>
<br/>

Copper is a Modern Command Line Tool for Apache Kafka. It is easy to install and use. Here is some of the features:


## Documentation

#### Linux Installation

Download [the latest copper binary](https://github.com/clivern/copper/releases). Make it executable from everywhere.

```zsh
$ export LATEST_VERSION=$(curl --silent "https://api.github.com/repos/clivern/copper/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/clivern/copper/releases/download/v{$LATEST_VERSION}/copper_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Copper is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/copper/releases) for changelogs for each release version of Helmet. It contains summaries of the most noteworthy changes made in each release.

## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/copper/issues

## Security Issues

If you discover a security vulnerability within Copper, please send an email to [hello@clivern.com](mailto:hello@clivern.com)

## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.

## License

© 2022, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Copper** is authored and maintained by [@Clivern](http://github.com/clivern).
