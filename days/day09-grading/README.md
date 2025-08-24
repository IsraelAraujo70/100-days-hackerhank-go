# Dia 9 - Grading Students

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: https://www.hackerrank.com/challenges/grading

### Enunciado

HackerLand University tem uma política de arredondamento de notas:

- Toda nota está no range de 0 a 100
- Notas menores que 38 são reprovação
- Se a diferença entre a nota e o próximo múltiplo de 5 é menor que 3, arredondar para cima
- Se a nota é menor que 38, não arredondar (continuará sendo reprovação)

**Exemplos:**
- 84 → 85 (diferença de 1 < 3)
- 29 → 29 (menor que 38)
- 57 → 57 (diferença de 3 ≥ 3)

## Solução

### Abordagem

1. Para cada nota, verificar se é menor que 38 (não arredondar)
2. Calcular o resto da divisão por 5
3. Se o resto é menor que 3, não arredondar
4. Caso contrário, arredondar para o próximo múltiplo de 5

### Complexidade

- **Tempo**: O(n) - percorre o array uma vez
- **Espaço**: O(n) - array de resposta

## Implementação

```go
func gradingStudents(grades []int32) []int32 {
    var response []int32
    for _, v := range grades {
        if v < 38 || v%5 < 3 {
            response = append(response, v)
        } else {
            a := -v%5 + 5
            response = append(response, a+v)
        }
    }
    return response
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Uso inteligente do operador módulo para calcular arredondamento
- Lógica condicional com múltiplas condições no if
- Fórmula matemática: `-v%5 + 5` para calcular quanto somar para próximo múltiplo de 5
