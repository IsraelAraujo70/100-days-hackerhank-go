# Dia 16 - Apple and Orange

## Problema

**Plataforma**: [HackerRank](https://www.hackerrank.com/challenges/apple-and-orange/problem)  
**Dificuldade**: Easy
**Link**: [URL do problema](https://www.hackerrank.com/challenges/apple-and-orange/problem)

### Enunciado

Determine o número de maçãs e laranjas que caem na casa de Sam, dada a posição das árvores e as distâncias que as frutas caem.

## Solução

### Abordagem

A abordagem é iterar sobre cada array de frutas (maçãs e laranjas). Para cada fruta, calculamos sua posição final somando a distância que ela caiu à posição de sua respectiva árvore. Em seguida, verificamos se essa posição final está dentro do intervalo inclusivo da casa de Sam `[s, t]`. Contamos quantas frutas de cada tipo atendem a essa condição e imprimimos os totais.

### Complexidade

- **Tempo**: O(m + n), onde `m` é o número de maçãs e `n` é o número de laranjas. Isso ocorre porque precisamos percorrer ambos os arrays uma vez.
- **Espaço**: O(1), pois usamos apenas algumas variáveis para armazenar as contagens, sem alocar espaço adicional proporcional ao tamanho da entrada.

## Implementação

```go
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    var appleCount int32
    for _, appleDist := range apples {
        if pos := a + appleDist; pos >= s && pos <= t {
            appleCount++
        }
    }
    fmt.Println(appleCount)

    var orangeCount int32
    for _, orangeDist := range oranges {
        if pos := b + orangeDist; pos >= s && pos <= t {
            orangeCount++
        }
    }
    fmt.Println(orangeCount)
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit

## Aprendizados

- Reforcei o uso de `for range` para iterar sobre slices.
- Pratiquei a lógica condicional para verificar se um número está dentro de um intervalo.
- O problema era simples e serviu para consolidar a sintaxe básica de Go.
