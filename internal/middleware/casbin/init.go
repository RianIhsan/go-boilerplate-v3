package casbin

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func InitCasbin(modelPath, policyPath string) *casbin.Enforcer {
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatalf("failed to initialize casbin enforcer: %v", err)
	}
	return enforcer
}
