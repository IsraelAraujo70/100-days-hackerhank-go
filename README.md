# 100 Dias de HackerRank em Go

📅 **Duração**: 100 dias consecutivos (sem faltar)
🎯 **Meta mínima**: 75 problemas resolvidos (HackerRank ou LeetCode)
💻 **Linguagem**: Go (não posso usar linguagens que já domino)
🤖 **IA**: proibido para resolver o problema; permitido para sintaxe/pesquisa
🔍 **Pesquisa**: posso usar docs oficiais e buscar conceitos, mas não a solução pronta
⏰ **Falhou um dia?** Volta pro dia 0 (streak zera)
✅ **Não consegui resolver no dia?** Conta como dia válido, mas não como problema resolvido

## Progresso

- **Dias cumpridos**: 43/100
- **Problemas resolvidos**: 43/75 (mínimo)
- **Streak atual**: 43
- **Última atualização**: 2025-09-26

[📊 Ver progresso detalhado em TRACK.md](./TRACK.md)

## Últimos Problemas Resolvidos

| Dia | Data | Problema | Plataforma | Dificuldade | Notas |
|-----|------|----------|------------|-------------|-------|
| 43 | 2025-09-26 | Library Fine | HackerRank | 🟢 Easy | Lógica condicional hierárquica para cálculo de multas |
| 42 | 2025-09-25 | Sherlock and Squares | HackerRank | 🟢 Easy | Matemática O(1) com raiz quadrada |
| 41 | 2025-09-24 | Append and Delete | HackerRank | 🟢 Easy | Transformação de strings com operações limitadas |
| 40 | 2025-09-23 | Extra Long Factorials | HackerRank | 🟡 Medium | Big integers com math/big |
| 39 | 2025-09-22 | Circular Array Rotation | HackerRank | 🟢 Easy | Aritmética modular para rotação |

## Estatísticas por Categoria

### 🟢 Easy: 36 problemas
### 🟡 Medium: 7 problemas
### 🔴 Hard: 0 problemas

### Por Tópico
- **Implementação**: 15 problemas
- **Strings**: 8 problemas
- **Arrays**: 7 problemas
- **Matemática**: 6 problemas
- **Algoritmos**: 4 problemas
- **Estruturas de Dados**: 3 problemas

## Prova Diária

Para cada dia, devo documentar:

1. Link do problema
2. Print do "Accepted"/resultado (ou do que tentei)
3. Link do repositório com o código
4. Breve relato do que aprendi

**Fuso horário**: America/Sao_Paulo (vale até 23:59 local)

## Regras Detalhadas

- **Linguagem**: Apenas Go
- **IA**: Proibido para resolver problemas. Permitido apenas para:
  - Dúvidas de sintaxe
  - Pesquisa de conceitos/documentação
- **Pesquisa**: Documentação oficial e conceitos são permitidos
- **Soluções prontas**: Proibido olhar antes de submeter
- **Consecutividade**: Falhar um dia = volta para o dia 0
- **Validação**: Não resolver ≠ falhar o dia (mas não conta como problema resolvido)

## Estrutura do Projeto

```
100-days-hackerrank-go/
├─ README.md
├─ TRACK.md
├─ days/
│  ├─ day01-problema-nome/
│  │  ├─ main.go
│  │  └─ README.md    # enunciado resumido, abordagem, complexidade
│  ├─ day02-...
│  └─ ...
└─ scripts/
   └─ new-day.sh      # script para criar pasta de novo dia
```

## Como Usar

### Criar um novo dia

```bash
./scripts/new-day.sh 44 "cut-the-sticks"
```

Isso criará a estrutura em `days/day44-cut-the-sticks/` com os arquivos base.

### Documentar solução

1. Resolver o problema em `days/dayXX-problema/main.go`
2. Preencher o `README.md` do dia com:
   - Enunciado resumido
   - Abordagem utilizada
   - Complexidade (tempo/espaço)
   - Dificuldades encontradas
   - O que aprendi
3. Atualizar o `TRACK.md` com o resultado

## Principais Aprendizados

### Conceitos de Go Aprendidos
- **Big integers**: Uso de `math/big` para números grandes
- **Slices**: Manipulação eficiente de arrays dinâmicos
- **Strings**: Operações de substring e transformação
- **Matemática**: Funções como `math.Sqrt`, `math.Abs`
- **I/O**: Leitura eficiente com `bufio.Scanner`

### Algoritmos e Estruturas de Dados
- **Aritmética modular**: Para rotações e ciclos
- **Busca binária**: Para otimização de buscas
- **Two pointers**: Para problemas de arrays
- **Greedy**: Para otimização local
- **Matemática**: Fórmulas para evitar bruteforce

### Padrões de Solução
- **Verificação de casos extremos**: Sempre considerar limites
- **Otimização O(1)**: Usar matemática quando possível
- **Reutilização de código**: Funções auxiliares
- **Leitura cuidadosa**: Entender bem o enunciado

## Links Úteis

- [HackerRank](https://www.hackerrank.com/)
- [LeetCode](https://leetcode.com/)
- [Documentação oficial do Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Playground](https://play.golang.org/)
- [Big O Cheat Sheet](https://www.bigocheatsheet.com/)

## Motivação

> "A jornada de mil milhas começa com um único passo" - Lao Tzu

Este desafio não é apenas sobre resolver problemas, mas sobre:
- **Disciplina**: Manter consistência por 100 dias
- **Aprendizado**: Dominar uma nova linguagem de programação
- **Crescimento**: Desenvolver pensamento algorítmico
- **Perseverança**: Continuar mesmo quando difícil

**Progresso atual**: 43% completo 🔥