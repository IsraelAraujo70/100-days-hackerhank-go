# Dia 42 - sherlock-and-squares

## Problema

**Plataforma**: [HackerRank]
**Dificuldade**: [Easy]
**Link**: [https://www.hackerrank.com/challenges/sherlock-and-squares]

### Enunciado

Watson likes to challenge Sherlock's math ability. He will provide a starting and ending value that describe a range of integers, inclusive of the endpoints. Sherlock must determine the number of square integers within that range.

Note: A square integer is an integer which is the square of an integer, e.g. 1, 4, 9, 16, 25.

Example: Range 3 to 9 contains squares 4 and 9, so the answer is 2.

## Solução

### Abordagem

A solução utiliza uma abordagem matemática eficiente:

1. Encontrar o menor inteiro cujo quadrado seja >= a (teto da raiz quadrada de a)
2. Encontrar o maior inteiro cujo quadrado seja <= b (piso da raiz quadrada de b)
3. A diferença entre esses valores + 1 dá a quantidade de quadrados perfeitos no intervalo

Esta abordagem é O(1) pois envolve apenas operações matemáticas constantes.

### Complexidade

- **Tempo**: O(1) - operações matemáticas constantes
- **Espaço**: O(1) - apenas variáveis temporárias

## Implementação

```go
func squares(a int32, b int32) int32 {
    // Find the smallest integer whose square is >= a
    start := int32(math.Ceil(math.Sqrt(float64(a))))

    // Find the largest integer whose square is <= b
    end := int32(math.Floor(math.Sqrt(float64(b))))

    // If start > end, no squares in range
    if start > end {
        return 0
    }

    return end - start + 1
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- [Uso de math.Ceil e math.Floor para encontrar limites de raiz quadrada]
- [Abordagem matemática O(1) para contar quadrados perfeitos em um intervalo]
- [Conversão cuidadosa entre int32 e float64 para cálculos de raiz quadrada]
- [Importância de verificar se start <= end para evitar resultados negativos]

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
