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

### Uso de GPU

Caso você tenha uma gpu da NVIDIA, o ollama pode usá-la para acelerar o processo de inferência. Para isso, basta seguir o passo a passo a seguir para instalar o `NVIDIA Container Toolkit`⁠ e instruir o Docker para usá-lo.


#### Instalação em Distros com Apt

1. Configure o repositório:
```
curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey \
    | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg
curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list \
    | sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' \
    | sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
sudo apt-get update
```

2. Instale o NVIDIA Container Toolkit:
```
sudo apt-get install -y nvidia-container-toolkit
```

#### Instalação em Distros com Yum ou Dnf

1. Configure o repositório:
```
curl -s -L https://nvidia.github.io/libnvidia-container/stable/rpm/nvidia-container-toolkit.repo \
    | sudo tee /etc/yum.repos.d/nvidia-container-toolkit.repo
```

2. Instale o NVIDIA Container Toolkit:
```
sudo yum install -y nvidia-container-toolkit
```

#### Configuração do Docker

1. Configure o Docker para usar o driver da Nvidia
```
sudo nvidia-ctk runtime configure --runtime=docker
sudo systemctl restart docker
```

2. Por fim, no docker-compose.yml, descomente a linha
```
# runtime: nvidia
```

OBS.: Mais informações podem ser obtidas na [página da imagem](https://hub.docker.com/r/ollama/ollama) do ollama no dockerhub.