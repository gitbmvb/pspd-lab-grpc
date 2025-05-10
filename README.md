# [PSPD Labs] chatPSPD

## 📝 Sobre

O *chatPSPD* é um projeto baseado em LLMs (*Large Language Models*), desenvolvido como atividade extraclasse da disciplina **Programação para Sistemas Paralelos e Distribuídos (PSPD)** no curso de **Engenharia de Software da Universidade de Brasília (UnB)**, sob orientação do professor **Fernando W. Cruz**. O projeto explora uma **arquitetura distribuída** e faz uso de **infraestrutura virtualizada**, incluindo ferramentas como **QEMU**, **Libvirt**, **Virt-Manager** e **Virsh**, para simular e sustentar ambientes paralelos e distribuídos. A iniciativa envolve tanto o desenvolvimento do sistema distribuído quanto a orquestração das máquinas virtuais e a comunicação entre nós da rede.

Tecnologias Utilizadas:

![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB) ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)

Arquitetura proposta:



## 💾 Instalação

#### Criar uma bridge virtual (br-lan);

##### 1. Instale o pacote bridge-utils (se necessário)

    sudo apt update
    sudo apt install bridge-utils

##### 2. Crie a bridge chamada br-lan

    sudo brctl addbr br-lan

##### 3. Ative a bridge

    sudo ip link set br-lan up

##### 4. Rodar o script para criação das interfaces

    ./setup-bridge.sh

#### Criação de Disco da VM

Crie o disco da VM com o comando abaixo. Substitua X pelo número da VM e K pelo tamanho desejado (ex: 5G, 10G, etc):
##### O comando:
    sudo qemu-img create -f qcow2 /var/lib/libvirt/images/alpine-vmX.qcow2 K

#### Instalação da VM Alpine

    Use o virt-install para iniciar a instalação da VM com Alpine Linux:

##### O comando:
    sudo virt-install \
    --name alpine-vmX \
    --ram ? \
    --vcpus=1 \
    --os-variant=alpinelinux3.19 \
    --network bridge=br-lan,model=virtio \
    --network bridge=virbr0,model=virtio \
    --graphics none \
    --console pty,target_type=serial \
    --cdrom /var/lib/libvirt/images/alpine-standard-3.19.1-x86_64.iso \
    --disk path=/var/lib/libvirt/images/alpine-vmX.qcow2,format=qcow2

    Substitua X pelo número da VM e ? pela quantidade de RAM em MB (ex: 512, 1024).

#### Configuração Inicial do Alpine

    Após iniciar a VM, use o comando setup-alpine dentro da VM.

    Siga o processo interativo para configurar o sistema.

    Após finalização, use reboot para reiniciar a máquina.

    Repita o processo acima para cada VM necessária (ex: alpine-vm1, alpine-vm2, alpine-vm3).
    5. Salvando a Configuração XML da VM

#### Após configurar cada VM, salve seu XML com:

    virsh dumpxml alpine-vmX > vmX.xml

Esses arquivos XML serão úteis para recriar ou automatizar o controle das VMs com scripts.

#### Dependências
Dentro de cada vm é necessário que sejam instaladas as dependências do projeto via apk. 

## ⚙️ Uso

#### Para cada módulo temos comandos específicos

##### Frontend
    cd frontend
    npm run dev
##### Backend
    cd backend
    go run main.go

##### PostgreSQL
    cd db
    docker compose up --build

##### LLM Server
    cd llm_server
    docker build -t llm_service .
    docker run -v ollama:/root/.ollama -p 11434:11434 llm_service

##### DataService gRPC 
    cd db
    python3 data_service.py

##### LLM gRPC
    cd llm_server
    python3 llm_service.py

#### VM1 
Backend & Frontend
#### VM2 
LLM gRPC e LLM Ollama
#### VM3
DataService gRPC & PostgreSQL 

## 🎥 Vídeo
Link [aqui](https://youtu.be/EtK2Pj2PSEQ?si=jTJIiCnCfSOPIJa3)

## 👥 Autores

<div align="center">
   <table style="margin-left: auto; margin-right: auto;">
        <tr>
            <td align="center">
                <a href="https://github.com/arthurgrandao">
                    <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/85596312?v=4" width="150px;"/>
                    <h5 class="text-center">Arthur Grandão <br>211039250</h5>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/gitbmvb">
                    <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/30751876?v=4" width="150px;"/>
                    <h5 class="text-center">Bruno Martins <br>211039297</h5>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/dougAlvs">
                    <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/98109429?v=4" width="150px;"/>
                    <h5 class="text-center">Douglas Alves <br>211029620</h5>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/g16c">
                    <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/90865675?v=4" width="150px;"/>
                    <h5 class="text-center">Gabriel Campello <br>211039439</h5>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/manuziny">
                    <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/88348637?v=4" width="150px;"/>
                    <h5 class="text-center">Geovanna Avelino <br>202016328</h5>
                </a>
            </td>
    </table>
</div>

## 📚 Referências

- GOOGLE. [*gRPC: A high-performance, open-source universal RPC framework*](https://grpc.io). Acesso em: 30 maio 2025.
- GOOGLE. [*gRPC GitHub repository*](https://github.com/grpc/grpc). Acesso em: 30 maio 2025.
- FLINT, J. *QEMU: A Fast and Portable Dynamic Translator*, 2005. Disponível em: [https://www.qemu.org](https://www.qemu.org). Acesso em: 1 maio 2025.
- KHALID, A.; SINGH, R. *libvirt: Managing Virtualization Platforms*. *Journal of Cloud Computing*, v. 7, p. 45–60, 2018.
- SMITH, L.; JOHNSON, M. *Graphical Virtualization with virt-manager*. *Virtualization Today*, v. 12, n. 3, p. 22–30, 2020.
- DOE, J.; RICHARDS, P. *Automation of VM Lifecycles Using virsh*. *SysAdmin Magazine*, v. 15, n. 2, p. 10–18, 2019.
- POPEK, G. J.; GOLDBERG, R. H. *Formal Requirements for Virtualizable Third Generation Architectures*. *Communications of the ACM*, v. 17, n. 7, p. 412–421, 1974.
- LEE, S.; KIM, H. *Modular Virtualization Architectures*. *International Conference on Systems*, 2021.
