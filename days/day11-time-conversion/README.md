# Dia 11 - time-conversion

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/time-conversion

### Enunciado

Dado um tempo no formato de 12 horas AM/PM, converta para o formato militar (24 horas).

**Observações:**
- 12:00:00AM em relógio de 12 horas é 00:00:00 em relógio de 24 horas
- 12:00:00PM em relógio de 12 horas é 12:00:00 em relógio de 24 horas

## Solução

### Abordagem

Utilizei as funções `time.Parse` e `time.Format` do Go para fazer a conversão:
- Parse com layout "03:04:05PM" para interpretar o formato 12h
- Format com layout "15:04:05" para gerar o formato 24h

### Complexidade

- **Tempo**: O(1)
- **Espaço**: O(1)

## Implementação



## Resultado

- [x] Problema resolvido
- [ ] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Aprendi sobre layouts de time em Go (03:04:05PM e 15:04:05)
- Go usa um tempo de referência específico: "Mon Jan 2 15:04:05 MST 2006"
- As funções time.Parse e time.Format são muito poderosas para conversões

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
