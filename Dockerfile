# Etapa de build do frontend
FROM node:16.20.2 as frontend

WORKDIR /app

# Copia apenas o package.json para fazer cache do npm install
COPY /web/package*.json ./

RUN npm install

COPY /web .

RUN npm run build

# Etapa de build do backend (Go) usando uma imagem base do Go no Ubuntu
FROM golang:1.23.1 as backend

WORKDIR /app

# Copia arquivos go.mod e go.sum primeiro para fazer cache do download de dependências
COPY go.mod go.sum ./

RUN go mod download

# Copia o restante dos arquivos
COPY main.go ./main.go
COPY config ./config
COPY internal ./internal

# Copia o build do frontend para o backend
COPY --from=frontend /app/dist ./web/dist

RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o server ./main.go

# Etapa final usando Alpine para a imagem leve de produção
FROM alpine:3.18

WORKDIR /app

# Copia o binário Go e o build do frontend
COPY --from=backend /app/server .
COPY --from=backend /app/web/dist ./web/dist

EXPOSE 8080

CMD ["./server"]
