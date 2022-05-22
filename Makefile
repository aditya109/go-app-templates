check_install:
	which swagger || GO111MODULE=off go install github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	GO111MODULE=off swagger generate spec -t=. -o ./api/swagger/swagger.yaml --scan-models