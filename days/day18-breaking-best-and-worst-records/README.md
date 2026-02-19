# Day 18: Breaking Best and Worst Records

## Enunciado

Maria joga basquete universitário e quer se tornar profissional. A cada temporada, ela mantém um registro de seu jogo. Ela tabula o número de vezes que quebra seu recorde da temporada de maior e menor número de pontos em um jogo. Os pontos marcados no primeiro jogo estabelecem seu recorde para a temporada, e ela começa a contar a partir daí.

[Link para o problema](https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem)

## Abordagem

A solução para este problema é um algoritmo de rastreamento de estado simples com complexidade de tempo O(n).

1.  **Inicialização**:
    *   Inicialize as variáveis `maxScore` e `minScore` com o valor da primeira pontuação (`scores[0]`).
    *   Inicialize os contadores `maxBreaks` e `minBreaks` como 0.

2.  **Iteração**:
    *   Percorra o array de pontuações a partir do segundo elemento (índice 1).
    *   Para cada pontuação:
        *   Se a pontuação atual for **estritamente maior** que `maxScore`, atualize `maxScore` e incremente `maxBreaks`.
        *   Se a pontuação atual for **estritamente menor** que `minScore`, atualize `minScore` e incremente `minBreaks`.

3.  **Retorno**:
    *   Retorne um array contendo `maxBreaks` e `minBreaks`.

Essa abordagem garante que o array seja percorrido apenas uma vez, tornando-a eficiente.

## Complexidade

- **Tempo**: O(n), onde `n` é o número de jogos. O array de pontuações é percorrido uma única vez.
- **Espaço**: O(1), pois usamos apenas um número constante de variáveis para armazenar os estados (`maxScore`, `minScore`, `maxBreaks`, `minBreaks`), independentemente do tamanho da entrada.

## O que aprendi

- Reforcei a lógica de rastreamento de estado mínimo/máximo em uma única passagem.
- A importância de ler atentamente os requisitos (por exemplo, "estritamente maior/menor") para evitar erros de contagem.
- Pratiquei a inicialização correta de valores de referência antes de iniciar um loop de comparação.