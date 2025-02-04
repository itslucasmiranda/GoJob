.PHONY: default run build test docs clean #declara as tasks que não são arquivos, ou seja, ações e não arquivos que o makefile deve procurar
# Variables
APP_NAME=gojopportunities

# Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
#gera os arquivos de documentação com o swag que é utilizado para gerar a documentação swagger de APIs Go
	@swag init 
#executa o app após gerar a documentação
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
#executa todos os testes da aplicação, buscando em todos os packages e subpackages
test:
	@go test ./ ...
docs:
#gera a documentação e cria ou atualiza os arquivos de doc necessários
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs