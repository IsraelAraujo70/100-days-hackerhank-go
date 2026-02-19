# Dia 39 - circular-array-rotation

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: [Circular Array Rotation](https://www.hackerrank.com/challenges/circular-array-rotation)

### Enunciado

Dado um array de inteiros, realizar k rotações circulares para a direita e responder q consultas sobre valores em posições específicas. Uma rotação move o último elemento para a primeira posição.

## Solução

### Abordagem

Em vez de simular as rotações, usei aritmética modular para calcular diretamente onde cada elemento estaria após k rotações. Para encontrar o valor na posição `query` após k rotações, busco o elemento que estava originalmente na posição `(query - k + n) % n`.

### Complexidade

- **Tempo**: O(q) - onde q é o número de consultas
- **Espaço**: O(1) - apenas variáveis auxiliares

## Implementação

```go
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
    n := int32(len(a))
    k = k % n  // Otimização: k rotações = k%n rotações
    
    result := make([]int32, len(queries))
    
    for i, query := range queries {
        originalIndex := (query - k + n) % n
        result[i] = a[originalIndex]
    }
    
    return result
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Aritmética modular para rotações: `(query - k + n) % n`
- Otimização k % n evita rotações desnecessárias
- Solução O(1) por consulta vs O(k) da simulação
- Fórmula inversa: se elemento vai de i para (i+k)%n, então posição j veio de (j-k+n)%n

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
