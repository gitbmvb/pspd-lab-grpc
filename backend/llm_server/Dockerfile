FROM ollama/ollama:latest


COPY scripts/start_llm_server.sh scripts/start_llm_server.sh

WORKDIR /scripts

RUN chmod +x start_llm_server.sh

ENTRYPOINT ["/usr/bin/bash", "start_llm_server.sh"]
