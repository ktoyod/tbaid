.PHONY:go/mod/tidy
go/mod/tidy:
	go mod tidy

.PHONY:install/tbaid
go/install/tbaid:
	GOBIN=$(CURDIR)/bin go install ./cmd/tbaid
