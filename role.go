package govoi

import (
	"errors"

	"github.com/mikespook/gorbac"
)

const (
	RoleEndUser = "enduser"
	RoleHunter  = "hunder"
	RoleAdmin   = "admin"
)

var ErrPermissionDenied = errors.New("Permission denied")

type RoleManager struct {
	rbac *gorbac.RBAC
}

func NewRoleManager() *RoleManager {
	rbac := gorbac.New()

	rEndUser := gorbac.NewStdRole(RoleEndUser)
	rHunter := gorbac.NewStdRole(RoleHunter)
	rAdmin := gorbac.NewStdRole(RoleAdmin)

	pLock := gorbac.NewStdPermission(EventLock)
	pUnlock := gorbac.NewStdPermission(EventUnlock)
	pCollect := gorbac.NewStdPermission(EventCollect)
	pDrop := gorbac.NewStdPermission(EventDrop)
	pServiceMode := gorbac.NewStdPermission(EventServiceMode)
	pUnkonwCheck := gorbac.NewStdPermission(EventUnknowCheck)
	pBetteryCheck := gorbac.NewStdPermission(EventBetteryCheck)

	// EndUser
	rEndUser.Assign(pLock)
	rEndUser.Assign(pUnlock)
	rbac.Add(rEndUser)

	// Hunter
	rHunter.Assign(pLock)
	rHunter.Assign(pUnlock)
	rHunter.Assign(pCollect)
	rHunter.Assign(pDrop)
	rbac.Add(rHunter)

	// Admin
	rAdmin.Assign(pLock)
	rAdmin.Assign(pUnlock)
	rAdmin.Assign(pCollect)
	rAdmin.Assign(pDrop)
	rAdmin.Assign(pServiceMode)
	rAdmin.Assign(pUnkonwCheck)
	rAdmin.Assign(pBetteryCheck)
	rbac.Add(rAdmin)

	return &RoleManager{rbac}
}

// IsAllow tests if the role has permission
func (r *RoleManager) IsAllow(role string, event string) (bool, error) {
	permisison := gorbac.NewStdPermission(event)
	if r.rbac.IsGranted(role, permisison, nil) {
		return true, nil
	} else {
		return false, ErrPermissionDenied
	}
}

// IsDeny tests if the role doen't have permission
func (r *RoleManager) IsDeny(role string, event string) (bool, error) {
	rslt, err := r.IsAllow(role, event)
	return !rslt, err
}
