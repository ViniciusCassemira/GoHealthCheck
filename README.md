# Go Health Check

Programa simples escrito em Go que verifica URLs dos protocolos HTTP e HTTPS, retornando informações sobre a requisição, como:
- [x] Código de status HTTP (`status_code`)
- [x] Versão do protocolo HTTP (`protocol_version`)
- [x] Tempo de execução da requisição em segundos (`execution_time_in_seconds`)

## Como usar

Crie um arquivo de texto com uma URL por linha:
```
https://youtube.com
https://google.com
https://github.com
```

### Formato das URLs

**Exemplos válidos:**
```text
https://google.com
http://example.com
https://api.github.com
```

**Exemplos inválidos:**
```text
google.com
www.google.com
ftp://example.com
```

### Arquivo de configuração

O programa espera um arquivo `.json` com as opções de consultas que serão feitas em cada uma das URLs. Certifique-se de ter esse arquivo em sua máquina, olhe o modelo desse arquivo na raiz do projeto: `config.example.json`
```json
{
    "protocol_version": true,
    "status_code": true
}
```

### Flags

| Flag | Descrição |
|------|-----------|
| `--url-file` | Caminho para o arquivo com as URLs (obrigatório) |
| `--json-path` | Diretório onde o JSON será salvo (obrigatório) |
| `--json-config-file` | Caminho para o arquivo JSON de configuração das consultas (obrigatório) |

> **Nota:** o diretório informado em `--json-path` precisa existir antes da execução, pois o programa não o cria automaticamente

### Executando via Go

```bash
go run ./cmd --url-file exemplo.txt --json-path ./resultados --json-config-file config.example.json
```

### Executando a partir do binário

> Para agilizar o seu uso, acesse a área das releases e baixe o binário compatível com a sua máquina

Dê permissão de execução ao binário baixado (exemplo em linux)
```bash
chmod +x gohealthcheck-linux-amd64 
```

Agora, execute o binário passando as flags necessárias
```bash
./gohealthcheck-linux-amd64 --url-file /home/vinicius/Downloads/file.txt --json-path /home/vinicius/Downloads/output --json-config-file config.example.json
```

### Log

Durante a execução, o programa cria (ou abre, se já existir) o arquivo `gohealthcheck.log` no diretório atual e registra nele todas as operações e erros encontrados.

## Exemplo de saída

O programa imprime o JSON no terminal e também o salva em um arquivo no diretório informado. O arquivo é nomeado com o timestamp da execução no formato `2006-01-02_15-04-05.json`:
```json
{
  "total_results": 3,
  "execution_time_in_seconds": 1.234,
  "results": [
    {
      "url": "https://youtube.com",
      "execution_time_in_seconds": 0.456,
      "status_code": 200,
      "protocol_version": "HTTP/2.0"
    },
    {
      "url": "https://google.com",
      "execution_time_in_seconds": 0.312,
      "status_code": 200,
      "protocol_version": "HTTP/2.0"
    },
    {
      "url": "https://github.com",
      "execution_time_in_seconds": 0.289,
      "status_code": 200,
      "protocol_version": "HTTP/2.0"
    }
  ]
}
```
