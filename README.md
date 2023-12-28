# go-lang-full-cycle-sqlc
Repositório para armazenar os códigos de exemplo do módulo de SQLC do curso Go-Expert da Full Cycle
- A pasta cmd/runSQLCTX contém exemplo de código para execução de queries com SQLC com transações.
- A pasta cmd/runSQLC contém exemplo de código para execução de queries com SQLC sem transações.

### Dependências
- [Go-Lang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html)
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
### Gerando código SQLC
```bash
 sqlc generate
```