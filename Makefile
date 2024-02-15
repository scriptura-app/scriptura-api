# Definir la variable BINARIO con el resultado del comando 'command -v'
SWAG_BIN := $(shell command -v swag || echo "")
GQLGEN_BIN := $(shell command -v gqlgen || echo "")

# Comprobar si el binario 'swag' está instalado.
check_swag:
ifndef SWAG_BIN
	$(error "El binario 'swag' no está instalado. Por favor, instálalo para continuar. https://github.com/swaggo/swag")
else
	@echo "El binario 'swag' está instalado, ejecutando 'swag init'..."
	swag init
endif

# Comprobar si el binario 'swag' está instalado.
check_gql:
ifndef GQLGEN_BIN
	$(error "El binario 'gqlgen' no está instalado. Por favor, instálalo para continuar. go install github.com/99designs/gqlgen@latest")
else
	@echo "El binario 'gqlgen' está instalado, ejecutando 'gqlgen generate'..."
	gqlgen generate --config graphql/gqlgen.yml
endif


# Compilar el proyecto con 'go build', precedido por 'swag init'
compile: check_swag
	check_gql
	@echo "Compilando el proyecto..."
	go build

# Ejecutar el proyecto con 'go run main.go', precedido por 'swag init'
run: check_swag
	check_gql
	@echo "Ejecutando la aplicación Go..."
	go run main.go
