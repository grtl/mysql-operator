FROM alpine:latest
ADD mysql-operator /go/bin/mysql-operator
ENTRYPOINT ["/go/bin/mysql-operator"]
