# Dia 33 - Utopian Tree

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: https://www.hackerrank.com/challenges/utopian-tree/problem

### Enunciado

A árvore utópica passa por dois ciclos de crescimento por ano: na primavera ela dobra de altura e no verão cresce 1 metro. Partindo de uma muda de 1 metro de altura, determine a altura da árvore após `n` ciclos de crescimento. Para cada caso de teste é informado o número de ciclos a simular.

## Solução

### Abordagem

Simulei os ciclos sequencialmente mantendo a altura atual. Para cada ciclo par (primavera) dobro a altura; para cada ciclo ímpar (verão) incremento 1 metro. O índice começa em zero para que o primeiro ciclo seja a primavera. Ao final retorno a altura acumulada.

### Complexidade

- **Tempo**: O(n), percorrendo todos os ciclos
- **Espaço**: O(1), apenas variáveis escalares

## Implementação

- Código: `days/day33-utopian-tree/main.go`

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforcei o padrão de simulação simples quando o problema descreve estados alternados.
- Atenção para iniciar o contador corretamente para casar primavera/verão.
- Uso de `int32` para alinhar com a assinatura imposta pelo HackerRank.

## Screenshots

_Aguardando captura do resultado Accepted._
