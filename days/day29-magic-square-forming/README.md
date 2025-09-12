# Dia 29 - Forming a Magic Square

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟡 Medium  
**Link**: https://www.hackerrank.com/challenges/magic-square-forming

### Enunciado

Dado um quadrado 3x3 com inteiros no intervalo [1..9], transforme-o em um quadrado mágico (todos os números distintos de 1 a 9 e somas iguais nas linhas/colunas/diagonais) com custo mínimo, onde o custo de alterar uma célula é a diferença absoluta entre o valor original e o novo.

## Solução

### Abordagem

Para 3x3, existem apenas 8 quadrados mágicos possíveis (todas as rotações/reflexões do quadrado de Lo Shu). A estratégia é:

- Pré-computar as 8 matrizes mágicas 3x3 válidas.
- Calcular o custo de converter a matriz de entrada em cada uma delas (soma das diferenças absolutas posição a posição).
- Retornar o menor custo entre os 8.

### Complexidade

- Tempo: O(1) — sempre 8 matrizes de tamanho fixo 3x3.
- Espaço: O(1) — apenas constantes auxiliares.

## Implementação

Ver `days/day29-magic-square-forming/main.go`.

## Resultado

- [x] Problema resolvido

## Aprendizados

- Para 3x3, o conjunto de quadrados mágicos é finito e pequeno; brute-force sobre os 8 casos é a solução mais simples e ideal.

## Screenshots

- (Adicionar quando submetido no HackerRank)
