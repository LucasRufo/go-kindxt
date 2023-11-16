# Go Kindxt

This project is a clone of the amazing [Kindxt](https://github.com/sergioprates/kindxt) made by [@sergioprates](https://github.com/sergioprates), but written in Go.

Kindxt is a CLI that wraps the [Kind](https://kind.sigs.k8s.io/) project, making it easy to setup a new local Kubernetes cluster and add common helm charts to it. Kindxt does all the job to bind the host port correctly on the `kindconfig.yaml`. Be aware that this project is just a clone and it's not ready to be used, if you really need the funcionality described here, use the real [Kindxt](https://github.com/sergioprates/kindxt) project. 

I'm still learning Go and had this ideia of making a small clone of a project that I already have some familiarity to learn new concepts. For this project I've used the [Cobra](https://github.com/spf13/cobra) library to create the CLI. 

## Required tools

- Docker 20.10.21
- Helm 3.13.1
- Kind 0.20.0

## Commands

To see all the commands on your terminal, you can use the `--help` command.

**To setup a new empty cluster:**

```bash
gokindxt create-cluster
```

**To add helm charts to the cluster you can use flags, for example:**

```bash
gokindxt create-cluster --mongodb
```

The above command will create your local cluster, install the MongoDB helm chart and bind the default MongoDB port (27017) to your host.

To see all the helm charts that can be installed, check the [charts](#charts) section. 

## Charts

| Chart | Flag | Host Port | Node Port |
|---|---|---|---|
| MongoDB | --mongodb | 27017 | 30001 |

## Next steps

- Include more commands, like `delete-cluster` and `update-cluster`.
- Include more helm charts options, like SQL Server, Istio and Redis.
- Refactor the helm charts installation so that we don't run all the commands in every installation.
- Write unit tests.
- Create a Github actions pipeline to deploy the binary on Github.