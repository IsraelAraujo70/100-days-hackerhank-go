# Dia 5 - plus-minus

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: [HackerRank - Plus Minus](https://www.hackerrank.com/challenges/plus-minus/problem)

### Enunciado

Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with 6 places after the decimal.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with small absolute error are acceptable.

Example

There are 5 elements: two positive, two negative and one zero. Their ratios are 0.4, 0.4 and 0.2. Results are printed as:

0.400000  
0.400000  
0.200000

Function Description

Complete the plusMinus function with the following parameter(s):

arr: an array of integers

Print

Print the ratios of positive, negative and zero values in the array. Each value should be printed on a separate line with 6 digits after the decimal. The function should not return a value.

Input Format

The first line contains an integer, n, the size of the array.  
The second line contains n space-separated integers that describe arr.

Sample Input

STDIN           Function  
-----           --------  
6               arr[] size n = 6  
-4 3 -9 0 4 1   arr = [-4, 3, -9, 0, 4, 1]

Sample Output

0.500000  
0.333333  
0.166667

Explanation

There are 3 positive numbers, 2 negative numbers, and 1 zero in the array.  
The proportions of occurrence are positive: 3/6, negative: 2/6 and zeros: 1/6.

## Solução

### Abordagem

- **Contagem**: percorrer o array uma única vez e contar quantos elementos são positivos, negativos e zeros.  
- **Cálculo**: dividir cada contagem por `n = len(arr)` para obter as razões.  
- **Saída**: imprimir cada razão com `fmt.Printf("%.6f\n", valor)`.

### Complexidade

- **Tempo**: O(n)
- **Espaço**: O(1)

## Implementação

```go
func plusMinus(arr []int32) {
    var positiveCounter, neagativeCounter, zeroCounter int
    len := len(arr)
    for i := range arr {
        if arr[i] > 0{
            positiveCounter ++
        }else if arr[i] < 0{
            neagativeCounter++
        }else{
            zeroCounter++
        }
    }
    if len > 0{
        plusRatio := float32(positiveCounter)/float32(len)
        minusRatio := float32(neagativeCounter)/float32(len)
        zeroRatios := float32(zeroCounter)/float32(len)        
        fmt.Printf("%.6f\n",plusRatio)
        fmt.Printf("%.6f\n",minusRatio)
        fmt.Printf("%.6f\n",zeroRatios)
    }   
}
```

## Resultado

- [x] Problema resolvido
- [ ] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Formatação de números de ponto flutuante com 6 casas usando `fmt.Printf("%.6f")`.
- Contagem de categorias em uma única passada do array (positivos, negativos e zeros).

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
