# Tracking - 100 Dias de HackerRank (Go)

**Meta**: 100 dias seguidos, 75+ problemas resolvidos  
**Sem IA** para resolver — apenas docs/pesquisa e perguntas de sintaxe

## Estatísticas

- **Dias cumpridos**: 13/100
- **Problemas resolvidos**: 13/75 (mínimo)
- **Streak atual**: 13
- **Última atualização**: 2025-08-27

## Progresso por Dia

| Dia | Data | Plataforma | Problema | Dificuldade | Código | Resultado | Tempo | Notas |
|-----|------|------------|----------|-------------|--------|-----------|--------|-------|
| 01  | 2025-08-15 | HR         | Simple Array Sum | 🟢          | days/day01-simple-array-sum/main.go | ✅        | -      | Aprendi a omitir o índice no `for range` (ex.: `for _, v := range ar`). Agradecimentos ao Hitesh Choudhary. |
| 02  | 2025-08-16 | HR         | Compare the Triplets | 🟢          | days/day02-compare-the-triplets/main.go | ✅        | -      | Implementei comparação de arrays usando range com índice. Algoritmo O(1) para comparar triplas e contar pontos. |
| 03  | 2025-08-17 | HR         | A Very Big Sum | 🟢          | days/day03-a-very-big-sum/main.go | ✅        | -      | Aprendi sobre overflow de int32 e a importância de usar int64 para números grandes. Padrão `for _, v := range` novamente aplicado. |
| 04  | 2025-08-18 | HR         | Diagonal Difference | 🟢          | days/day04-diagonal-difference/main.go | ✅        | -      | Aprendi sobre matrizes 2D em Go e acesso a diagonais. Me confundi com formato da entrada - achei que seria nested array mas era só tamanho + linhas. |
| 05  | 2025-08-19 | HR         | Plus Minus           | 🟢          | days/day05-plus-minus/main.go | ✅        | -      | Ratios de positivos/negativos/zeros; impressão com 6 casas decimais. 
| 06  | 2025-08-20 | HR         | Staircase | 🟢          | days/day06-staircase/main.go | ✅        | -      | Aprendi `strings.Repeat()` para gerar strings repetidas. Problema de formatação/impressão - função não retorna valor, apenas imprime. |
| 07  | 2025-08-21 | HR         | Mini-Max Sum | 🟢          | days/day07-mini-max-sum/main.go | ✅        | -      | Otimização O(n) vs O(n log n) - evitei sort e fiz em uma passagem. Sempre pensar em complexidade! |
| 08  | 2025-08-23 | HR         | Birthday Cake Candles | 🟢          | days/day08-birthday-cake-candles/main.go | ✅        | -      | Problema simples de busca por máximo e contagem. Abordagem de duas passadas pelo array é clara e eficiente. |
| 09  | 2025-08-24 | HR         | Grading Students | 🟢          | days/day09-grading/main.go | ✅        | -      | Lógica de arredondamento com módulo. Fórmula `-v%5 + 5` para calcular próximo múltiplo de 5. Condições múltiplas no if. |
| 10  | 2025-08-24 | HR         | CamelCase | 🟢          | days/day10-camelcase/main.go | ✅        | -      | Uso do pacote `unicode` para verificar caracteres maiúsculos. Conversão string para `[]rune` para Unicode. Contagem simples: cada maiúscula = nova palavra. |
| 11  | 2025-08-26 | HR         | Time Conversion | 🟢          | days/day11-time-conversion/main.go | ✅        | -      | Aprendi sobre layouts de time em Go. Uso de `time.Parse` e `time.Format` com layouts específicos. Go usa tempo de referência "Mon Jan 2 15:04:05 MST 2006". |
| 12  | 2025-08-27 | HR         | Counting Valleys | 🟢          | days/day12-counting-valleys/main.go | ✅        | -      | Algoritmo de tracking de estado para contar vales. Contagem precisa ao sair do vale (`altitude == 0 && v == 'U'`). Uso de `[]rune(path)` para Unicode. |
| 13  | 2025-08-27 | HR         | The Birthday Bar | 🟢          | days/day13-the-birthday-bar/main.go | ✅        | -      | Problema de sliding window. Solução O(n*m) com loops aninhados. Reforçou a importância dos limites de loop corretos. |

<!-- Continue até o dia 100... -->

## Legenda

**Resultado:**
- ✅ Problema resolvido (Accepted)
- ❌ Não consegui resolver
- ⏳ Ainda não tentado
- 🔄 Em progresso

**Dificuldade:**
- 🟢 Easy
- 🟡 Medium  
- 🔴 Hard

**Plataformas:**
- HR: HackerRank
- LC: LeetCode

## Regras (Lembrete)

1. **Streak**: Falhar um dia = volta para o dia 0
2. **IA**: Proibido para resolver (só sintaxe/pesquisa)
3. **Pesquisa**: Docs oficiais OK, solução pronta NÃO
4. **Tempo limite**: Até 23:59 America/Sao_Paulo
5. **Validação**: Não resolver ≠ falhar o dia

## Problemas por Categoria

### Arrays/Strings
- [x] Dia 1: Simple Array Sum
- [x] Dia 2: Compare the Triplets
- [x] Dia 3: A Very Big Sum
- [x] Dia 4: Diagonal Difference
- [x] Dia 5: Plus Minus
- [x] Dia 6: Staircase
- [x] Dia 7: Mini-Max Sum
- [x] Dia 8: Birthday Cake Candles
- [x] Dia 9: Grading Students
- [x] Dia 10: CamelCase
- [x] Dia 11: Time Conversion
- [x] Dia 12: Counting Valleys
- [x] Dia 13: The Birthday Bar

### Matemática
- [ ] Dia X: Problema Y

### Algoritmos
- [ ] Dia X: Problema Y

### Estruturas de Dados
- [ ] Dia X: Problema Y

---

*Atualizar este arquivo a cada dia concluído*