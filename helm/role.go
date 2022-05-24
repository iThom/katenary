package helm

type Rule struct {
	ApiGroup  []string `yaml:"apiGroup,omitempty"`
	Resources []string `yaml:"resource,omitempty"`
	Verbs     []string `yaml:"verbs,omitempty"`
}

type Role struct {
	*K8sBase `yaml:",inline"`
	Rules    []Rule `yaml:"rules,omitempty"`
}

func NewCronRole(name string) *Role {
	role := &Role{
		K8sBase: NewBase(),
	}

	role.K8sBase.Metadata.Name = ReleaseNameTpl + "-" + name + "-cron-executor"
	role.K8sBase.Kind = "Role"
	role.K8sBase.ApiVersion = "rbac.authorization.k8s.io/v1"
	role.K8sBase.Metadata.Labels[K+"/component"] = name

	role.Rules = []Rule{
		{
			ApiGroup:  []string{""},
			Resources: []string{"pods", "pods/exec"},
			Verbs:     []string{"get", "list", "create"},
		},
	}

	return role
}
