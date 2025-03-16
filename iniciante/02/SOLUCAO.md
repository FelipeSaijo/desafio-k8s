# Template Config Error

## Soluções

- Remover `nodeSelector` para o pod ser criado em qualquer node

- Corrigir os valores dos `requests`, eles não podem ser maiores que os `limits`

- Corrigir `selector` no service para identificar o pod.