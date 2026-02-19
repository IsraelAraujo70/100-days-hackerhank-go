# Dia 14 - Between Two Sets

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy
**Link**: [https://www.hackerrank.com/challenges/between-two-sets/problem](https://www.hackerrank.com/challenges/between-two-sets/problem)

### Enunciado

Dado dois arrays de inteiros, `a` e `b`, determine quantos números existem que satisfazem as seguintes condições:
1. Todos os elementos em `a` são fatores do número.
2. O número é um fator de todos os elementos em `b`.

Esses números são considerados "entre" os dois conjuntos.

## Solução

### Abordagem

A solução itera através de um intervalo de números e verifica se cada número atende às duas condições do problema.

1.  **Definir o Intervalo de Busca**: O número `x` que estamos procurando deve ser um múltiplo de todos os elementos em `a` e um fator de todos os elementos em `b`. Portanto, `x` deve ser maior ou igual ao maior elemento em `a` (`max(a)`) e menor ou igual ao menor elemento em `b` (`min(b)`). O intervalo de busca é `[max(a), min(b)]`.

2.  **Iterar e Verificar**: O código percorre cada inteiro `x` neste intervalo. Para cada `x`, ele verifica:
    - **Condição 1**: Se `x` é um múltiplo de todos os elementos em `a` (ou seja, `x % valA == 0` para todo `valA` em `a`).
    - **Condição 2**: Se `x` é um fator de todos os elementos em `b` (ou seja, `valB % x == 0` para todo `valB` em `b`).

3.  **Contagem**: Se ambas as condições forem verdadeiras, um contador é incrementado.

Essa abordagem de força bruta é eficiente o suficiente devido às baixas restrições do problema (valores de array até 100).

### Complexidade

- **Tempo**: `O((max(b) - min(a)) * (len(a) + len(b)))`. Como os valores são pequenos, a complexidade é baixa.
- **Espaço**: `O(1)`, pois usamos apenas algumas variáveis para contagem e iteração, sem alocar nova estrutura de dados proporcional à entrada.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit

## Aprendizados

- **Fatores e Múltiplos**: O problema foi um bom exercício sobre a definição de fatores e múltiplos.
- **Delimitação de Intervalo**: A chave para uma solução eficiente foi identificar o intervalo correto para a busca (`[max(a), min(b)]`), o que evita verificações desnecessárias.
- **Força Bruta vs. Otimização**: Embora fosse possível resolver usando LCM (Mínimo Múltiplo Comum) de `a` e GCD (Máximo Divisor Comum) de `b`, a abordagem de força bruta dentro do intervalo correto é mais simples de implementar e perfeitamente aceitável para as restrições dadas.

## Screenshots

*Ainda não adicionei o screenshot.*
