# Dia 23 - Drawing Book

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/drawing-book/problem

### Enunciado

Dado um livro com `n` páginas e uma página alvo `p`, calcule o menor número de viradas de página necessárias para chegar em `p`, podendo começar a virar a partir da frente ou de trás do livro. Cada virada avança duas páginas (exceto possível última página ímpar).

## Solução

### Abordagem

Usa simetria das viradas: o número de viradas a partir da frente é `p/2` e a partir do fim é `n/2 - p/2` (divisão inteira). A resposta é `min(front, back)`. Isso já lida corretamente com `n` par/ímpar e com a última página isolada.

### Complexidade

- **Tempo**: O(1)
- **Espaço**: O(1)

## Implementação

- Função: `pageCount(n, p) int32` em `days/day23-drawing-book/main.go:21`
- Lógica: `fromFront := p/2` e `fromBack := n/2 - p/2`; retorna o menor.

## Resultado

- [x] Problema resolvido
- [ ] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Divisão inteira em Go com `int32` funciona naturalmente para o cálculo (`/` trunca para baixo).
- A modelagem por pares de páginas torna o problema trivial com `p/2` e `n/2`.
- Importância de considerar a última página quando `n` é ímpar — já coberto pela fórmula.

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
