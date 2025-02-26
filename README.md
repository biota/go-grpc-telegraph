# go-grpc-telegraph
gRPC based transmission line

Creates communication channels between local devices (aka clients) and
remote services over gRPC.

## What is it for

A bidirectional secure communication channel over gRPC (device initiated),
which allows the device (aka the client) to stream events to a service
and additionally receive work from a service.

## How do I build it

The short answer is run make:

```shell
make
```

The `build` target (default for `all`) installs the dependencies to
compile and build protobuf code and then runs the golang build.

And to run a clean build, you will need to remove all the dependencies
and run a `make clean` before running `make` ...

```shell
make depclean; make clean; make
```

## How do I run the samples

### work in progress

## How do I test the code

To run all the tests, use the makefile `test` or `tests` target.

```shell
make test  #  == make tests
```

## How do I generate a test coverage report

```shell
make test
make coverage
```

## How do I lint the code

To run all the linter tests, use the `check` or `lint` makefile targets.

```shell
make lint   #  or check
```

## Issues

This is still work in progress ...
