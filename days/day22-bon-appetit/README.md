# Dia 22 - Bon Appetit

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/bon-appetit/problem

### Enunciado

Dois amigos (Anna e Brian) dividem a conta do jantar. Anna não comeu o item de índice `k`. Dado o vetor `bill` com os custos, `k` e o valor `b` cobrado para Anna, determine se Brian cobrou corretamente. Se sim, imprimir `Bon Appetit`; caso contrário, imprimir quanto Brian deve devolver (diferença inteira).

## Solução

### Abordagem

Somar todos os itens exceto o de índice `k`, dividir por 2 para obter a parte justa de Anna e comparar com `b`. Se `b` for igual à parte justa, imprimir `Bon Appetit`; caso contrário, imprimir `b - parte_justa`.

### Complexidade

- **Tempo**: O(n)
- **Espaço**: O(1)

## Implementação
Arquivo: `days/day22-bon-appetit/main.go`

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforço do padrão de leitura de entrada e impressão direta.
- Lidar com exclusão de um item pelo índice e comparação direta do valor cobrado.
- Validação simples de divisão da conta.

## Screenshots

Sem screenshot neste commit.
