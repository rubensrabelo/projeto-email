## Ordem de Inicialização

Suba os serviços na ordem correta para garantir que tudo inicialize sem problemas.

### 1. Subir o banco de dados

```bash
docker compose --env-file .env -f database/docker-compose.yml up -d
```

### 2. Subir o Keycloak

```bash
docker compose --env-file .env -f keycloak/docker-compose.yml up -d
```

---

## Parar os containers

Para derrubar os containers, siga a ordem inversa:

### 1. Parar o Keycloak

```bash
docker compose --env-file .env -f keycloak/docker-compose.yml down -v
```

### 2. Parar o banco de dados

```bash
docker compose --env-file .env -f database/docker-compose.yml down -v
```
