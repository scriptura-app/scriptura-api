# Definir la variable BINARIO con el resultado del comando 'command -v'
SWAG_BIN := $(shell command -v swag || echo "")

# Comprobar si el binario 'swag' está instalado.
check_swag:
ifndef SWAG_BIN
	$(error "El binario 'swag' no está instalado. Por favor, instálalo para continuar. https://github.com/swaggo/swag")
else
	@echo "El binario 'swag' está instalado, ejecutando 'swag init'..."
	swag init
endif

# Compilar el proyecto con 'go build', precedido por 'swag init'
compile: check_swag
	@echo "Compilando el proyecto..."
	go build

# Ejecutar el proyecto con 'go run main.go', precedido por 'swag init'
run: check_swag
	@echo "Ejecutando la aplicación Go..."
	go run main.go
