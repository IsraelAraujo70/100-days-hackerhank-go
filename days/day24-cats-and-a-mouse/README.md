# Dia 24 - Cats and a Mouse

## Problema

- Plataforma: HackerRank  
- Dificuldade: 🟢 Easy  
- Link: https://www.hackerrank.com/challenges/cats-and-a-mouse/problem

### Enunciado

Dados três posições em uma reta: `x` (gato A), `y` (gato B) e `z` (rato C), determine qual gato alcança o rato primeiro assumindo mesma velocidade. Se ambos chegam ao mesmo tempo, o rato escapa.

## Solução

### Abordagem

- Calcular as distâncias absolutas até o rato: `|x - z|` e `|y - z|`.
- Se `|x - z| < |y - z|` → "Cat A".
- Se `|y - z| < |x - z|` → "Cat B".
- Se iguais → "Mouse C".

### Complexidade

- Tempo: O(1)
- Espaço: O(1)

## Implementação

- Código: `days/day24-cats-and-a-mouse/main.go`

## Resultado

- [x] Problema resolvido

## Aprendizados

- Diferença absoluta simples resolve; atenção a empate.

## Screenshots

- Adicionar print do Accepted quando submetido.
