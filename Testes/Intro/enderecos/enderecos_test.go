package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel();

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua João Limoeiro", "Rua"},
		{"Estrada dos Pintos", "Estrada"},
		{"Avenida Caxangá", "Avenida"},
		{"Praça Faria Neves", "Tipo Inválido!"},
		{"Rodovia Federal", "Rodovia"},
		{"RUA PROFESSOR CLAUDIO SELVA", "Rua"},
		{"", "Tipo Inválido!"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do tipo esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestQualquer(t *testing.T) {
	
	t.Parallel();

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
