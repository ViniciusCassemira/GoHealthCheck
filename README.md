# Go Health Check

Programa simples em Go que verifica URLs e retorna:

- [x] Código de status HTTP (`status_code`)
- [x] Versão do protocolo HTTP (`protocol_version`)

## Como usar

Crie um arquivo de texto com uma URL por linha:

```
https://youtube.com
https://google.com
https://github.com
```


### Flags

| Flag | Descrição |
|------|-----------|
| `--url-file` | Caminho para o arquivo com as URLs (obrigatório) |
| `--json-path` | Diretório onde o JSON será salvo (obrigatório) |


### Executando via Go

```bash
go run ./cmd --url-file exemplo.txt --json-path ./resultados
```

### Executando a partir do binário
> Para agilizar o seu uso, acesse a área das releases e baixe o binário compatível com a sua máquina


```bash
go build -o healthcheck ./cmd
./healthcheck --url-file exemplo.txt --json-path ./resultados
```


## Exemplo de saída

O programa gera um arquivo JSON no diretório informado trazendo o resultado das consultas:

```json
{
  "total_results": 3,
  "results": [
    {
      "url": "https://youtube.com",
      "status_code": 200,
      "protocol_version": "HTTP/2.0"
    },
    {
      "url": "https://google.com",
      "status_code": 200,
      "protocol_version": "HTTP/2.0"
    },
    {
      "url": "https://github.com",
      "status_code": 200,
      "protocol_version": "HTTP/2.0"
    }
  ]
}
```