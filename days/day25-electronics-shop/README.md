# Dia 25 - Electronics Shop

## Problema

- Plataforma: HackerRank  
- Dificuldade: Easy  
- Link: https://www.hackerrank.com/challenges/electronics-shop/problem

### Enunciado

Dado um orçamento `b` e duas listas de preços (`keyboards` e `drives`), determine o maior valor possível da soma de um teclado e um USB drive que não ultrapasse o orçamento. Se não for possível comprar ambos, retorne `-1`.

## Solução

### Abordagem

- Brute force simples: iterar por todas as combinações de `keyboards[i] + drives[j]` e manter o máximo `<= b`.
- Como as constraints do problema permitem, a solução `O(n*m)` é suficiente e direta. Alternativamente, poderia ordenar e usar two-pointers para melhorar para `O(n log n + m log m)` no pior caso, mas não é necessário aqui.

### Complexidade

- Tempo: O(n*m)
- Espaço: O(1)

## Implementação

- Código: `days/day25-electronics-shop/main.go`

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforço da checagem de máximo condicionado por orçamento.
- Padrão de leitura de entrada/saída no template do HackerRank para Go.

## Screenshots

- Ver `screenshots/day25-electronics-shop/` (adicionar quando disponível)
