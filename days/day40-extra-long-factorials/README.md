# Dia 40 - extra-long-factorials

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Medium  
**Link**: [Extra Long Factorials](https://www.hackerrank.com/challenges/extra-long-factorials)

### Enunciado

Calcular e imprimir o fatorial de um número inteiro dado. Para números grandes (como 25!), o resultado não cabe em tipos de dados comuns como int64, sendo necessário usar big integers.

Exemplo: 25! = 15511210043330985984000000

## Solução

### Abordagem

Utilizei o pacote `math/big` do Go para trabalhar com números arbitrariamente grandes:
1. Inicializei um `big.Int` com valor 1
2. Multipliquei sequencialmente por todos os números de 2 até n
3. Imprimi o resultado final

### Complexidade

- **Tempo**: O(n * log(n!)) - devido às operações de multiplicação em big integers
- **Espaço**: O(log(n!)) - para armazenar o resultado

## Implementação

```go
func extraLongFactorials(n int32) {
    result := big.NewInt(1)
    
    for i := int32(2); i <= n; i++ {
        result.Mul(result, big.NewInt(int64(i)))
    }
    
    fmt.Println(result)
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Aprendi sobre o pacote `math/big` para trabalhar com números grandes em Go
- Entendi quando usar big integers vs tipos primitivos
- O método `Mul` modifica o receptor, permitindo operações in-place eficientes
- Go não tem sobrecarga de operadores, então usamos métodos como `Mul`, `Add`, etc.

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
