<p align="center">
    <img src="/static/logo.png?v=0.1.2" width="250" />
    <h3 align="center">Copper</h3>
    <p align="center">Kubernetes Monitoring and Observability Platform</p>
    <p align="center">
        <a href="https://github.com/clivern/copper/actions/workflows/api.yml">
            <img src="https://github.com/clivern/copper/actions/workflows/api.yml/badge.svg">
        </a>
        <a href="https://github.com/clivern/copper/actions/workflows/ui.yml">
            <img src="https://github.com/clivern/copper/actions/workflows/ui.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Copper/actions/workflows/release.yml">
            <img src="https://github.com/Clivern/Copper/actions/workflows/release.yml/badge.svg">
        </a>
        <a href="https://github.com/clivern/copper/releases">
            <img src="https://img.shields.io/badge/Version-0.1.2-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/clivern/copper">
            <img src="https://goreportcard.com/badge/github.com/clivern/copper?v=0.1.2">
        </a>
        <a href="https://godoc.org/github.com/clivern/copper">
            <img src="https://godoc.org/github.com/clivern/copper?status.svg">
        </a>
        <a href="https://github.com/clivern/copper/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>


## Documentation

#### Linux Deployment

Download [the latest copper binary](https://github.com/clivern/copper/releases). Make it executable from everywhere.

```zsh
$ export LATEST_VERSION=$(curl --silent "https://api.github.com/repos/clivern/copper/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/clivern/copper/releases/download/v{$LATEST_VERSION}/copper_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

Create the configs file `config.yml` from `config.dist.yml`. Something like the following:

```yaml
# App configs
app:
  # App name
  name: ${APP_NAME:-copper}
  # Env mode (dev or prod)
  mode: ${APP_MODE:-prod}
  # HTTP port
  port: ${API_PORT:-80}
  # Hostname
  hostname: ${API_HOSTNAME:-127.0.0.1}
  # TLS configs
  tls:
    status: ${API_TLS_STATUS:-off}
    crt_path: ${API_TLS_PEMPATH:-cert/server.crt}
    key_path: ${API_TLS_KEYPATH:-cert/server.key}

  # Global timeout
  timeout: ${API_TIMEOUT:-50}

  # Kubernetes Clusters
  clusters:
      -
          name: ${KUBE_CLUSTER_01_NAME:-production}
          inCluster: ${KUBE_CLUSTER_01_IN_CLUSTER:-false}
          kubeconfig: ${KUBE_CLUSTER_01_CONFIG_FILE:-/app/configs/production-cluster-kubeconfig.yaml}
      -
          name: ${KUBE_CLUSTER_02_NAME:-staging}
          inCluster: ${KUBE_CLUSTER_02_IN_CLUSTER:-false}
          kubeconfig: ${KUBE_CLUSTER_02_CONFIG_FILE:-/app/configs/staging-cluster-kubeconfig.yaml}

  # Log configs
  log:
    # Log level, it can be debug, info, warn, error, panic, fatal
    level: ${LOG_LEVEL:-info}
    # Output can be stdout or abs path to log file /var/logs/copper.log
    output: ${LOG_OUTPUT:-stdout}
    # Format can be json
    format: ${LOG_FORMAT:-json}
```

The run the `copper` with `systemd`

```zsh
$ copper server -c /path/to/config.yml
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
