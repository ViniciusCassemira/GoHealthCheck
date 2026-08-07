# Go Health Check

Programa simples escrito em Go que verifica URLs dos protocolos HTTP e HTTPS, retornando informações sobre a requisição, como:
- [x] Informações HTTP (`http_info`): código de status (`status_code`) e versão do protocolo (`protocol_version`)
- [x] Informações DNS (`dns_info`): registros A, AAAA, CNAME, MX, NS e TXT
- [x] Tempo de execução da requisição em segundos (`execution_time`)

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
    "http_info": true,
    "dns_info": true
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

Durante a execução, o programa exibe no terminal uma barra de progresso com o andamento das verificações e salva o resultado em um arquivo JSON no diretório informado. O arquivo é nomeado com o timestamp da execução no formato `2006-01-02_15-04-05.json`:

```json
{
  "total_results": 1,
  "execution_time": 0.456,
  "search_options": {
    "http_info": true,
    "dns_info": true
  },
  "results": [
    {
      "url": "https://youtube.com",
      "execution_time": 0.456,
      "http_info": {
        "protocol_version": "HTTP/2.0",
        "status_code": 200
      },
      "dns_info": {
        "a_records": [
          "142.251.128.14"
        ],
        "aaaa_records": [
          "2800:3f0:4001:82e::200e"
        ],
        "cname_records": [
          "youtube.com."
        ],
        "mx_records": [
          "smtp.google.com."
        ],
        "ns_records": [
          "ns1.google.com."
        ],
        "txt_records": [
          "v=spf1 include:_spf.google.com ~all"
        ]
      }
    }
  ]
}
```

> **Nota:** quando uma consulta é desabilitada no arquivo de configuração, o campo correspondente é salvo vazio (`{}` para `http_info` e campos `null`/`[]` para `dns_info`). O campo `search_options` reflete as opções usadas na execução.
