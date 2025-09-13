# Dia 30 - Designer PDF Viewer

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: Easy  
**Link**: https://www.hackerrank.com/challenges/designer-pdf-viewer/problem

### Enunciado

Dado um array `h` de 26 inteiros representando as alturas das letras `a` a `z` e uma palavra `word`, determine a área do retângulo destacado ao selecionar a palavra em um visualizador de PDF. Cada letra tem largura 1, e a altura do retângulo é a altura da letra mais alta presente em `word`.

Ex.: se as alturas de `a..z` são dadas e `word = "abc"`, o retângulo terá altura igual à maior altura entre `a`, `b` e `c`, e largura igual a `len(word)`.

## Solução

### Abordagem

- Mapear cada caractere `c` da palavra para seu índice `c - 'a'` e obter `h[idx]`.
- Manter o máximo das alturas encontradas.
- A área é `maxAltura * len(word)`.

### Complexidade

- Tempo: O(n), onde n = tamanho de `word`.
- Espaço: O(1).

## Implementação

- Código: `days/day30-designer-pdf-viewer/main.go`

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Reforço do mapeamento `byte` -> índice (`c - 'a'`) para letras minúsculas ASCII.
- Lembrar que `len(string)` em Go mede bytes, o que é suficiente aqui pois o domínio é `a-z`.

## Screenshots

Adicione aqui o print do "Accepted" quando disponível.
