#!/bin/bash

ollama serve &
sleep 5

# Pull models
echo "Pulling model llama3.2..."
ollama pull llama3.2

echo "Pulling model llama3.2:1b..."
ollama pull llama3.2:1b

# Espera o servidor terminar
wait $!