# Basis-Image
FROM golang:1.21

# Arbeitsverzeichnis setzen
WORKDIR /app

# Module kopieren und installieren
COPY go.mod ./
RUN go mod tidy

# Restliche Dateien kopieren
COPY . .

# Go Programm bauen
RUN go build -o ollama-server

# Port freigeben
EXPOSE 8080

# Startbefehl
CMD ["./ollama-server"]
