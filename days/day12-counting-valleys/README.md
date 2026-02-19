# Dia 12 - counting-valleys

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/counting-valleys

### Enunciado

Um caminhante mantém registros meticulosos de suas caminhadas. Durante a última caminhada de exatamente `steps` passos, foi anotado se cada passo foi uma subida ('U') ou descida ('D'). As caminhadas sempre começam e terminam ao nível do mar, e cada passo representa uma mudança de 1 unidade na altitude.

- **Vale**: sequência de passos consecutivos abaixo do nível do mar, começando com uma descida do nível do mar e terminando com uma subida ao nível do mar
- **Montanha**: sequência de passos consecutivos acima do nível do mar, começando com uma subida do nível do mar e terminando com uma descida ao nível do mar

O objetivo é contar quantos vales foram atravessados durante a caminhada.

## Solução

### Abordagem

1. Mantenho uma variável `altitude` para rastrear a altitude atual
2. Para cada passo:
   - Se for 'D': diminuo a altitude
   - Se for 'U': aumento a altitude
3. Um vale é contado quando retornamos ao nível do mar (altitude == 0) com um passo 'U'
4. Isso garante que contamos apenas quando saímos de um vale

### Complexidade

- **Tempo**: O(n) - percorro o caminho uma única vez
- **Espaço**: O(n) - conversão da string para array de runes

## Implementação

```go
func countingValleys(steps int32, path string) int32 {
    pathArr := []rune(path)
    var altitude, valleyCount int32
    for _, v := range pathArr {
        if v == 'D' {
            altitude--
        } else {
            altitude++
        }
        if altitude == 0 && v == 'U' {
            valleyCount++
        }
    }
    return valleyCount
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Algoritmo de tracking de estado simples mas efetivo
- Importância de identificar exatamente quando contar (ao sair do vale, não ao entrar)
- Uso de `[]rune(path)` para trabalhar com caracteres Unicode de forma segura
- Condição `altitude == 0 && v == 'U'` garante contagem precisa dos vales

## Screenshots

[Print do resultado "Accepted" será adicionado]
