# check-tester

a docker container used for consul health check based integration tests

building: `GOOS=linux go build; docker build .`

endpoints:

`{containerip}:8888/check` consul compatible health check

`{containerip}:8888/fail` make `/check` fail

`{containerip}:8888/pass` make `/check` pass
