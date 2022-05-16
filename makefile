gql:
	go get github.com/99designs/gqlgen@v0.17.2
	go run github.com/99designs/gqlgen generate --config ./gqlgen.yml

server:
	go run main.go server

mock:
	mockgen -source pkg/interface/service.go -destination dist/mock/service_mock.go -package mock
	mockgen -source pkg/interface/repo.go -destination dist/mock/repo_mock.go -package mock
	mockgen -source pkg/interface/script.go -destination dist/mock/script_mock.go -package mock

sqlc:
	sqlc compile && sqlc generate

mc:
	migrate create -ext sql -seq -digits 3 -dir ./db/migrations $(n)

test:
	go test ./pkg/service/...