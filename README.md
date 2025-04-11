
# Gateway API

Projeto de gateway de APIs desenvolvido com Go e TypeScript, integrando serviços de backend, frontend e um sistema antifraude baseado em NestJS.

## 📦 Estrutura do Projeto

- **backend/**: Serviço principal desenvolvido em Go, responsável por orquestrar as requisições e integrar os demais serviços.
- **frontend/**: Interface web desenvolvida com tecnologias modernas de frontend (possivelmente React ou Angular).
- **nestjs-anti-fraud/**: Módulo antifraude construído com NestJS, focado em segurança e validação de transações.

## 🚀 Tecnologias Utilizadas

- **Go**: Linguagem principal para o serviço de gateway.
- **TypeScript**: Utilizado no frontend e no módulo antifraude.
- **NestJS**: Framework Node.js para o serviço antifraude.
- **CSS/JavaScript**: Complementam o desenvolvimento da interface web.

## 🛠️ Como Executar

Instruções básicas para rodar os serviços (substitua pelos comandos reais conforme sua implementação):

### Backend (Go)

```bash
cd backend
go run main.go
```

### Frontend

```bash
cd frontend
npm install
npm start
```

### Antifraude (NestJS)

```bash
cd nestjs-anti-fraud
npm install
npm run start
```

## 🧪 Testes

Execute os testes para cada serviço individualmente:

### Backend

```bash
cd backend
go test ./...
```

### Frontend

```bash
cd frontend
npm test
```

### Antifraude

```bash
cd nestjs-anti-fraud
npm run test
```

## 📷 Capturas de Tela

![Interface do Projeto](./Captura%20de%20tela%20de%202025-04-10%2013-24-49.png)

## 📄 Licença

Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](./LICENSE) para mais informações.
