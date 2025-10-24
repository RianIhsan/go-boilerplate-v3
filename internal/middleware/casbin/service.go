package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"strconv"
)

type Service struct {
	enforcer *casbin.Enforcer
}

func NewService(enforcer *casbin.Enforcer) *Service {
	return &Service{enforcer: enforcer}
}

func (s *Service) RegisterUser(user UserCasbin) error {
	id := strconv.FormatUint(uint64(user.ID), 10)
	role := strconv.FormatUint(uint64(user.RoleId), 10)

	_, err := s.enforcer.AddRoleForUser(id, role)
	if err != nil {
		return err
	}

	// default role assignment
	//_, err = s.enforcer.AddRoleForUser(id, "1")
	//if err != nil {
	//	return err
	//}

	return s.enforcer.SavePolicy()
}

func (s *Service) UpdateRole(user UserCasbin) error {
	idStr := strconv.FormatUint(uint64(user.ID), 10)
	lastRoleStr := strconv.FormatUint(uint64(user.LastRole), 10)
	newRoleStr := strconv.FormatUint(uint64(user.RoleId), 10)

	_, err := s.enforcer.DeleteRoleForUser(idStr, lastRoleStr)
	if err != nil {
		return err
	}

	_, err = s.enforcer.AddRoleForUser(idStr, newRoleStr)
	if err != nil {
		return err
	}

	return s.enforcer.SavePolicy()
}

func (s *Service) DeleteRole(user UserCasbin) error {
	idStr := strconv.FormatUint(uint64(user.ID), 10)
	roleStr := strconv.FormatUint(uint64(user.LastRole), 10)

	_, err := s.enforcer.DeleteRoleForUser(idStr, roleStr)
	if err != nil {
		return err
	}

	return s.enforcer.SavePolicy()
}

func (s *Service) AddRole(user UserCasbin) error {
	idStr := strconv.FormatUint(uint64(user.ID), 10)
	roleStr := strconv.FormatUint(uint64(user.RoleId), 10)

	_, err := s.enforcer.AddRoleForUser(idStr, roleStr)
	if err != nil {
		return err
	}

	return s.enforcer.SavePolicy()
}

func (s *Service) AddPolicy(role RolePolicy) error {
	roleIDStr := strconv.FormatUint(uint64(role.ID), 10)
	var policies [][]string

	for _, v := range role.Policy {
		policies = append(policies, []string{roleIDStr, v.Path, v.Method})
	}

	_, err := s.enforcer.AddPolicies(policies)
	if err != nil {
		return err
	}

	return s.enforcer.SavePolicy()
}

func (s *Service) RemovePolicy(role RolePolicy) error {
	roleIDStr := strconv.FormatUint(uint64(role.ID), 10)
	var policies [][]string

	for _, v := range role.LastPolicy {
		policies = append(policies, []string{roleIDStr, v.Path, v.Method})
	}

	_, err := s.enforcer.RemovePolicies(policies)
	if err != nil {
		return err
	}

	return s.enforcer.SavePolicy()
}

func (s *Service) UpdatePolicy(role RolePolicy) error {
	roleIDStr := strconv.FormatUint(uint64(role.ID), 10)
	var newPolicies, oldPolicies [][]string

	for _, p := range role.Policy {
		newPolicies = append(newPolicies, []string{roleIDStr, p.Path, p.Method})
	}
	for _, p := range role.LastPolicy {
		oldPolicies = append(oldPolicies, []string{roleIDStr, p.Path, p.Method})
	}

	_, err := s.enforcer.UpdatePolicies(oldPolicies, newPolicies)
	if err != nil {
		return err
	}

	return s.enforcer.SavePolicy()
}

func (s *Service) UpdatePathForPolicies(oldPath, newPath, method string) error {
	policies, err := s.enforcer.GetPolicy()
	if err != nil {
		return fmt.Errorf("failed to get policies: %w", err)
	}
	var oldPolicies, newPolicies [][]string

	for _, p := range policies {
		if len(p) >= 3 && p[1] == oldPath && p[2] == method {
			oldPolicies = append(oldPolicies, p)
			newPolicies = append(newPolicies, []string{p[0], newPath, method})
		}
	}

	if len(oldPolicies) == 0 {
		return fmt.Errorf("no matching policy for path: %s and method: %s", oldPath, method)
	}

	_, err = s.enforcer.UpdatePolicies(oldPolicies, newPolicies)
	if err != nil {
		return err
	}

	return s.enforcer.SavePolicy()
}
