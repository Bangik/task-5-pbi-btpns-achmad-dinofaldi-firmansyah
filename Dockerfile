FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o task-5-pbi-btpns-achmad-dinofaldi-firmansyah

ENTRYPOINT ["/app/task-5-pbi-btpns-achmad-dinofaldi-firmansyah"]