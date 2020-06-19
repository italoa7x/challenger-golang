package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/italoa7x/simple-project-golang/domain"
)

func testNewUser(t *testing.T) {

	_, err := domain.User.NewStructUser()
	require.Nil(t, err)
}
