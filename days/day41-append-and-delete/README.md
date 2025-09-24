# Dia 41 - append-and-delete

## Problema

**Plataforma**: HackerRank
**Dificuldade**: Easy
**Link**: https://www.hackerrank.com/challenges/append-and-delete

### Enunciado

Você tem duas strings de letras minúsculas em inglês. Pode realizar dois tipos de operações na primeira string:
1. Anexar uma letra minúscula no final da string
2. Deletar o último caractere da string (string vazia permanece vazia)

Dado um inteiro k e duas strings s e t, determine se é possível converter s em t executando exatamente k operações.

## Solução

### Abordagem

1. **Encontrar prefixo comum**: Determinar quantos caracteres iniciais são iguais entre s e t
2. **Calcular operações mínimas**: (len(s) - prefixo_comum) + (len(t) - prefixo_comum)
3. **Verificar viabilidade**:
   - Se k < operações_mínimas: impossível
   - Se k == operações_mínimas: possível
   - Se k > operações_mínimas: verificar se operações extras podem ser desperdiçadas
4. **Desperdiçar operações**:
   - Deletar string inteira e reconstruir (len(s) + len(t) operações)
   - Ou desperdiçar em pares (operações extras devem ser pares)

### Complexidade

- **Tempo**: O(min(len(s), len(t))) - para encontrar prefixo comum
- **Espaço**: O(1) - apenas variáveis auxiliares

## Implementação

A solução encontra o prefixo comum máximo e calcula se é possível usar exatamente k operações, considerando que operações extras podem ser desperdiçadas em pares ou deletando/reconstruindo a string inteira.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Lógica de transformação de strings com operações limitadas
- Estratégia de desperdiçar operações extras (em pares ou reconstrução completa)
- Importância de encontrar prefixo comum para otimizar operações
- Análise de casos extremos (string vazia, operações insuficientes)

## Screenshots

[Print do resultado "Accepted" seria adicionado aqui]
