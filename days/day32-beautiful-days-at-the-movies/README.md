# Dia 32 - Beautiful Days at the Movies

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem

### Enunciado

Dados dois inteiros `i` e `j` que definem um intervalo fechado de dias e um divisor `k`, conte quantos dias possuem uma diferença absoluta entre o número original e o número invertido que seja divisível por `k`.

## Solução

### Abordagem

- Iterar de `i` até `j`, invertendo os dígitos do dia atual.
- Calcular a diferença absoluta entre o valor original e o invertido.
- Incrementar o contador quando a diferença for divisível por `k`.

### Complexidade

- **Tempo**: O(j - i + 1)
- **Espaço**: O(1)

## Implementação

- Código: `days/day32-beautiful-days-at-the-movies/main.go`

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforcei a manipulação de dígitos diretamente com operações aritméticas, evitando conversões para string.

## Screenshots

Adicione aqui o print do "Accepted" quando disponível.
