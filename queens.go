package main

import (
	"fmt"
	"sort"
)

const (
	Empty cuadro = iota
	Queen
	Occupied
)

func run(tablero *Tablero) *Tablero {
	if tablero.reinas() == tablero.n {
		return tablero
	}
	for _, coordinate := range tablero.cuadrosDisponibles() {
		tablero.colocarReina(coordinate)
		next := run(tablero)
		if next != nil {
			return next
		}
		tablero.quitarReina(coordinate)
	}
	return nil
}

type cuadro int

func (f cuadro) String() string {
	switch f {
	case Queen:
		return "Q"
	case Occupied:
		return "X"
	default:
		return "?"
	}
}

type Tablero struct {
	Fields []cuadro
	n      int
}

func nuevoTablero(n int) *Tablero {
	if n < 4 {
		panic("Failed")
	}
	return &Tablero{
		Fields: make([]cuadro, n*n),
		n: n,
	}
}

func (b Tablero) String() string {
	sb := ""
	for i, v := range b.Fields {
		if i%b.n == 0 {
			sb += "\n"
		}
		sb += fmt.Sprintf("%v", v)
	}
	return sb
}

func (b *Tablero) colocarReina(n int) {
	b.Fields[n] = Queen
	for _, v := range b.cuadrosOcupados(n) {
		if v == n {
			continue
		}
		b.Fields[v] = Occupied
	}
}

func (b *Tablero) quitarReina(n int) {
	b.Fields[n] = Empty
	for _, v := range b.cuadrosOcupados(n) {
		b.Fields[v] = Empty
	}
}

func (b Tablero) reinas() int {
	count := 0
	for _, v := range b.Fields {
		if v == Queen {
			count++
		}
	}
	return count
}

func (b Tablero) cuadrosDisponibles() []int {
	res := make([]int, 0)
	for i, v := range b.Fields {
		if v == Empty {
			res = append(res, i)
		}
	}
	return res
}

func (b Tablero) cuadrosOcupados(n int) []int {
	line := n / b.n
	column := n % b.n
	res := make([]int, 0)
	res = append(res, b.getNumerosLinea(line)...)
	res = append(res, b.getNumerosColumna(column)...)
	res = append(res, b.getDiagonales(line, column)...)
	return sortQuitarDuplicados(res)
}

func (b Tablero) getNumerosLinea(line int) []int {
	res := make([]int, 0)
	for i := line * b.n; i < line*b.n+b.n; i++ {
		res = append(res, i)
	}
	return sortQuitarDuplicados(res)
}

func (b Tablero) getNumerosColumna(col int) []int {
	res := make([]int, 0)
	for i := col; i < b.n*b.n; i += b.n {
		res = append(res, i)
	}
	return sortQuitarDuplicados(res)
}

func (b Tablero) getDiagonales(line, col int) []int {
	res := make([]int, 0)
	res = append(res, getArribaDerecha(line, col, b.n)...)
	res = append(res, getArribaIzquierda(line, col, b.n)...)
	res = append(res, getFondoDerecha(line, col, b.n)...)
	res = append(res, getFondoIzquierda(line, col, b.n)...)
	return sortQuitarDuplicados(res)
}

func getArribaDerecha(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line-1, col-1
	for {
		if i < 0 || j < 0 {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i--
		j--
	}
	return res
}

func getArribaIzquierda(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line-1, col+1
	for {
		if i < 0 || j >= limit {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i--
		j++
	}
	return res
}

func getFondoDerecha(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line+1, col-1
	for {
		if i >= limit || j < 0 {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i++
		j--
	}
	return res
}

func getFondoIzquierda(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line+1, col+1
	for {
		if i >= limit || j >= limit {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i++
		j++
	}
	return res
}

func sortQuitarDuplicados(nums []int) []int {
	seen := make(map[int]bool, 0)
	res := make([]int, 0)

	for v := range nums {
		if seen[nums[v]] == false {
			seen[nums[v]] = true
			res = append(res, nums[v])
		}
	}
	sort.Ints(res)
	return res
}

func main() {
	board := nuevoTablero(6)
	run(board)
	fmt.Println(board)
}