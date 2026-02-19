# Dia 8 - birthday-cake-candles

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: [Birthday Cake Candles](https://www.hackerrank.com/challenges/birthday-cake-candles/problem)

### Enunciado

Você está encarregado do bolo para o aniversário de uma criança. O bolo terá uma vela para cada ano de idade. A criança só consegue apagar as velas mais altas. Sua tarefa é contar quantas velas são as mais altas.

**Exemplo**: Para velas com alturas [3, 2, 1, 3], as velas mais altas têm 3 unidades de altura e há 2 delas, então a função deve retornar 2.

## Solução

### Abordagem

1. Encontrar a altura máxima entre todas as velas percorrendo o array uma vez
2. Percorrer o array novamente contando quantas velas têm essa altura máxima
3. Retornar a contagem

### Complexidade

- **Tempo**: O(n) - duas passadas pelo array
- **Espaço**: O(1) - apenas variáveis auxiliares

## Implementação

```go
func birthdayCakeCandles(candles []int32) int32 {
    biggest := candles[0]
    count := 0
    for _, value := range candles {
        if value >= biggest {
            biggest = value
        }
    }
    for _, value := range candles {
        if value == biggest {
            count++
        }
    }
    return int32(count)
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Problema simples de busca por máximo e contagem
- Abordagem de duas passadas é clara e eficiente
- Prática com range loops em Go

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
