# Dia 37 - jumping-on-the-clouds-revisited

## Problema

**Plataforma**: HackerRank  
**Dificuldade**: 🟢 Easy  
**Link**: https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem

### Enunciado

A child is playing a cloud hopping game where they jump from cloud to cloud in a circular path until reaching the start again. Each jump costs 1 energy, and landing on a thundercloud costs an additional 2 energy.

## Solução

### Abordagem

Simulate the game by jumping k positions at a time using modulo arithmetic for circular movement. Track energy loss for each jump and additional loss for thunderclouds.

### Complexidade

- **Tempo**: O(n) - worst case visits all clouds once
- **Espaço**: O(1) - only uses constant extra space

## Implementação

```go
func jumpingOnClouds(c []int32, k int32) int32 {
	n := len(c)
	energy := int32(100)
	pos := 0

	for {
		pos = (pos + int(k)) % n
		energy--
		if c[pos] == 1 {
			energy -= 2
		}
		if pos == 0 {
			break
		}
	}

	return energy
}
```

## Resultado

- [x] Problema resolvido
- [x] Accepted no primeiro submit
- [ ] Precisou de múltiplas tentativas

## Aprendizados

- Uso de aritmética modular para movimento circular
- Simulação simples com controle de energia
- Loop até retornar à posição inicial

## Screenshots

[Adicionar print do resultado "Accepted" ou das tentativas]
