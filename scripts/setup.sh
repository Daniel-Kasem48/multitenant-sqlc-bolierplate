#!/bin/sh

export GOBIN=$(pwd)/bin

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go get github.com/pressly/goose/v3
go install github.com/dgrijalva/jwt-go
go install github.com/lib/pq
go get github.com/labstack/echo
go get github.com/joho/godotenv
go get github.com/labstack/echo/v4/middleware
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgxpool@v5.6.0
go get github.com/go-playground/validator/v10
go get github.com/jackc/pgx/v5/pgtype
go get google.golang.org/grpc
go get github.com/google/uuid
go get go.uber.org/fx
go get github.com/lib/pq
go get github.com/dgrijalva/jwt-go
