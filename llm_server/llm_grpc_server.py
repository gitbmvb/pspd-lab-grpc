from concurrent import futures
import json
import requests

import grpc

from grpc_services import service_pb2
from grpc_services import service_pb2_grpc

class LLMServiceServicer(service_pb2_grpc.LLMServiceServicer):
    def __init__(self):
        self.ollama_url = "http://localhost:11434/api/generate"

    def GenerateText(self, request, context):
        payload = {
            "model": request.model or "llama3.2",
            "prompt": request.prompt,
            "stream": True,
            "options": {
                "temperature": request.temperature or 0.8,
                "num_predict": request.max_tokens or 256
            }
        }

        try:
            with requests.post(self.ollama_url, json=payload, stream=True, timeout=30) as resp:
                resp.raise_for_status()
                for line in resp.iter_lines():
                    if line:
                        chunk = json.loads(line)
                        yield service_pb2.TextResponse(
                            text=chunk.get("response", ""),
                            is_final=chunk.get("done", False)
                        )
        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Ollama error: {str(e)}")
            raise

    def HealthCheck(self, request, context):
        try:
            resp = requests.get("http://localhost:11434")
            return service_pb2.HealthResponse(
                ready=resp.ok,
                model="llama3.2"
            )
        except Exception:
            return service_pb2.HealthResponse(ready=False)

    def LoadModel(self, request, context):
        try:
            resp = requests.post("http://localhost:11434/api/pull", json={"model": request.model_name})
            resp.raise_for_status()
            return service_pb2.HealthResponse(
                ready=resp.ok,
                model=request.model
            )
        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Load model error: {str(e)}")
            raise


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_LLMServiceServicer_to_server(
        LLMServiceServicer(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    print("LLM Server started on port 50052")
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
