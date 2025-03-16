# Postgres Error

## Soluções

- Adicionar campo de env com a senha admin do postgres no template `StatefulSet`:
``` yaml
env:
  - name: POSTGRES_PASSWORD
    value: mysecretpassword
```

- Corrigir a porta do postgres no service para `5432`

- Subir um pod postgres no cluster e acessar o postgres com o DNS `postgres.default.svc.cluster.local`
``` bash
psql -U postgres -h postgres.default.svc.cluster.local -W
```

- Criar um port-forward para o service do postgres e criar um container docker na maquina do host e acessar via porta exposta no portforward
```
kubectl port-forward svc/postgres <HOST PORT>:5432

docker exec -it <container id> psql postgres -h localhost -p <HOST PORT> -W
```