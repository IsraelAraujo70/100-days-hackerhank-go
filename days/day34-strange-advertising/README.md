# Dia 34 - Strange Advertising

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: https://www.hackerrank.com/challenges/strange-advertising/problem

### Enunciado

A campanha começa com um anúncio compartilhado para 5 pessoas. A cada dia, apenas a metade (arredondada para baixo) dessas pessoas curte o post, e cada curtida gera 3 novos compartilhamentos para o dia seguinte. Ninguém recebe o anúncio duas vezes. O objetivo é calcular o total acumulado de curtidas após `n` dias.

## Solução

### Abordagem

Simulei o processo dia a dia mantendo duas variáveis: `shared`, quantidade de pessoas alcançadas no dia, e `cumulative`, total de curtidas já obtidas. Para cada dia calculo `likes = shared / 2`, somo ao total e atualizo `shared` para `likes * 3`, que alimenta o próximo dia. O laço roda `n` vezes, seguindo a ordem descrita pelo enunciado.

### Complexidade

- **Tempo**: O(n), percorrendo cada dia uma única vez
- **Espaço**: O(1), uso apenas variáveis escalares

## Implementação

- Código: `days/day34-strange-advertising/main.go`

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforcei a leitura do enunciado para respeitar o arredondamento para baixo ao dividir por 2.
- Mantive o uso de `int32` para alinhar com a assinatura imposta pelo HackerRank.
- A atualização sequencial das variáveis elimina a necessidade de armazenar históricos intermediários.

## Screenshots

_Aguardando captura do resultado Accepted._
