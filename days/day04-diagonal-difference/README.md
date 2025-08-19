# #100DiasDeHackerRank — Dia 4/100 (Go)

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/diagonal-difference/

### Enunciado

Diagonal Difference - Dada uma matriz quadrada, calcular a diferença absoluta entre as somas de suas diagonais.

**Exemplo**:
```
1 2 3
4 5 6
9 8 9
```
- Diagonal principal (esquerda-direita): 1 + 5 + 9 = 15
- Diagonal secundária (direita-esquerda): 3 + 5 + 9 = 17
- Diferença absoluta: |15 - 17| = 2

**Função**: `diagonalDifference(arr [][]int32) int32`
- **Parâmetro**: matriz 2D de inteiros
- **Retorno**: diferença absoluta entre as somas das diagonais

## Solução

### Abordagem

• Iterar pela matriz uma única vez usando o índice i
• Diagonal principal: `arr[i][i]` (linha i, coluna i)
• Diagonal secundária: `arr[i][len-i]` (linha i, coluna len-i)
• Calcular diferença absoluta manualmente (sem usar abs())

### Complexidade

- **Tempo**: O(n) onde n é o tamanho da matriz
- **Espaço**: O(1) - apenas variáveis para as somas

## Implementação

Solução otimizada em uma única passada pela matriz, calculando ambas as diagonais simultaneamente.

## Resultado

- [x] Problema resolvido
- [ ] Accepted no primeiro submit
- [x] Precisou de múltiplas tentativas

## Aprendizados

- Interpretação de entrada: o primeiro número é apenas o tamanho da matriz, não um indicador de formato especial
- Matrizes 2D em Go: `[][]int32` 
- Acesso a diagonais: principal `[i][i]`, secundária `[i][len-i]`
- Valor absoluto manual: `if dif <= 0 { return dif * -1 }`
- **Dificuldade**: Me confundi com o formato da entrada, achei que seria `[[3][1,2,3][1,3,2][3,2,1]]` mas na verdade era apenas o tamanho seguido das linhas da matriz

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
