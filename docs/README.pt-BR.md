# 📧 Sketch Service Email

O **sketch-service-email** é um microserviço desenvolvido em Go, focado no processamento e envio de e-mails de forma assíncrona.

> **Projeto de Estudos:** Este repositório foi criado exclusivamente para fins de aprendizado. O objetivo principal é explorar a integração de frameworks modernos, gerenciamento de dependências e padrões de concorrência em Go.

---

## 🛠️ Tecnologias e Conceitos Explorados

O projeto utiliza uma stack robusta para simular um ambiente de produção real:

* **Go (1.25.5):** Linguagem base do projeto.
* **Uber fx:** Utilizado para **Injeção de Dependência**, facilitando a modularização e o ciclo de vida da aplicação.
* **Fiber v3:** Framework web de alta performance para a criação das rotas de API.
* **Gomail:** Biblioteca para a construção e envio de e-mails via SMTP.
* **Concorrência (Worker Pool):** Implementação de filas com `channels` para processamento de e-mails em background, evitando o bloqueio da requisição HTTP.

---

## 🏗️ Arquitetura do Sistema

O projeto segue uma estrutura organizada por módulos:

* **`config/`**: Gerenciamento de variáveis de ambiente usando [dotenv](https://github.com/rickferrdev/dotenv).
* **`internal/handlers/`**: Camada de transporte (HTTP) que recebe as requisições de assinatura.
* **`internal/services/`**: Regras de negócio, onde a lógica de assinatura e preparação do e-mail reside.
* **`pkg/mail/`**: Pacote utilitário responsável pela comunicação direta com o servidor SMTP e gerenciamento do Worker.

---

## 🚀 Como Executar

### Pré-requisitos
* Go 1.25.5 ou superior instalado.
* Um servidor SMTP para testes (ex: Mailtrap).

### Passo a Passo
1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/rickferrdev/sketch-service-email.git
    ```

2.  **Configure as variáveis de ambiente:**
    Crie um arquivo `.env` na raiz com base nas chaves definidas no código:
    ```env
    MAIL_PORT=587
    MAIL_HOST=smtp.exemplo.com
    MAIL_USERNAME=seu_usuario
    MAIL_PASSWORD=sua_senha
    ```

3.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

4.  **Inicie a aplicação:**
    ```bash
    go run main.go
    ```
    O servidor iniciará na porta **8080**.

---

## 📨 Endpoint Principal

### Criar Assinatura
`POST /api/v1/subs/`

**Request Body:**
```json
{
  "email": "usuario@exemplo.com"
}
```

**O que acontece internamente?**
1. O Handler valida o corpo da requisição.
2. O Service gera o conteúdo do e-mail.
3. O Mailer adiciona a mensagem a um **channel**.
4. Um **Worker** em segundo plano consome esse canal e realiza o envio via SMTP.

---

## 📝 Licença

Este projeto está sob a licença **MIT**.

---
