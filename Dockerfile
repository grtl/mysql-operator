FROM alpine:latest
ADD mysql-operator /go/bin/mysql-operator
ADD artifacts artifacts
ENTRYPOINT ["/go/bin/mysql-operator"]
