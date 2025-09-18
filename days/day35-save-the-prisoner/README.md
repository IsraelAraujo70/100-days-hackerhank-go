# Dia 35 - Save the Prisoner

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: [Save the Prisoner](https://www.hackerrank.com/challenges/save-the-prisoner/problem)

### Enunciado

Uma prisão tem `n` prisioneiros e `m` doces para distribuir. Os prisioneiros estão sentados em uma mesa circular com cadeiras numeradas de 1 a n. A distribuição começa na cadeira `s` e continua sequencialmente até todos os doces serem distribuídos. O último doce tem sabor horrível. Determine qual prisioneiro receberá esse doce.

## Solução

### Abordagem

Este é um problema clássico de aritmética modular em arranjo circular.

**Intuição**: 
- Começamos na posição `s`
- Distribuímos `m` doces sequencialmente 
- O último doce vai para a posição `s + m - 1` (se fosse linear)
- Como é circular, usamos módulo `n`
- Convertemos entre indexação 1-based e 0-based

**Fórmula**: `((s - 1 + m - 1) % n) + 1`

### Complexidade

- **Tempo**: O(1) - cálculo direto com aritmética modular
- **Espaço**: O(1) - apenas variáveis auxiliares

## Implementação

```go
func saveThePrisoner(n int32, m int32, s int32) int32 {
    return ((s-1+m-1)%n)+1
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Aritmética modular é fundamental para problemas circulares
- Conversão entre indexação 0-based e 1-based requer cuidado
- Problemas que parecem de simulação podem ter soluções matemáticas O(1)
- A fórmula `(start + offset) % size` é padrão para arranjos circulares

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
