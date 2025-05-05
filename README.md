# PSPD Lab gRPC

## Como Rodar Localmente

Cada módulo deve ser instanciado em um processo diferente, com seu respectivo comando.

* Frontend

```
cd frontend
npm run dev
```

* Backend

```
cd backend
go run main.go
```

* PostgreSQL

```
cd db
docker compose up --build
```

* Ollama

```
cd llm_server
docker run -v ollama:/root/.ollama -p 11434:11434 llm_service
```

* DataService gRPC 

```
cd db
python3 data_service.py
```

* LLMService gRPC 

```
cd llm_server
python3 llm_service.py
```
