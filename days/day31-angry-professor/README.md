# Dia 31 - Angry Professor

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/angry-professor/problem

### Enunciado

Um professor de Matemática cancela a aula se menos que `k` alunos estiverem presentes no horário. Dadas as chegadas dos alunos `a[i]` (onde `a[i] <= 0` significa no horário/adiantado e `a[i] > 0` significa atraso), determine se a aula será cancelada.

Retorne `YES` se a aula for cancelada, caso contrário `NO`.

## Solução

### Abordagem

- Contar quantos alunos chegaram no horário: `a[i] <= 0`.
- Se a contagem for menor que `k`, retornar `YES` (cancelada); senão, `NO`.

### Complexidade

- Tempo: O(n), onde n é o número de alunos.
- Espaço: O(1).

## Implementação

- Código: `days/day31-angry-professor/main.go`

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Leitura cuidadosa do enunciado: o retorno é `YES` quando cancela (contagem < k), não o contrário.

## Screenshots

Adicione aqui o print do "Accepted" quando disponível.
