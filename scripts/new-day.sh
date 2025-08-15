#!/bin/bash

# Script para criar estrutura de um novo dia do desafio
# Uso: ./scripts/new-day.sh <numero_do_dia> <nome_do_problema>
# Exemplo: ./scripts/new-day.sh 1 "two-sum"

set -e

if [ $# -ne 2 ]; then
    echo "Uso: $0 <numero_do_dia> <nome_do_problema>"
    echo "Exemplo: $0 1 \"two-sum\""
    exit 1
fi

DAY_NUM=$1
PROBLEM_NAME=$2

# Formatar número do dia com zero à esquerda
DAY_FORMATTED=$(printf "%02d" $DAY_NUM)

# Nome da pasta
FOLDER_NAME="day${DAY_FORMATTED}-${PROBLEM_NAME}"
FOLDER_PATH="days/${FOLDER_NAME}"

# Verificar se a pasta já existe
if [ -d "$FOLDER_PATH" ]; then
    echo "❌ Erro: A pasta '$FOLDER_PATH' já existe!"
    exit 1
fi

# Criar diretório
mkdir -p "$FOLDER_PATH"

# Criar main.go
cat > "$FOLDER_PATH/main.go" << 'EOF'
package main

import "fmt"

func main() {
	fmt.Println("Hello, Day X!")
	
	// TODO: Implementar solução aqui
}

// TODO: Implementar função(ões) necessária(s)
EOF

# Criar README.md do dia
cat > "$FOLDER_PATH/README.md" << EOF
# Dia ${DAY_NUM} - ${PROBLEM_NAME}

## Problema

**Plataforma**: [HackerRank/LeetCode]  
**Dificuldade**: [Easy/Medium/Hard]  
**Link**: [URL do problema]

### Enunciado

[Descrever brevemente o problema aqui]

## Solução

### Abordagem

[Explicar a estratégia utilizada]

### Complexidade

- **Tempo**: O(?)
- **Espaço**: O(?)

## Implementação

```go
// Código principal em main.go
```

## Resultado

- [x] Problema resolvido
- [ ] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- [O que aprendi hoje]
- [Dificuldades encontradas]
- [Conceitos novos]

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
EOF

echo "✅ Estrutura criada com sucesso em: $FOLDER_PATH"
echo ""
echo "📁 Arquivos criados:"
echo "   - $FOLDER_PATH/main.go"
echo "   - $FOLDER_PATH/README.md"
echo ""
echo "🚀 Próximos passos:"
echo "   1. Ir para: cd $FOLDER_PATH"
echo "   2. Editar main.go com sua solução"
echo "   3. Testar: go run main.go"
echo "   4. Documentar no README.md"
echo "   5. Atualizar TRACK.md"