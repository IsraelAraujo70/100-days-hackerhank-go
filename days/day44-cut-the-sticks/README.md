# Dia 44 - Cut the Sticks

## Problema

**Plataforma**: HackerRank
**Dificuldade**: Easy
**Link**: https://www.hackerrank.com/challenges/cut-the-sticks

### Enunciado

Dado um array de bastões com diferentes comprimentos, você deve iterativamente cortar os bastões em pedaços menores, descartando os pedaços mais curtos até que não reste nenhum. A cada iteração:
1. Determine o comprimento do bastão mais curto restante
2. Corte esse comprimento de todos os bastões mais longos
3. Descarte todos os pedaços desse comprimento mais curto
4. Conte quantos bastões restam antes de cada corte

## Solução

### Abordagem

Simulação iterativa do processo de corte:
1. A cada iteração, contar quantos bastões ainda restam
2. Encontrar o comprimento mínimo entre os bastões restantes
3. Subtrair esse mínimo de todos os bastões
4. Remover todos os bastões que ficaram com comprimento zero
5. Repetir até não restarem bastões

### Complexidade

- **Tempo**: O(n²) no pior caso (n iterações × n elementos para encontrar mínimo)
- **Espaço**: O(n) para armazenar o array de resultados e array temporário

## Implementação

```go
func cutTheSticks(arr []int32) []int32 {
    var result []int32

    for len(arr) > 0 {
        // Contar bastões antes do corte
        result = append(result, int32(len(arr)))

        // Encontrar comprimento mínimo
        minLength := arr[0]
        for _, stick := range arr {
            if stick < minLength {
                minLength = stick
            }
        }

        // Cortar e remover bastões zerados
        var newArr []int32
        for _, stick := range arr {
            newStick := stick - minLength
            if newStick > 0 {
                newArr = append(newArr, newStick)
            }
        }

        arr = newArr
    }

    return result
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- **Simulação iterativa**: Implementar passo a passo exatamente o que o problema descreve
- **Busca de mínimo**: Percorrer array para encontrar menor elemento
- **Filtragem de arrays**: Criar novo array removendo elementos que atendem critério
- **Remoção em Go**: Usar slice vazio e append apenas elementos válidos
