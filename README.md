# Comunicação gRPC com Golang

Esse projeto tem como finalidade mostrar as formas de comunicação gRPC entre um *server* e um *client*

As comunicações utilizadas nesse projeto são:

- Comunicação entre server e client.
- Comunicação Stream do lado do server.
- Comunicação Stream do lado do client.
- Comunicação Stream bi-direcional.

Para executar o projeto, basta entrar na pasta do mesmo e executar o comando ```go run ./cmd/server/server.go``` para executar o server e ```go run ./cmd/client/client.go``` para executar o client.