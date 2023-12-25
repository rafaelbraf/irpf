FROM golang:latest

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
ENV PORT=8081

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 5432
EXPOSE 8081

COPY . .

RUN go build -o irpf-api irpf-ws/cmd/app/.

CMD [ "./irpf-api" ]