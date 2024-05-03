# Especifica uma imagem base
FROM golang:1.22.2

# Cria um diretório 'app' para conter o código-fonte do seu app
WORKDIR /app

# Copia tudo do diretório raiz para /app no contêiner
COPY . .

# Instala as dependências do Go
RUN go mod download

# Compila seu app com configuração opcional
RUN go build -o app

# Indica a porta de rede que o seu contêiner escuta
EXPOSE 3002

# Especifica o comando executável que roda quando o contêiner é iniciado
CMD ["./app"]