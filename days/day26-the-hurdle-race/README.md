# Dia 26 - The Hurdle Race

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/the-hurdle-race/problem

### Enunciado

Dados a altura máxima de salto natural `k` e uma lista de alturas de obstáculos `height`, determine quantas doses de poção (cada dose aumenta o salto em 1) são necessárias para que o atleta consiga pular todos os obstáculos. Se `k` já for suficiente, retorne `0`.

## Solução

### Abordagem

Encontrar a maior altura em `height` e calcular `max(0, maxHeight - k)`. Se `maxHeight <= k`, não são necessárias doses; caso contrário, a diferença indica o mínimo de doses.

### Complexidade

- **Tempo**: O(n)
- **Espaço**: O(1)
- Código: `days/day26-the-hurdle-race/main.go`
## Implementação



## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforço do padrão de "clamp em zero" ao subtrair capacidades de requerimentos: `max(0, necessidade - capacidade)`.
- Implementação direta e eficiente com passagem única para achar o máximo.

## Screenshots

- Adicionar print do resultado "Accepted" quando disponível.
