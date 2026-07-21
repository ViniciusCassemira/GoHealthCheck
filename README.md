# Go Health Check

Programa simples em Go que verifica URLs e retorna algumas informações:

- [x] Status Code
- [x] Versão do Protocolo HTTP

## Como usar

Crie um arquivo de texto com uma URL por linha:

```
https://youtube.com
https://blog.cassemira.com
```

Execute o programa passando o caminho do arquivo:

```bash
go run ./cmd --path arquivo.txt
```

Ou use o binário pré-compilado:

```bash
./healthcheck --path arquivo.txt
```

## Build

```bash
go build -o healthcheck ./cmd
```
