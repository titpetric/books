# The 12 Factors of Docker & Go

This is a general roadmap what to cover within the book. It lists external
projects and references for the type of software which can be used to
satisfy individual guidelines of the 12-Factor app. The criteria for
inclusion will be based on idiomatic guidelines and the traction which
individual projects managed to get.

Leanpub: [12 Factor Applications with Docker and Go](https://leanpub.com/12fa-docker-golang)

## I. Codebase - One codebase tracked in revision control, many deploys

The obvious choice in regards to codebase today is git. There are two
aspects which are interesting here. One is the software being used
on the server side, and options for self-hosting (GitLab, Gogs) and the
other is development methodologies for teams.

- [x] Self hosted git option: [Gogs](https://gogs.io/)
- [x] Workflows for development with git - [Feature branches](https://www.atlassian.com/git/tutorials/comparing-workflows)

## II. Dependencies - Explicitly declare and isolate dependencies

There are a few package managers which are available for Go. Depending on your needs,
any might be fine. My personal favorite is [gvt](https://github.com/FiloSottile/gvt) and
I've seen [godep](https://github.com/tools/godep) many times in the wild.

In addition to a quick how-to and why in regards to vendoring, there should be some
things explained here in terms of our own dependencies and how to structure the project
in such a way, that reuse from other projects is possible.

## III. Config - Store config in the environment

- [x] [spf13/viper](https://github.com/spf13/viper)
- [x] [namsral/flag](https://github.com/namsral/flag)
- [x] [joho/godotenv](https://github.com/joho/godotenv)

The project `spf13/viper` will be covered due to terrific adption by several large projects.
The other packages are listed as simpler options which may be considered for a
nano-services approach, where only few configuration/environment variables are needed.

## IV. Backing services - Treat backing services as attached resources

Demonstrate use of common backing services.

- [x] MySQL [jmoiron/sqlx](https://github.com/jmoiron/sqlx)
- [x] Redis [garyburd/redigo](https://github.com/garyburd/redigo)
- ~~Minio/S3 [minio/minio-go](https://github.com/minio/minio-go)~~

## V. Build, release, run - Strictly separate build and run stages

There are a few considerations to make when publishing 12 Factor apps. While applications
with multiple files may be the norm, there's existing software which is used for packaging
these apps, and more. I'd like to cover a few things:

### Build

Set up and use a CI system which fully supports Docker.

- [x] Codeship CI
- ~~GitLab CI?~~
- ~~[Buildkite](https://buildkite.com/)~~

### Release

This is very variable and subject to changes.

- [x] Setting up a docker registry [docker/registry](https://docs.docker.com/registry/),
- ~~Using the GitLab container registry,~~
- [x] Registry review: [Amazon ECR](https://aws.amazon.com/ecr/), Google GCR, Docker Hub, Quay.io
- [x] Building your own release system (Codeship, release to Docker Hub and Github)

### Run

Subject to change.

- [x] Docker to run go applications, scaling with docker swarm,
- [x] Migrating your container(s) to the cloud (Digital Ocean `doctl`, in intro)
- [ ] set up a scalable docker swarm with doctl?

## VI. Processes - Execute the app as one or more stateless processes

The principle says that the application itself shouldn't keep a local state - a local cache or
in-memory values. For persistent data storage a database or a caching server like redis or memcached
should be used as a backing service. The intent of this guideline is to improve reliability of
the service in face of random outages. If individual apps don't have their own caches or data,
it means that you can tolerate the outage with low impact.

This chapter is closely related to "IX. Disposability" in the sense that there's just one
logical conclusion towards testing it - reap some of the running application processes in
order to discover possible problems in terms of caches and data storage.

## VII. Port binding - Export services via port binding

Exposing ports via `net/http` and Docker `-p` option is the baseline for this chapter.
There are possible areas of extending the principles behind it:

1. Exposing a multi-host service with docker swarm (networking),
2. Private network topology with an exposed reverse-proxy setup,
3. Multi-host Docker private networking

## VIII. Concurrency - Scale out via the process model

While not applicable to Go programs directly, due to it's strong concurrency model and ability to
scale it's processing to multiple CPU cores, it is worth to look at this option as a way to provide
redundancy via the process model, and with this also a graceful upgrade path when deploying new
application versions. By using Docker we can also satisfy basic process management requirements
set forth in the 12FA guidelines - respond to crashed processes and handle user initiated restarts
and shutdowns.

## IX. Disposability - Maximize robustness with fast startup and graceful shutdown

Much like the principle set out in "VI. Processes", the 12FA guidelines push home the notion
that each application may fail at any given moment, and that it should be resilient to these failures.

Several approaches to automating failure and testing this should be noted - one of the
most known is [Netflix Chaos Monkey](https://blog.codinghorror.com/working-with-the-chaos-monkey/).

> One of the first systems our engineers built in AWS is called the Chaos Monkey. The Chaos Monkey’s job is to randomly kill instances and services within our architecture. If we aren’t constantly testing our ability to succeed despite failure, then it isn’t likely to work when it matters most – in the event of an unexpected outage.

In terms of recovery, docker swarm and docker infrakit might be good options to start with, to
demonstrate how it's possible to tolerate failure in a clustered environment. There are some
secondary goals within this chapter in terms of how a program behaves in terms of it's workload.

## X. Dev/prod parity - Keep development, staging, and production as similar as possible

Even if we're not using Docker to keep these environments as similar as possible, we have
to account for the nature of the 12FA guidelines - each application is a microservice, which
should only define it's dependencies explicitly (12FA, item II.).

While I realize that some dependencies (e.g. `rsync`) are provided completely outside the Go app,
it is possible to have "pre-flight" check within the application itself which would report
possible issues (much like automake `./configure` does).

## XI. Logs - Treat logs as event streams

Output logs appropriately to stdout/stderr and have external tooling to review this output
when applicable. There are tools that work within the Docker ecosystem like [Rancher](http://rancher.com/),
and there are external tools like [Logstash](https://www.elastic.co/products/logstash) which may
provide insight to your logs.

- [x] Use Papertrail from Docker (stdout),
- [x] Papertrail as a backing-service

## XII. Admin processes - Run admin/management tasks as one-off processes
