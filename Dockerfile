FROM golang:1.18

ENV PORT 5000

WORKDIR /app

ADD . /app

RUN go mod tidy
RUN go mod verify

RUN go build -o /food

EXPOSE $PORT

# CMD [ "/food", "serve" ]
ENTRYPOINT [ "bash", "./devops/entrypoint.sh", "/food", "serve", "docker-config-template.yml" ]