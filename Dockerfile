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

# Usar a mesma imagem do Go como base para a imagem final
FROM golang:1.22

# Definir o diretório de trabalho na imagem final
WORKDIR /app

# Instalar o Air na imagem final
RUN go install github.com/cosmtrek/air@latest

# Adicionar o diretório de binários do Go ao PATH na imagem final
ENV PATH="/go/bin:${PATH}"

# Copiar o binário compilado e os arquivos de código-fonte para a imagem final
COPY --from=builder /app/cms-school /cms-school
COPY . .

# Expõe a porta em que sua aplicação estará ouvindo
EXPOSE 3333

# Define o comando para executar o aplicativo usando o Air para hot reloading
ENTRYPOINT ["air"]