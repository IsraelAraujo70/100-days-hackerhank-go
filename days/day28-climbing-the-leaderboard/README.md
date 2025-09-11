# Dia 28 - Climbing the Leaderboard

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟡 Medium  
**Link**: https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem

### Enunciado

Um jogo usa ranking denso (dense ranking). Dada a leaderboard com scores em ordem decrescente (com possíveis duplicados) e os scores da jogadora em ordem crescente, calcule a posição (rank) dela após cada partida, aplicando a regra de ranking denso: scores iguais têm o mesmo rank e o próximo rank é imediatamente seguinte.

## Solução

### Abordagem

- Compactar a leaderboard removendo duplicados, mantendo ordem decrescente (array `uniq`).
- Usar um ponteiro a partir do final de `uniq` (menores scores) e percorrer os `player` em ordem crescente.
- Para cada `p`, mover o ponteiro para a esquerda enquanto `p >= uniq[i]`.
- O rank é `i + 2` (1-based e contando quantos são estritamente maiores).

Racional: como os scores de `player` estão em ordem crescente, o ponteiro nunca volta, resultando em varredura linear no total.

### Complexidade

- Tempo: O(n + m), onde n = tamanho de `ranked` e m = tamanho de `player`.
- Espaço: O(n) para o array de únicos.

## Implementação

- Código: `days/day28-climbing-the-leaderboard/main.go`

Trecho principal:

```go
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    uniq := make([]int32, 0, len(ranked))
    for _, s := range ranked {
        if len(uniq) == 0 || s != uniq[len(uniq)-1] {
            uniq = append(uniq, s)
        }
    }

    res := make([]int32, 0, len(player))
    i := len(uniq) - 1
    for _, p := range player {
        for i >= 0 && p >= uniq[i] {
            i--
        }
        res = append(res, int32(i+2))
    }
    return res
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Técnica de dois ponteiros aproveitando ordem dos dados para obter O(n+m).
- Definição de ranking denso e cálculo do rank com base em "quantos são estritamente maiores".

## Screenshots

- Adicionar print do resultado "Accepted" quando disponível.
