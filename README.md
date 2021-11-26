[![Build Status](https://travis-ci.org/openservicebrokerapi/osb-checker.svg?branch=master)](https://travis-ci.org/openservicebrokerapi/osb-checker "Travis")

# osb-checker

An automatic checker to verify an Open Service Broker API implementation
against the
[specification](https://github.com/openservicebrokerapi/servicebroker).

# Project Status

This project should be considered **experimental**. You should validate the
results against the released
[specification](https://github.com/openservicebrokerapi/servicebroker). In the
case of any discrepancy, the specification should be considered correct.

# Usage

## Test

Deploy your own service broker to be tested. You can use the mock broker (see
below). Then, run the test suite:

```console
vim test/configs/config_mock.yaml # OPTIONAL: change this file if testing a custom service broker
go test -v ./test 
```

## Mock broker

This project provides a mock broker as reference implementation for enforcing
conformance test, and validated results of this mock broker should be
considered more reliable compared with other 3rd party brokers.

For deploying and using this mock broker, you should execute these commands
below:

```console
make
./bin/mockbroker
```

## Generate model and mock broker

All the model and mock broker frameworks will be generated from
service-catalog's
[swagger.yaml](https://raw.githubusercontent.com/openservicebrokerapi/servicebroker/master/swagger.yaml)
automatically. Here are some steps for developers to generate them:

```console
make autogenerated
```
