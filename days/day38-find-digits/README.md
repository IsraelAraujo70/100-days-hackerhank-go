# Dia 38 - find-digits

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/find-digits

### Enunciado

Dado um inteiro n, para cada dígito que compõe o inteiro, determinar se ele é um divisor. Contar o número de divisores dentro do inteiro.

## Solução

### Abordagem

Extrair cada dígito usando módulo 10 e divisão por 10. Para cada dígito não-zero, verificar se divide n evenly (n % digit == 0).

### Complexidade

- **Tempo**: O(d) onde d é número de dígitos
- **Espaço**: O(1)

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Extração de dígitos com operações aritméticas: temp % 10 para dígito atual, temp /= 10 para próximo
- Verificação de divisibilidade com módulo
- Cuidado especial com divisão por zero (digit != 0)
