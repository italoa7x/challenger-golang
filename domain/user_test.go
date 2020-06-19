package domain_test

import (
	"github.com/italoa7x/simple-project-golang/domain"
	"testing"


	"github.com/stretchr/testify/require"
)

func testNewUser(t *testing.T) {

	_, err := domain.NewStructUser("Italo", "italo@gmail.com", "123")
	require.Nil(t, err)
	// testa se a validacao do email estar correta
	_, err2 := domain.NewStructUser("Maria", "maria@gmail.com", "123")
	require.Error(t, err2)
}
