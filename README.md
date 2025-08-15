# 100 Dias de HackerRank em Go

📅 **Duração**: 100 dias consecutivos (sem faltar)  
🎯 **Meta mínima**: 75 problemas resolvidos (HackerRank ou LeetCode)  
💻 **Linguagem**: Go (não posso usar linguagens que já domino)  
🤖 **IA**: proibido para resolver o problema; permitido para sintaxe/pesquisa  
🔍 **Pesquisa**: posso usar docs oficiais e buscar conceitos, mas não a solução pronta  
⏰ **Falhou um dia?** Volta pro dia 0 (streak zera)  
✅ **Não consegui resolver no dia?** Conta como dia válido, mas não como problema resolvido  

## Progresso

- **Dias cumpridos**: 0/100
- **Problemas resolvidos**: 0/75 (mínimo)
- **Streak atual**: 0

[📊 Ver progresso detalhado em TRACK.md](./TRACK.md)

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
./scripts/new-day.sh 1 "two-sum"
```

Isso criará a estrutura em `days/day01-two-sum/` com os arquivos base.

### Documentar solução

1. Resolver o problema em `days/dayXX-problema/main.go`
2. Preencher o `README.md` do dia com:
   - Enunciado resumido
   - Abordagem utilizada
   - Complexidade (tempo/espaço)
   - Dificuldades encontradas
   - O que aprendi
3. Atualizar o `TRACK.md` com o resultado

## Links Úteis

- [HackerRank](https://www.hackerrank.com/)
- [LeetCode](https://leetcode.com/)
- [Documentação oficial do Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)