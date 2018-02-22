FROM alpine:latest
ADD mysql-operator /go/bin/mysql-operator
ADD artifacts /go/bin/artifacts
ENTRYPOINT ["/go/bin/mysql-operator"]
