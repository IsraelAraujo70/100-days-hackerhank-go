# Dia 15 - Day of the Programmer

## Problema

**Plataforma**: HackerRank
**Dificuldade**: Easy
**Link**: [Day of the Programmer](https://www.hackerrank.com/challenges/day-of-the-programmer/problem)

### Enunciado

Encontrar a data do 256º dia do ano na Rússia, considerando a mudança do calendário Juliano para o Gregoriano.

- **1700-1917 (Juliano)**: Ano bissexto se for divisível por 4.
- **1918 (Transição)**: O dia seguinte a 31 de janeiro foi 14 de fevereiro.
- **1919+ (Gregoriano)**: Ano bissexto se for divisível por 400, ou divisível por 4 mas não por 100.

A saída deve ser no formato `dd.mm.yyyy`.

## Solução

### Abordagem

A solução foi implementada com uma série de condicionais para tratar os três casos distintos do calendário russo:

1.  **Ano da Transição (1918)**: Este é um caso especial. Em 1918, 13 dias foram pulados em fevereiro. O 256º dia do ano cai em 26 de setembro. Este caso é tratado com um `if` específico que retorna `26.09.1918` diretamente.

2.  **Calendário Juliano (1700-1917)**: Para este intervalo, um ano é bissexto se for divisível por 4. Se for bissexto, o 256º dia é 12 de setembro; caso contrário, é 13 de setembro.

3.  **Calendário Gregoriano (1919 em diante)**: A regra do ano bissexto é mais complexa: é bissexto se for divisível por 400, ou se for divisível por 4 e não por 100. A lógica para determinar o dia (12 ou 13 de setembro) é a mesma do calendário juliano.

A função determina se o ano é bissexto com base nessas regras e, em seguida, formata a string da data correspondente.

### Complexidade

-   **Tempo**: O(1), pois a lógica envolve apenas algumas verificações condicionais e operações aritméticas, sem loops ou recursão dependente do tamanho da entrada.
-   **Espaço**: O(1), pois a memória usada é constante e não depende do ano de entrada.

## Implementação

```go
func dayOfProgrammer(year int32) string {
	if year == 1918 {
		return fmt.Sprintf("26.09.%d", year)
	}

	bissexto := false
	if year < 1918 {
		if year%4 == 0 {
			bissexto = true
		}
	} else {
		if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
			bissexto = true
		}
	}

	if bissexto {
		return fmt.Sprintf("12.09.%d", year)
	}
	return fmt.Sprintf("13.09.%d", year)
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit

## Aprendizados

- Reforcei o conhecimento sobre as regras dos calendários Juliano e Gregoriano.
- Pratiquei o uso de `fmt.Sprintf` para formatação de strings em Go.
- A principal dificuldade foi entender e calcular corretamente o caso especial do ano de 1918, onde a contagem de dias é diferente devido à transição do calendário.

## Screenshots

*A ser adicionado após a submissão no HackerRank.*