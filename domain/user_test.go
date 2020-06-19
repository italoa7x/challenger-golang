package domain_test

import (
	"github.com/italoa7x/simple-project-golang/domain"
	"testing"


	"github.com/stretchr/testify/require"
)

func testNewUser(t *testing.T) {

	_, err := domain.NewUser("Italo", "italo@gmail.com", "123")
	require.Nil(t, err)
}
