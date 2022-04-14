FROM golang:1.18

ENV PORT 5000
ARG UID=1000
ARG GID=1000

RUN groupadd -r --gid $GID user && useradd -r --uid $UID -g user user

WORKDIR /app

ADD . /app

RUN go mod tidy
RUN go mod verify

RUN go build -o /food

EXPOSE $PORT

# CMD [ "/food", "serve" ]
ENTRYPOINT [ "bash", "./devops/entrypoint.sh", "/food", "serve", "docker-config-template.yml" ]