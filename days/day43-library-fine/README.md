# Dia 43 - library-fine

## Problema

**Plataforma**: HackerRank
**Dificuldade**: Easy
**Link**: https://www.hackerrank.com/challenges/library-fine

### Enunciado

Calcular multa de biblioteca baseada na data de retorno vs data esperada:
- Sem multa: retorno até a data esperada
- 15 Hackos/dia: atraso no mesmo mês/ano
- 500 Hackos/mês: atraso no mesmo ano
- 10000 Hackos: atraso após o ano esperado

## Solução

### Abordagem

Comparação condicional hierárquica:
1. Verificar se não há multa (retorno no prazo)
2. Verificar atraso por ano (multa fixa 10000)
3. Verificar atraso por mês (500 × diferença de meses)
4. Calcular atraso por dias (15 × diferença de dias)

### Complexidade

- **Tempo**: O(1)
- **Espaço**: O(1)

## Implementação

```go
func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
    if y1 < y2 || (y1 == y2 && m1 < m2) || (y1 == y2 && m1 == m2 && d1 <= d2) {
        return 0
    }

    if y1 > y2 {
        return 10000
    }

    if m1 > m2 {
        return 500 * (m1 - m2)
    }

    return 15 * (d1 - d2)
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Lógica condicional hierárquica para diferentes tipos de penalidade
- Comparação de datas usando múltiplas condições
- Estrutura de if-else para casos mutuamente exclusivos
- Aritmética simples para cálculo de diferenças

## Screenshots

[Problema aceito com sucesso]
