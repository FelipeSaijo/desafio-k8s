# Template Config Error

## Soluções

- Adicionar o campo `storageClassName: standard` dentro de spec nos templates pvc e pv.

- No template PVC adicionar o campo `volumeName: nginx-config` dentro de spec.

- Corrigir `claimName` no template deployment.