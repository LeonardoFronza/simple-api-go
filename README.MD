# README

Este guia explica como a aplicação funciona e como utilizá-la com Docker e Kubernetes.

## Funcionamento da Aplicação

A aplicação é desenvolvida em GO e possui as seguintes funcionalidades:

1. Swagger - `get_random_list`
2. Hello, World!

Essa aplicação foi criada como um desafio para aprender a utilizar o Kubernetes e explorar a linguagem GO.

## Utilizando com Docker

Para utilizar a aplicação com Docker, siga os passos abaixo:

1. Certifique-se de ter o Docker instalado em sua máquina.
2. Clone o repositório do projeto.
3. Navegue até o diretório raiz do projeto.
4. Execute o comando `docker compose build` para construir a imagem Docker.
5. Execute o comando `docker compose up -d` para iniciar o contêiner Docker.
6. O projeto estará disponível na porta `http://localhost:8080`.

## Utilizando com Kubernetes

Para utilizar a aplicação com Kubernetes, siga os passos abaixo:

1. Certifique-se de ter o Kubernetes instalado em seu ambiente.
2. Clone o repositório do projeto.
3. Navegue até o diretório raiz do projeto.
4. Execute o comando `kubectl apply -f kubernetes-deploy.yaml` para criar os recursos no Kubernetes.
5. Execute o comando `minikube tunnel` para criar um túnel entre as portas.
6. Verifique o status dos pods com o comando `kubectl get pods`.
7. Acesse a aplicação através do serviço exposto pelo Kubernetes.

Link Docker Hub: https://hub.docker.com/repository/docker/leonardofronza/api-go-gurizes-app/general