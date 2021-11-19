package server

import (
	"testing"
)

func TestStartServer(t *testing.T) {

	// TODO optimization test login
	t.Run("start server", func(t *testing.T) {
		start()
	})

}
