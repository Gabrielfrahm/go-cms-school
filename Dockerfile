# Usar a imagem oficial do Go como base para a etapa de construção
FROM golang:1.22 AS builder

# Definir o diretório de trabalho no container
WORKDIR /app

# Copiar os arquivos do módulo Go para baixar as dependências
COPY go.mod ./
COPY go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o restante do código-fonte para o container
COPY . .

# Compilar o aplicativo para um binário estático
# Aqui especificamos o caminho exato do arquivo main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cms-school ./cmd/main.go

# Etapa final usa a imagem scratch por ser a mais leve e segura
FROM scratch

# Copiar o binário compilado para a nova etapa
COPY --from=builder /app/cms-school /

# Expõe a porta em que sua aplicação estará ouvindo
EXPOSE 3333

# Define o comando para executar o aplicativo
ENTRYPOINT ["/cms-school"]