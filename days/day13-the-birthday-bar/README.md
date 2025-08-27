# Dia 13 - The Birthday Bar

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy
**Link**: [https://www.hackerrank.com/challenges/the-birthday-bar](https://www.hackerrank.com/challenges/the-birthday-bar)

### Enunciado

Lily tem uma barra de chocolate com quadrados, e cada quadrado tem um número inteiro. Ela quer compartilhar um segmento contíguo da barra com Ron para seu aniversário. O segmento deve ter um comprimento igual ao mês de nascimento de Ron (`m`) e a soma dos números nos quadrados deve ser igual ao seu dia de nascimento (`d`). A tarefa é contar de quantas maneiras ela pode dividir o chocolate.

## Solução

### Abordagem

A solução utiliza uma abordagem de "janela deslizante" (sliding window) para encontrar os segmentos válidos.

1.  Iteramos pela lista de chocolates `s`, começando do índice `0` até o último ponto possível onde um segmento de comprimento `m` pode começar (`len(s) - m`).
2.  Para cada posição inicial `i`, um segundo loop calcula a soma dos `m` elementos seguintes (da posição `i` até `i + m - 1`).
3.  Se a `soma` calculada for igual ao dia do aniversário `d`, um contador de segmentos válidos é incrementado.
4.  Ao final, o contador é retornado.

### Complexidade

-   **Tempo**: `O(n * m)`, onde `n` é o número de quadrados na barra de chocolate e `m` é o comprimento do segmento. Isso ocorre porque para cada um dos `n-m+1` segmentos possíveis, fazemos `m` somas.
-   **Espaço**: `O(1)`, pois usamos apenas um número fixo de variáveis para armazenar a soma e a contagem, independentemente do tamanho da entrada.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit

## Aprendizados

-   **Sliding Window**: Este problema é um bom exemplo de aplicação do padrão de "janela deslizante". A implementação com loops aninhados é a forma mais direta, embora uma otimização para `O(n)` seja possível (calculando a soma da primeira janela e, em seguida, deslizando-a subtraindo o elemento que sai e adicionando o que entra).
-   **Limites de Loop**: É crucial definir corretamente o limite do loop externo (`len(s) - m`) para evitar erros de "index out of bounds" ao acessar os elementos do array.
-   **Conversão de Tipos**: O problema exige a manipulação de `int32`, o que reforça a atenção aos tipos de dados em Go.

## Screenshots

*Adicionar print do "Accepted" aqui*

