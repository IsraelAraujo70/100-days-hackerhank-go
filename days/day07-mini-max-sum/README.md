# #100DiasDeHackerRank — Dia 7/100 (Go)

## Problema

**Plataforma**: HackerRank
**Dificuldade**: Easy
**Link**: https://www.hackerrank.com/challenges/mini-max-sum/

### Enunciado

Mini-Max Sum - Dado 5 inteiros positivos, encontrar os valores mínimo e máximo que podem ser calculados somando exatamente 4 dos 5 inteiros.

**Exemplo**: Para [1, 2, 3, 4, 5]
- Soma mínima (excluindo 5): 1+2+3+4 = 10
- Soma máxima (excluindo 1): 2+3+4+5 = 14
- Output: "10 14"

**Função**: `miniMaxSum(arr []int32)` (não retorna valor, apenas imprime)
- **Parâmetro**: array de 5 inteiros int32
- **Ação**: Imprimir soma mínima e máxima separadas por espaço

**Constraints**: 1 ≤ arr[i] ≤ 10^9

## Solução

### Abordagem

• Percorrer array uma vez para:
  - Calcular soma total
  - Encontrar maior e menor valor
• Soma mínima: total - maior valor
• Soma máxima: total - menor valor
• Usar int64 para evitar overflow
• Imprimir com `fmt.Printf("%d %d", min, max)`

### Complexidade

- **Tempo**: O(n) - uma única passagem pelo array
- **Espaço**: O(1) - apenas variáveis para soma e extremos

## Implementação

Solução eficiente que calcula tudo em uma única passagem, evitando ordenação.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Otimização de algoritmos: O(n) vs O(n log n)
- Importância de evitar overflow com int64
- Funções que apenas imprimem (sem retorno)
- **Aproveitamento**: Considerei usar sort (O(n log n)) mas optei pela solução O(n) mais eficiente!
- Uso de `fmt.Printf()` para formatação específica
- Sempre pensar na complexidade antes de implementar

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
