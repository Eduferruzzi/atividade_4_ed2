package main

// Preprocessamento guarda o que Preprocessa calcula e Distancia consulta.
type Preprocessamento struct {
	// Slice que garda a distancia de v até um dos extremos da linha. O calculo é feito com o mesmo extremo para todos
	distExtremo []int
}

// Preprocessa recebe um grafo linha, com pelo menos dois vértices e pesos
// positivos, e prepara as consultas de distância. O custo deve ser O(V).
func Preprocessa(g *GrafoP) *Preprocessamento {
	extremo := 0
	dists := make([]int, g.V())
	for i := 0; i < g.V(); i++ {
		if g.Grau(i) == 1 {
			extremo = i
			break
		}
	}

	v := extremo
	distPercorrida := 0
	ant := -1
	for g.Grau(v) != 1 || v == extremo {
		dists[v] = distPercorrida
		adjs := g.Adj(v)
		for j := 0; j < 2; j++ {
			if adjs[j].Outro(v) != ant {
				ant = v
				v = adjs[j].Outro(v)
				distPercorrida += adjs[j].Peso
				break
			}
		}
		if g.Grau(v) == 1 {
			dists[v] = distPercorrida
		}
	}

	return &Preprocessamento{distExtremo: dists}
}

// Distancia devolve o comprimento do caminho mais curto entre os vértices v e
// w. O custo deve ser O(1).
func (p *Preprocessamento) Distancia(v, w int) int {
	// Subtrai as distâncias até o extremo dos 2 vértices
	return abs(p.distExtremo[v] - p.distExtremo[w])
}

func abs(valor int) int {
	if valor < 0 {
		return valor * (-1)
	}
	return valor
}
