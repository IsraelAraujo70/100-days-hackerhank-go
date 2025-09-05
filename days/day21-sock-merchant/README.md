# Dia 21 - Sock Merchant (Sales by Match)

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: https://www.hackerrank.com/challenges/sock-merchant/problem

### Enunciado

Dada uma lista de inteiros representando as cores das meias, determine quantos pares com a mesma cor podem ser formados.

## Solução

### Abordagem

- Usei um mapa de frequências para contar quantas vezes cada cor aparece.
- Para cada cor, o número de pares é `freq // 2` (divisão inteira).
- Somei os pares de todas as cores para obter o total.

### Complexidade

- **Tempo**: O(n), onde `n` é o número de meias.
- **Espaço**: O(k), onde `k` é o número de cores distintas (no pior caso, O(n)).

## Implementação

Veja `days/day21-sock-merchant/main.go` na função `sockMerchant`.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit

## Aprendizados

- Reforcei o padrão de contar frequências com `map` e extrair pares via divisão inteira.

## Screenshots

- [Adicionar print do "Accepted" quando disponível]
