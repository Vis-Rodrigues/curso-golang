# curso-golang
Projeto com exemplos do curso de golang

## Comandos úteis em Go

Aqui estão alguns comandos essenciais para trabalhar com Go no seu projeto:

### **Gerenciamento do projeto**
- Inicializar um novo módulo Go:
  ```sh
  go mod init nome-do-projeto
  ```
- Adicionar uma dependência ao projeto:
  ```sh
  go get github.com/exemplo/pacote
  ```
- Atualizar todas as dependências:
  ```sh
  go get -u ./...
  ```
- Verificar e limpar dependências não utilizadas:
  ```sh
  go mod tidy
  ```

### **Compilação e execução**
- Executar um arquivo Go:
  ```sh
  go run main.go
  ```
- Compilar um binário:
  ```sh
  go build -o meu-programa main.go
  ```
  
### **Testes**
- Rodar todos os testes do projeto:
  ```sh
  go test ./...
  ```
- Rodar testes com detalhes:
  ```sh
  go test -v ./...
  ```
- Rodar um teste específico:
  ```sh
  go test -run TestNomeDoTeste
  ```

### **Formatação e linting**
- Formatar código Go:
  ```sh
  go fmt ./...
  ```
- Verificar erros estáticos no código:
  ```sh
  go vet ./...
  ```

### **Outros comandos úteis**
- Ver versão do Go instalada:
  ```sh
  go version
  ```
- Ver informações do ambiente Go:
  ```sh
  go env
  ```
