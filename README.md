# PSPD Lab gRPC

## LLM Server

Para subir o servidor do ollama execute o comando `docker compose up --build`. O servidor aceita requisições em `localhost:11434/api`, exemplo do corpo de uma requisição que pergunta o motivo do céu ser azul no endpoint `/generate`:

```
{
  "model": "llama3.2",
  "prompt": "Why is the sky blue?",
  "stream": false
}
```

Mais informações sobre a api do ollama na [documentação](https://github.com/ollama/ollama/blob/main/docs/api.md).