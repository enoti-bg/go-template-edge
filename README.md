# An opinionated GO Service template

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/enoti-bg/go-template-edge/master/LICENSE)


## Why?

We prefer libraries to frameworks and after being asked repeatedly "What do we use for X w/ GO?" we've decided to create this repository and try to maintain a small up to date reflection of our production golang stack.

## Requirements

The template can be consumed with [Cookiecutter](https://github.com/cookiecutter/cookiecutter).
Current golang target is **1.15**

## Usage

Once the package is installed, it is enough to

```shell
$ git clone https://github.com/enoti-bg/go-template-edge.git
```

and then start the creation of a new template with

```shell
$ cookiecutter <path-to-cloned-repo>
```

During creation the template exposed 4 parameters:
* `project_name` (default: hello-world) - This is the name of the project and accidentally the created binary output
* `gomodule_uri` (default: hello-world) - This is the path to the respective module - e.g. "github.com/enoti-bg/from-go-template"
* `goproxy_uri` (default: "")- If goproxy is being used, please indicate the full uri with protocol + port - e.g. "<protocol>://<domain-and-path>:<port>"
* `service_port` (default: 8889)- Port on which the service will be exposed by default. (adjustable by environment variable)


The Makefile shipping with the template supports four main targets:
* **go-deps** - Syncs go.(mod/sum) with deps.
* **go-build** - Packages an executable.
* **go-test** - Executes all tests included with the module.
* **go-run** - Executes current codebase in **development** environment


## Content

The template currently ships with the following libraries:

* [Chi](https://github.com/go-chi/chi): v1.5.1 - Opinionated minimal router library.
* [Chi/Render](https://github.com/go-chi/render): v1.0.1 - Response helpers for Chi.
* [Zerolog](https://github.com/rs/zerolog): v1.20.0 - Performant structured logging.
* [Ozzo-Validation](https://github.com/go-ozzo/ozzo-validation) v4.3.0 - Structural and value validation library.
* [testify](https://github.com/stretchr/testify): v1.4.0 - Testing helpers.
* [Cobra](https://github.com/spf13/cobra): v1.1.1 - CLI creation library.
* [Viper](https://github.com/spf13/viper): v1.7.0  - Environment and configuration loading.

Coming soon(for certain values of soon, because cookiecutter optional implementation physically hurts us):
* (optional) [GraphQL](https://github.com/graph-gophers/graphql-go)
* (optional) [SQL Storage w/ Upper](https://github.com/upper/db)
* (optional) [GRPC](https://github.com/grpc/grpc-go)
* (optional) Several common middlewares

After creating a new service from the template, you will have a working API
with a single demo entity, memory repository, chi router & endpoint, logger
and application configuration.
Running a service would produce an output:

```shell
$ make go-run
SERVICE_ENV=development GOPROXY= SERVICE_PORT=8888 SERVICE_LOG=debug go run main.go server
{"level":"debug","time":"<timestamp>","message":"Starting server on port: [8888]"}
```

Once running it is possible to call the demo endpoint:

```shell
curl -X POST --data-binary '{"label":"je suis label"}' localhost:8888/api/demo
```

Which will create a demo record with label `je suis label` that can now be fetched under the demo key

```
curl -X GET localhost:8888/api/demo/demo
```

### Contributors

* [BB](https://github.com/bbsbb)
* [PA](https://github.com/pepi1707)
* [MP](https://github.com/merilinpisina)

### License

Copyright © 2021 Enoti.BG

Distributed under the MIT License
