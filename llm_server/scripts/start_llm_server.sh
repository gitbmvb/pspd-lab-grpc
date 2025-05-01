#!/bin/bash

ollama serve &
sleep 5

# Pull models
echo "Pulling model ${OLLAMA_MODEL}..."
ollama pull ${OLLAMA_MODEL}

echo "Pulling model llama3.2:1b..."
ollama pull llama3.2:1b

# Espera o servidor terminar
wait $!