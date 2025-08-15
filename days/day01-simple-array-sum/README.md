# Dia 1 - Simple Array Sum

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/simple-array-sum/

### Enunciado

Dado um array de inteiros, retorne a soma de seus elementos.

## Solução

### Abordagem

- Somar todos os elementos do array iterando com `range`, descartando o índice com `_` para não armazená-lo na memória.
- Exemplo do laço usado:

```go
for _, v := range ar {
    sum += int(v)
}
```

### Complexidade

- **Tempo**: O(n)
- **Espaço**: O(1)

## Implementação

Função `simpleArraySum(ar []int32) int32` em `days/day01-simple-array-sum/main.go`.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Agradecimentos

Agradecimentos especiais ao Hitesh Choudhary pelo curso gratuito de Golang no YouTube, que tem ajudado bastante neste início.

## Aprendizados

- No Go é possível omitir o índice no `for range` usando `_`, por exemplo `for _, v := range ar`, evitando manter esse valor na memória quando não é necessário.
- Curti muito a sintaxe do Go; até agora achei bem direta e prazerosa de escrever.

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
