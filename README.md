# API de Questões em Go

Uma API REST desenvolvida em Go para gerenciamento de questões, integrada com Supabase.

## 🚀 Funcionalidades

- Buscar questões com filtros
- Adicionar novas questões
- Integração com banco de dados Supabase

## 📋 Pré-requisitos

- Go 1.24 ou superior
- Conta no Supabase

## ⚙️ Configuração

1. Clone o repositório:
```bash
git clone <url-do-seu-repositorio>
cd go_api
```

2. Copie o arquivo de exemplo das variáveis de ambiente:
```bash
cp .env.example .env
```

3. Configure suas credenciais do Supabase no arquivo `.env`:
```
SUPABASE_KEY=sua_chave_supabase_aqui
SUPABASE_URL=https://seu_projeto.supabase.co
```

4. Instale as dependências:
```bash
go mod download
```

5. Execute a aplicação:
```bash
go run main/main.go
```

## 📡 Endpoints

### GET/POST `/questoes`
Busca questões com filtros opcionais

### POST `/add-questao`
Adiciona uma nova questão

## 🛠️ Tecnologias Utilizadas

- Go
- Supabase
- HTTP Router nativo do Go

## 📝 Estrutura do Projeto

```
├── config/          # Configurações da aplicação
├── controllers/     # Controllers/endpoints HTTP
├── main/           # Arquivo principal
├── services/       # Lógica de negócio
├── structs/        # Estruturas de dados
├── .env.example    # Exemplo de variáveis de ambiente
└── .gitignore      # Arquivos ignorados pelo Git
```
