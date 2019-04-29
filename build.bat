@echo off
go build -ldflags "-X github.com/yeejlan/maru.BuildDir=%cd%" main.go