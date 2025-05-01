FROM ollama/ollama:latest


COPY llm_server/scripts/start_llm_server.sh scripts/start_llm_server.sh

WORKDIR /scripts

RUN chmod +x start_llm_server.sh
EXPOSE 11434
ENTRYPOINT ["/usr/bin/bash", "start_llm_server.sh"]
