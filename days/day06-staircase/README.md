# #100DiasDeHackerRank — Dia 6/100 (Go)

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/staircase/

### Enunciado

Staircase - Imprimir uma escada de tamanho n usando # símbolos e espaços. A escada deve ser alinhada à direita.

**Exemplo para n=4**:
```
   #
  ##
 ###
####
```

**Função**: `staircase(n int32)` (não retorna valor, apenas imprime)
- **Parâmetro**: inteiro n (tamanho da escada)
- **Ação**: Imprimir escada alinhada à direita

**Constraints**: 0 < n ≤ 100

## Solução

### Abordagem

• Para cada linha i (de 1 até n):
  - Imprimir (n-i) espaços em branco
  - Imprimir i símbolos #
• Usar `strings.Repeat()` para gerar as strings de espaços e #
• Concatenar as strings e imprimir com `fmt.Println()`

### Complexidade

- **Tempo**: O(n²) - loop O(n) com operações strings.Repeat O(n)
- **Espaço**: O(n) - string temporária para cada linha

## Implementação

Solução usando `strings.Repeat()` para construir cada linha da escada de forma limpa e legível.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Função `strings.Repeat(string, count)` para repetir strings
- Concatenação de strings com `+=`
- Problemas de formatação/impressão não retornam valores
- Alinhamento à direita: calcular espaços como (total - posição atual)
- Conversão int32 para int quando necessário para funções da stdlib

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
