# Dia 19 - divisible-sum-pairs

## Problema

**Plataforma**: HackerRank
**Dificuldade**: Easy
**Link**: https://www.hackerrank.com/challenges/divisible-sum-pairs

### Enunciado

Dado um array de inteiros `ar` e um inteiro positivo `k`, determine o número de pares `(i, j)` onde `i < j` e a soma `ar[i] + ar[j]` é divisível por `k`.

## Solução

### Abordagem

A abordagem utilizada foi a de força bruta, que é suficiente dadas as constraints do problema (n <= 100).

1.  Inicializei um contador `count` para armazenar o número de pares válidos.
2.  Utilizei um loop aninhado para gerar todos os pares `(i, j)` possíveis com a condição `i < j`.
    - O loop externo itera de `i = 0` até `len(ar) - 1`.
    - O loop interno itera de `j = i + 1` até `len(ar) - 1`.
3.  Para cada par, verifiquei se a soma `ar[i] + ar[j]` é divisível por `k` usando o operador módulo (`%`).
4.  Se `(ar[i] + ar[j]) % k == 0`, o contador `count` é incrementado.
5.  Ao final, a função retorna o valor de `count`.

### Complexidade

- **Tempo**: O(n^2), onde `n` é o número de elementos no array. Isso se deve aos loops aninhados.
- **Espaço**: O(1), pois apenas uma variável de contagem é utilizada, sem alocação de espaço adicional que dependa do tamanho da entrada.

## Implementação

```go
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    var count int32 = 0
    for i := 0; i < len(ar); i++ {
        for j := i + 1; j < len(ar); j++ {
            if (ar[i]+ar[j])%k == 0 {
                count++
            }
        }
    }
    return count
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Este problema foi um bom exercício para reforçar o uso de loops aninhados para encontrar pares em um array.
- A utilização do operador módulo (`%`) é fundamental para verificar a divisibilidade.
- A complexidade de tempo O(n^2) é aceitável para pequenas entradas, mas para problemas com `n` maior, uma abordagem mais otimizada (talvez usando um mapa de frequências dos restos) seria necessária.