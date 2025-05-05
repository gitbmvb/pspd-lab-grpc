package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"backend/grpc_services"
	"log"
	"time"
)

const (
	maxRetries         = 30               // Número máximo de tentativas
	retryInterval      = 5 * time.Second  // Intervalo entre tentativas
	defaultModel       = "llama3.2"       // Modelo padrão
	modelDownloadTimeout = 10 * time.Minute // Timeout para download do modelo
)

func AskHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Validação do método HTTP
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Parsing do payload
	var payload struct {
		Input string `json:"input"`
		Model string
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	// 3. Define modelo padrão se não fornecido
	if payload.Model == "" {
		payload.Model = defaultModel
	}

	// 4. Verifica saúde do modelo e baixa se necessário
	if err := ensureModelReady(r.Context(), payload.Model); err != nil {
		log.Printf("Falha ao preparar modelo: %v", err)
		http.Error(w, "Modelo não disponível", http.StatusServiceUnavailable)
		return
	}

	// 5. Chama o serviço gRPC para geração de texto
	stream, err := grpc_services.ClientLLM.GenerateText(r.Context(), &grpc_services.PromptRequest{
		Prompt: payload.Input,
		Model:  &payload.Model,
	})
	if err != nil {
		log.Printf("Erro na chamada gRPC: %v", err)
		http.Error(w, "Falha ao iniciar geração de texto", http.StatusInternalServerError)
		return
	}

	// 6. Configura streaming
	setupStreamingHeaders(w)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming não suportado", http.StatusInternalServerError)
		return
	}

	// 7. Processa respostas
	processStream(stream, w, flusher)
}

// ensureModelReady verifica se o modelo está pronto e o baixa se necessário
func ensureModelReady(ctx context.Context, model string) error {
	// Cria contexto com timeout para o download
	dlCtx, cancel := context.WithTimeout(ctx, modelDownloadTimeout)
	defer cancel()

	for i := 0; i < maxRetries; i++ {
		// Verifica saúde do modelo
		healthResp, err := grpc_services.ClientLLM.HealthCheck(ctx, &grpc_services.HealthRequest{})
		
		if err != nil {
			log.Printf("Erro no health check: %v", err)
			return err
		}

		if healthResp.Ready {
			return nil // Modelo pronto
		}

		// Se não estiver pronto, tenta baixar
		if i == 0 {
			log.Printf("Modelo %s não disponível, iniciando download...", model)
			_, err := grpc_services.ClientLLM.LoadModel(dlCtx, &grpc_services.ModelRequest{
				ModelName: model,
			})
			if err != nil {
				log.Printf("Erro ao iniciar download: %v", err)
				return err
			}
		}

		// Aguarda antes de tentar novamente
		select {
		case <-time.After(retryInterval):
		case <-dlCtx.Done():
			return dlCtx.Err()
		}
	}

	return fmt.Errorf("modelo %s não ficou pronto após %d tentativas", model, maxRetries)
}

// setupStreamingHeaders configura os headers para streaming
func setupStreamingHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

// processStream processa o stream de respostas
func processStream(stream grpc_services.LLMService_GenerateTextClient, w http.ResponseWriter, flusher http.Flusher) {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Erro no stream: %v", err)
			return
		}

		if _, err := w.Write([]byte(resp.Text)); err != nil {
			log.Printf("Erro ao escrever resposta: %v", err)
			return
		}
		flusher.Flush()

		if resp.IsFinal {
			break
		}
	}
}