# go-lang-full-cycle-sqlc
Repositório para armazenar os códigos de exemplo do módulo de SQLC do curso Go-Expert da Full Cycle

# Go-Lang Migrations
### Dependências
- [Go-Lang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Gerando migrations
```bash
 make create-migration name=<nome-da-migration>
```
### Rodando migrations
```bash
 make migrate
```
### Rollback migrations
```bash
 make rollback
```