# Dia 20 - Migratory Birds

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: [https://www.hackerrank.com/challenges/migratory-birds](https://www.hackerrank.com/challenges/migratory-birds)

### Enunciado

Given an array of bird sightings where every element represents a bird type id, determine the id of the most frequently sighted type. If more than 1 type has been spotted that maximum amount, return the smallest of their ids.

## Solução

### Abordagem

Utilizei um slice de inteiros como um mapa de frequência para contar as ocorrências de cada tipo de pássaro, aproveitando que os tipos são limitados (1 a 5).

1.  Criei um slice `freq` de tamanho 6 para mapear cada tipo de pássaro (índices 1 a 5) à sua contagem.
2.  Iterei sobre o array de entrada `arr`, incrementando a contagem em `freq[birdType]` para cada pássaro.
3.  Em um segundo loop, iterei de 1 a 5 para encontrar a frequência máxima.
4.  Mantive variáveis para a frequência máxima (`maxFreq`) e o ID do pássaro correspondente (`mostCommonBird`).
5.  A frequência e o ID só são atualizados se uma contagem **estritamente maior** for encontrada. Isso garante que, em caso de empate, o primeiro ID encontrado (que é o menor, devido à ordem crescente do loop) seja mantido.

### Complexidade

-   **Tempo**: O(n), onde `n` é o número de avistamentos. O primeiro loop é O(n) para contar as frequências, e o segundo é O(1) (tamanho constante de 5), resultando em uma complexidade dominada pelo primeiro loop.
-   **Espaço**: O(1), pois o slice de frequência tem um tamanho constante (6), independentemente do tamanho da entrada.

## Implementação

```go
func migratoryBirds(arr []int32) int32 {
    // Usa um slice como mapa de frequência para os tipos de 1 a 5.
    freq := make([]int32, 6)

    // Conta a frequência de cada tipo de pássaro.
    for _, birdType := range arr {
        freq[birdType]++
    }

    var maxFreq int32 = 0
    var mostCommonBird int32 = 0

    // Itera dos tipos 1 a 5 para encontrar a frequência máxima.
    // A ordem crescente garante que o menor ID seja retornado em caso de empate,
    // pois a atualização só ocorre para frequências estritamente maiores.
    for birdType := 1; birdType < len(freq); birdType++ {
        if freq[birdType] > maxFreq {
            maxFreq = freq[birdType]
            mostCommonBird = int32(birdType)
        }
    }

    return mostCommonBird
}
```

## Resultado

-   [x] Problema resolvido
-   [x] Accepted no primeiro submit

## Aprendizados

-   Reforcei o uso de slices como mapas de frequência para chaves inteiras limitadas, uma abordagem mais performática que `map` para este caso.
-   A ordem de iteração pode ser uma ferramenta simples e poderosa para lidar com regras de desempate (como "retorne o menor ID em caso de empate").