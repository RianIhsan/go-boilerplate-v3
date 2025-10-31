# ==============================================================================
# Golang commands
general:
	go run ./cmd/api/main.go

ws:
	go run ./cmd/realtime/main.go

worker:
	go run ./cmd/worker/main.go

build-api:
	go build -o ./api ./cmd/api/main.go

build-worker:
	go build -o ./worker ./cmd/worker/main.go

build-realtime:
	go build -o ./ws ./cmd/realtime/main.go

build: build-api build-worker build-realtime
	@echo "✅ All services built"

# Supervisor comma-backend

restart-worker:
	supervisorctl restart dms-worker

restart-realtime:
	supervisorctl restart dms-realtime

restart-all: restart-api restart-worker restart-realtime
	@echo "♻️ All services restarted"

test:
	go test

user-repository-test:
	go test -v ./internal/user/repository

user-service-test:
	go test -v ./internal/user/service

user-controller-test:
	go test -v ./internal/user/controllers/http
#===============================================================================

# ==============================================================================
# Docker compose commands
local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yaml up --build -d

develop:
	echo "Starting docker environment"
	docker compose -f docker-compose.dev.yaml up --build -d
#===============================================================================

# ==============================================================================
# SSL/TLS commands

#generate private key self-signed certificate (public key)
gen_private_key:
	openssl genrsa -out server.key 2048
	openssl ecparam -genkey -name secp384r1 -out server.key

#generate self-signed certificate (public key)
gen_self_signed_cert:
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
#===============================================================================

# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES) 	#perintah Docker untuk menghentikan semua container yang id nya ada di dalam variable files
	docker rm $(FILES) 		#perintah ini akan menghapus semua container yang id nya ada di dalam variabel files

# membersihkan resource Docker yang tidak digunakan seperti container berhenti,
# image lama, volume, dan network yang tidak terpakai
clean:
	docker system prune -f

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	#go mod vendor


dcbuild:
	docker build -t dms-evolution -f docker/Dockerfile .


MOCKGEN=mockgen
FEATURE?=

mock:
	@echo "Please provide feature name. Example: make mock access"

mock-%:
	$(MOCKGEN) \
		-source=internal/features/$*/repository_interface.go \
		-destination=internal/features/$*/mocks/repository_mock_interface.go \
		-package=mocks
	@echo "✅ Mock generated for feature: $*"
