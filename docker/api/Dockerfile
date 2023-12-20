FROM golang:1.21.3-alpine3.18 AS builder
RUN apk add --no-cache openssh
RUN mkdir ~/.ssh/
RUN ssh-keyscan git.mulk.net >> ~/.ssh/known_hosts
RUN apk add --no-cache git
RUN git config --global url."git@git.mulk.net:".insteadOf "https://git.mulk.net/"
WORKDIR /application

EXPOSE 80

#dev stage - start#
FROM builder AS dev
RUN go install github.com/cosmtrek/air@latest

COPY go.mod /application/go.mod
COPY go.sum /application/go.sum

ENV GOPRIVATE=git.mulk.net
ENV GONOPROXY=git.mulk.net

ENV VAULT_HOST=http://vault:8200
ENV VAULT_TOKEN=root

#RUN go mod download
CMD ["air", "-c", "docker/api/.air.api.toml"]
#dev stage - end#

FROM builder AS container_builder
ARG SSH_PRIVATE_KEY
RUN echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa
COPY . /application
RUN go build -o app /application/cmd/api/main.go

FROM container_builder AS test
CMD ["/application/app"]

FROM container_builder AS prod
CMD ["/application/app"]
