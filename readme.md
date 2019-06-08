# audit-services-cli-plugin

A CF CLI Plugin to list service instances details

## Installation

* Download the latest plugin from the [release section](https://github.com/oahcran/audit-services-cli-plugin/releases) of this repository. 
* Install the plugin with `cf install-plugin <path_to_binary>`. Use `-f` flag to uninstall existing plugin if any and install the new one.

**macOS**: `cf install-plugin -f https://github.com/oahcran/audit-services-cli-plugin/releases/download/v0.1.0/audit-services_0.1.0.osx`

## Usage

`$ cf audit-service-instances [--csv]`

```
$ cf audit-service-instances
+------------+-------------+-------------------+--------------+-----------------------+--------------------------------------+-----------------+
|    ORG     |    SPACE    |   SERVICE NAME    | SERVICE PLAN | SERVICE INSTANCE NAME |        SERVICE INSTANCE GUID         | BOUND APP COUNT |
+------------+-------------+-------------------+--------------+-----------------------+--------------------------------------+-----------------+
| pivot-cran | development | app-autoscaler    | standard     | svc-autoscaler        | afa00d7f-0849-4c31-89c5-27472359c2c3 |               1 |
| pivot-cran | development | metrics-forwarder | unlimited    | metrics-forwarder     | a9caa104-92d7-43f7-9657-9b62baa7b589 |               0 |
| pivot-cran | development | p-config-server   | standard     | service-config        | acb8ce1a-f2f1-4116-b7b3-c8c919adfe9f |               0 |
| pivot-cran | development | p-config-server   | standard     | config-server         | fd8fa63b-7272-4cbc-ad72-ad098ba125c0 |               1 |
| pivot-cran | development | p-rabbitmq        | standard     | rabbitmq-service      | 60f8edf8-001d-4938-84ce-8aa0d443073d |               1 |
| pivot-cran | development | p.redis           | cache-small  | svc-redis             | 55eece71-9cbc-4dac-94fb-aef574242dbf |               2 |
+------------+-------------+-------------------+--------------+-----------------------+--------------------------------------+-----------------+

```

## Build

compile

```
go get https://github.com/oahcran/audit-services-cli-plugin
cd $GOPATH/src/github.com/oahcran/audit-services-cli-plugin
./build-binaries.sh
```

use [counterfeiter v6](https://github.com/maxbrunsfeld/counterfeiter) to generate test doubles for `CfClient` interface

```
$ counterfeiter -generate
Writing `FakeCfClient` to `cffakes/fake_cf_client.go`... Done
```