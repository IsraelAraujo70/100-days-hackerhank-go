# Dia 27 - Picking Numbers

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: https://www.hackerrank.com/challenges/picking-numbers/problem

### Enunciado

Dado um array de inteiros, encontre o maior multiconjunto (subconjunto considerando multiplicidades) em que a diferença absoluta entre quaisquer dois elementos seja menor ou igual a 1. Retorne o tamanho máximo desse multiconjunto.

Exemplos de entradas mostram que o conjunto ótimo sempre é formado por elementos iguais a `x` e/ou `x+1` para algum `x` presente no array.

## Solução

### Abordagem

- Conte as frequências de cada valor.
- Para cada valor `v`, considere o tamanho do conjunto formado por `v` e `v+1`, que é `freq[v] + freq[v+1]`.
- O resultado é o máximo desses somatórios ao percorrer todos os `v` presentes.

Racional: se a diferença entre quaisquer dois elementos é ≤ 1, então todos os elementos pertencem a dois valores consecutivos no máximo.

### Complexidade

- Tempo: O(n), onde n é o tamanho do array (contagem + varredura das chaves)
- Espaço: O(k), número de valores distintos (no HackerRank, valores estão em faixa pequena)
- Código: `days/day27-picking-numbers/main.go`

## Implementação

Trecho principal da função:

```go
func pickingNumbers(a []int32) int32 {
    freq := make(map[int32]int32)
    for _, v := range a {
        freq[v]++
    }

    var best int32 = 0
    for v, c := range freq {
        c2 := freq[v+1]
        if c+c2 > best {
            best = c + c2
        }
    }
    return best
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforço do padrão de solução via mapa de frequência e regra de negócio (valores consecutivos) para reduzir busca.
- Verificação simples de emparelhamento `v` e `v+1` cobre todos os conjuntos válidos.

## Screenshots

- Adicionar print do resultado "Accepted" quando disponível.
