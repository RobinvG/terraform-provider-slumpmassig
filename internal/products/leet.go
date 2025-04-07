// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package products

import (
	"math/rand"
	"strings"
)

func getLeetList() [][]string {
	A := []string{"A", "4"}
	B := []string{"B", "I3"}
	C := []string{"C", "["}
	D := []string{"D", ")", "|)"}
	E := []string{"E", "3"}
	F := []string{"F", "|="}
	G := []string{"G", "6", "&"}
	H := []string{"H", "#", "/-/"}
	I := []string{"I", "1", "!"}
	J := []string{"J", "_|", "_/", "¿", "</", "_]", "(/"}
	K := []string{"K", "|<", "|{", "]{", "|X"}
	L := []string{"L", "1", "£", "7", "|_"}
	O := []string{"O", "0", "()", "oh", "[]", "p", "<>"}
	P := []string{"P", "|*", "|o", "|^(o)", "|>"}
	S := []string{"S", "5", "$", "2"}
	T := []string{"T", "7", "+", "-|-"}
	U := []string{"U", "(_)", "|_|", "v", "L|", "µ", "บ"}
	W := []string{"W", "vv", "'//", "dubya", "(n)", "พ"}
	X := []string{"X", "><", "Ж", "}{", "ecks", "?", ")(", "][", "}{"}
	Y := []string{"Y", "j", "`/", "Ч", "7"}
	Z := []string{"Z", "2", "7_", "-/_"}

	return [][]string{A, B, C, D, E, F, G, H, I, J, K, L, O, P, S, T, U, W, X, Y, Z}
}

func getLeetIndex(c string, leetIndex []string) int {
	for p, v := range leetIndex {
		if c == v {
			return p
		}
	}

	return -1
}

func convertToLeet(input string) string {
	leets := getLeetList()

	var leetIndex []string
	for i := 0; i < len(leets); i++ {
		leetIndex = append(leetIndex, leets[i][0])
	}

	leetStr := ""

	for _, c := range input {
		upperC := strings.ToUpper(string([]rune{c}))

		n := getLeetIndex(upperC, leetIndex)
		if n > 0 {
			rnum := rand.Intn(len(leets[n]))
			leetStr = leetStr + leets[n][rnum]
		} else {
			leetStr = leetStr + string(c)
		}

	}

	return leetStr
}
