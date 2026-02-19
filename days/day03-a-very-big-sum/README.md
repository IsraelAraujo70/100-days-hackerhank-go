# #100DiasDeHackerRank — Dia 3/100 (Go)

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/a-very-big-sum/

### Enunciado

A Very Big Sum - Calcular e imprimir a soma dos elementos de um array, considerando que alguns inteiros podem ser muito grandes (exceder o range de int32).

**Função**: `aVeryBigSum(ar []int64) int64`
- **Parâmetro**: array de inteiros int64
- **Retorno**: soma de todos os elementos como int64

**Constraints**: 1 ≤ n ≤ 10, 0 ≤ ar[i] ≤ 10^10

## Solução

### Abordagem

• Usar int64 para evitar overflow de int32
• Iterar pelo array somando cada elemento
• Usar `for range` omitindo o índice (`for _, v := range ar`)
• Retornar a soma total

### Complexidade

- **Tempo**: O(n) onde n é o tamanho do array
- **Espaço**: O(1) - apenas variável para soma

## Implementação

A solução é direta: iterar pelos elementos e somar, usando int64 para suportar números grandes.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Importância de usar tipos adequados para evitar overflow
- int32 range: -2^31 a 2^31-1 (±2.1 bilhões)
- int64 suporta números muito maiores
- Continuidade do padrão `for _, v := range` aprendido no Dia 1

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
