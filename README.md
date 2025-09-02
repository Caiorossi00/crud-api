# API CRUD Multi-Linguagem

Este repositório serve como um portfólio de aprendizado prático que demonstra a implementação de **APIs CRUD** (Criar, Ler, Atualizar, Deletar) em diversas linguagens e frameworks de backend.

Cada pasta contém uma versão completa e funcional da API, permitindo a comparação de abordagens e a prática de conceitos fundamentais como **roteamento**, **manipulação de dados** e **respostas HTTP**.

## Estrutura do Repositório

O repositório é organizado em pastas, com cada uma representando uma implementação da API. Por exemplo:

`/python-flask`: Implementação da API usando Python e o micro-framework Flask.  
`/node-express`: Implementação da API usando Node.js e o framework Express.

<br>

## Conteúdo

### API em Python + Flask (`/python-flask`)

Esta API gerencia uma coleção de **receitas culinárias** e implementa as quatro operações básicas de um CRUD:

#### **Endpoints**

| Método | Rota                        | Descrição                               | Exemplo de Retorno |
|--------|-----------------------------|----------------------------------------|---------------------|
| **GET**    | `/receitas`                  | Retorna todas as receitas cadastradas | Lista de receitas |
| **GET**    | `/receitas/<id>`             | Retorna uma receita específica pelo **ID** | Receita individual |
| **POST**   | `/receitas`                  | Adiciona uma nova receita à lista | Nova receita criada |
| **PUT**    | `/receitas/<id>`             | Atualiza uma receita existente | Receita atualizada |
| **DELETE** | `/receitas/<id>`             | Remove uma receita pelo **ID** | Mensagem de sucesso |

---

### API em Node + Express (`/node-express`)

Esta API gerencia uma coleção de **clientes** e implementa as quatro operações básicas de um CRUD:

#### **Endpoints**

| Método | Rota                        | Descrição                               | Exemplo de Retorno |
|--------|-----------------------------|----------------------------------------|---------------------|
| **GET**    | `/clientes`                  | Retorna todos os clientes cadastrados | Lista de clientes |
| **GET**    | `/clientes/:id`             | Retorna um cliente específico pelo **ID** | Cliente individual |
| **POST**   | `/clientes`                  | Adiciona um novo cliente à lista | Novo cliente criado |
| **PUT**    | `/clientes/:id`             | Atualiza um cliente existente | Cliente atualizado |
| **DELETE** | `/clientes/:id`             | Remove um cliente pelo **ID** | Mensagem de sucesso |

---


