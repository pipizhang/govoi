package govoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleEndUser(t *testing.T) {
	var ok bool

	r := NewRoleManager()

	ok, _ = r.IsAllow(RoleEndUser, EventUnlock)
	assert.True(t, ok)

	ok, _ = r.IsDeny(RoleEndUser, EventUnlock)
	assert.False(t, ok)

	ok, _ = r.IsAllow(RoleEndUser, EventCollect)
	assert.False(t, ok)
}

func TestRoleHunter(t *testing.T) {
	var ok bool

	r := NewRoleManager()

	ok, _ = r.IsAllow(RoleHunter, EventUnlock)
	assert.True(t, ok)

	ok, _ = r.IsAllow(RoleEndUser, EventServiceMode)
	assert.False(t, ok)
}

func TestRoleAdmin(t *testing.T) {
	var ok bool

	r := NewRoleManager()

	ok, _ = r.IsAllow(RoleAdmin, EventDrop)
	assert.True(t, ok)

	ok, _ = r.IsDeny(RoleAdmin, EventServiceMode)
	assert.False(t, ok)
}

func BenchmarkIsAllow(b *testing.B) {
	r := NewRoleManager()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.IsAllow(RoleEndUser, EventUnlock)
	}
}
