@echo off
go run -ldflags "-X github.com/yeejlan/maru.BuildDir=%cd%" main.go