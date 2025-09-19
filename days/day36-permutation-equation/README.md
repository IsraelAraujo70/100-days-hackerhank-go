# Dia 36 - permutation-equation

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: https://www.hackerrank.com/challenges/permutation-equation/problem

### Enunciado

Given a permutation p of integers from 1 to n, for each x from 1 to n, find y such that p[p[y] - 1] = x.

## Solução

### Abordagem

Create a position array pos where pos[val] = index of val in p (0-based). Then for each x, y = pos[pos[x] + 1] + 1 (convert to 1-based).

### Complexidade

- **Tempo**: O(n)
- **Espaço**: O(n)

## Implementação

```go
func permutationEquation(p []int32) []int32 {
    n := len(p)
    pos := make([]int32, n+1)
    for i, v := range p {
        pos[v] = int32(i)
    }
    result := make([]int32, n)
    for x := int32(1); x <= int32(n); x++ {
        px := pos[x]
        y := pos[px+1]
        result[x-1] = y + 1
    }
    return result
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Trabalhando com permutações e mapeamento de índices.
- Uso de array para posição ao invés de map para eficiência.
- Conversão entre índices 0-based e 1-based.

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
