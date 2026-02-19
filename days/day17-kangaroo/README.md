# Dia 17 - Kangaroo

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: [https://www.hackerrank.com/challenges/kangaroo/problem](https://www.hackerrank.com/challenges/kangaroo/problem)

### Enunciado

Dois cangurus em uma reta numérica pulam na direção positiva. Dadas suas posições iniciais (x1, x2) e velocidades (v1, v2), o desafio é determinar se eles se encontrarão no mesmo ponto e ao mesmo tempo.

## Solução

### Abordagem

A solução se baseia em uma equação matemática simples. A posição de cada canguru pode ser descrita como `pos = x + j * v`, onde `j` é o número de pulos. Para que eles se encontrem, as posições devem ser iguais para o mesmo `j`:
`x1 + j*v1 = x2 + j*v2`

Isolando `j`, temos:
`j = (x2 - x1) / (v1 - v2)`

Um encontro só é possível se `j` for um inteiro não-negativo. A lógica implementada verifica duas condições principais:
1.  O canguru que está atrás (`x1`) deve ser estritamente mais rápido que o canguru da frente (`v1 > v2`). Caso contrário, ele nunca o alcançará.
2.  A diferença de posição (`x2 - x1`) deve ser um múltiplo exato da diferença de velocidade (`v1 - v2`). Isso garante que o número de pulos `j` seja um número inteiro.

Se ambas as condições forem satisfeitas, a resposta é "YES". Caso contrário, é "NO".

### Complexidade

- **Tempo**: O(1) - A solução envolve um número constante de operações aritméticas.
- **Espaço**: O(1) - Nenhuma estrutura de dados adicional é utilizada.

## Implementação

O código foi implementado em Go no arquivo `main.go`.

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit

## Aprendizados

- **Modelagem Matemática**: Este problema foi um ótimo exercício para transformar um cenário físico em uma equação matemática. A chave foi converter a pergunta "eles se encontram?" em "existe uma solução inteira e não-negativa para a equação de encontro?".
- **Eficiência**: A abordagem matemática com complexidade O(1) é muito mais eficiente do que uma simulação de pulos, que poderia falhar em casos com grandes números ou não terminar se o canguru de trás for mais lento.
- **Operador Módulo**: O uso do operador de módulo (`%`) provou ser a forma mais direta e limpa de verificar a divisibilidade e garantir que o número de pulos é um inteiro.
