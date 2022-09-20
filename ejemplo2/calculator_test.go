package main_test

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Error("La suma no es correcta")
		t.Fail()
	} else {
		t.Log("Finalizado correctamente")
	}
}
