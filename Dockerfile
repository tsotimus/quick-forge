FROM golang:1.24

RUN apt-get update && \
    apt-get install -y bash git curl zsh && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o quickforge main.go

RUN echo '# dummy .zshrc' > /root/.zshrc

SHELL ["/bin/zsh", "-c"]
CMD ["zsh", "-i"]