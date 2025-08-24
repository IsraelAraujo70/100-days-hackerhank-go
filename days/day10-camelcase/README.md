# Dia 10 - camelcase

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: [CamelCase](https://www.hackerrank.com/challenges/camelcase/problem)

### Enunciado

Dada uma string em CamelCase, determinar o número de palavras na string. Uma string CamelCase tem as seguintes propriedades:

- É uma concatenação de uma ou mais palavras consistindo de letras inglesas
- Todas as letras da primeira palavra são minúsculas
- Para cada palavra subsequente, a primeira letra é maiúscula e o resto são minúsculas

**Exemplo**: `oneTwo Three` contém 3 palavras: 'one', 'Two', 'Three'

## Solução

### Abordagem

A solução conta as letras maiúsculas na string e adiciona 1 (para a primeira palavra que sempre começa com minúscula). Cada letra maiúscula indica o início de uma nova palavra.

1. Converter a string em slice de runes para lidar corretamente com caracteres Unicode
2. Iterar pelos caracteres e contar quantos são maiúsculos usando `unicode.IsUpper()`
3. Retornar o count + 1 (primeira palavra)

### Complexidade

- **Tempo**: O(n) - percorre a string uma vez
- **Espaço**: O(n) - cria slice de runes da string

## Implementação

```go
func camelcase(s string) int32 {
    char := []rune(s)
    wordCount := 1
    for _, v := range char {
        if unicode.IsUpper(v) {
            wordCount++
        }
    }
    return int32(wordCount)
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Uso do pacote `unicode` para verificar se caracteres são maiúsculos/minúsculos
- Conversão de string para `[]rune` para lidar corretamente com caracteres Unicode
- Algoritmo simples de contagem - cada letra maiúscula = nova palavra
- Padrão `for _, v := range` aplicado novamente para iteração em slice

## Screenshots

[Problema resolvido com sucesso no HackerRank]
