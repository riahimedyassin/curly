package config

import (
	"fmt"
	"slices"
)

type ConfigResolver struct {
	teamConfig     *TeamConfig
	templateConfig *TemplateConfig
	templateInfo   *TemplateInfo
}

func NewConfigResolver(team *TeamConfig, template *TemplateConfig, info *TemplateInfo) *ConfigResolver {
	return &ConfigResolver{
		teamConfig:     team,
		templateConfig: template,
		templateInfo:   info,
	}
}

// Resolve creates the global config by applying team config dominance
func (r *ConfigResolver) Resolve() (*GlobalConfig, error) {
	resolvedVars, err := r.resolveVariables()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve variables: %w", err)
	}

	return &GlobalConfig{
		Team:         r.teamConfig.Team,
		Organization: r.teamConfig.Organization,
		Template:     *r.templateInfo,
		Variables:    *resolvedVars,
		Files:        r.templateConfig.Files,
	}, nil
}

func (r *ConfigResolver) resolveVariables() (*ResolvedVariables, error) {
	vars := &ResolvedVariables{}

	vars.ComponentName = r.resolveStringVariable(
		r.templateConfig.Variables.ComponentName,
		r.teamConfig.Defaults.ComponentPath,
		r.teamConfig.Enforcement.ComponentPath != "",
	)

	vars.ComponentPath = r.resolveStringVariable(
		r.templateConfig.Variables.ComponentPath,
		r.teamConfig.Defaults.ComponentPath,
		r.teamConfig.Enforcement.ComponentPath != "",
	)

	vars.IncludeTests = r.resolveBoolVariable(
		r.templateConfig.Variables.IncludeTests,
		r.teamConfig.Defaults.IncludeTests,
		r.teamConfig.Enforcement.IncludeTests,
	)

	vars.IncludeProps = r.resolveBoolVariable(
		r.templateConfig.Variables.IncludeProps,
		r.teamConfig.Defaults.IncludeProps,
		r.teamConfig.Enforcement.IncludeProps,
	)

	var err error
	vars.Styling, err = r.resolveStylingVariable()
	if err != nil {
		return nil, err
	}

	vars.ExportType, err = r.resolveExportTypeVariable()
	if err != nil {
		return nil, err
	}

	return vars, nil
}

func (r *ConfigResolver) resolveStringVariable(templateVar Variable[string], teamDefault string, enforced bool) ResolvedVariable[string] {
	value := templateVar.Default
	if teamDefault != "" {
		value = teamDefault
	}

	return ResolvedVariable[string]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    enforced,
	}
}

func (r *ConfigResolver) resolveBoolVariable(templateVar Variable[bool], teamDefault bool, enforced bool) ResolvedVariable[bool] {
	value := templateVar.Default
	if enforced {
		value = teamDefault
	}

	return ResolvedVariable[bool]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    enforced,
	}
}

func (r *ConfigResolver) resolveStylingVariable() (ResolvedVariable[string], error) {
	templateVar := r.templateConfig.Variables.Styling
	value := templateVar.Default

	// Apply team default if set
	if r.teamConfig.Defaults.Styling != "" {
		value = r.teamConfig.Defaults.Styling
	}

	// Validate against restrictions
	if err := r.validateStyling(value); err != nil {
		return ResolvedVariable[string]{}, err
	}

	return ResolvedVariable[string]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    r.teamConfig.Defaults.Styling != "",
	}, nil
}

func (r *ConfigResolver) resolveExportTypeVariable() (ResolvedVariable[string], error) {
	templateVar := r.templateConfig.Variables.ExportType
	value := templateVar.Default

	// Apply team default if set
	if r.teamConfig.Defaults.ExportType != "" {
		value = r.teamConfig.Defaults.ExportType
	}

	// Validate against restrictions
	if err := r.validateExportType(value); err != nil {
		return ResolvedVariable[string]{}, err
	}

	return ResolvedVariable[string]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    r.teamConfig.Defaults.ExportType != "",
	}, nil
}

// Validation methods
func (r *ConfigResolver) validateStyling(value string) error {
	restrictions := r.teamConfig.Restrictions.Styling

	// Check forbidden list
	if slices.Contains(restrictions.Forbidden, value) {
		return fmt.Errorf("styling '%s' is forbidden by team config", value)
	}

	// Check allowed list (if specified)
	if len(restrictions.Allowed) > 0 && !slices.Contains(restrictions.Allowed, value) {
		return fmt.Errorf("styling '%s' is not in allowed list: %v", value, restrictions.Allowed)
	}

	return nil
}

func (r *ConfigResolver) validateExportType(value string) error {
	restrictions := r.teamConfig.Restrictions.ExportType

	// Check forbidden list
	if slices.Contains(restrictions.Forbidden, value) {
		return fmt.Errorf("export type '%s' is forbidden by team config", value)
	}

	// Check allowed list (if specified)
	if len(restrictions.Allowed) > 0 && !slices.Contains(restrictions.Allowed, value) {
		return fmt.Errorf("export type '%s' is not in allowed list: %v", value, restrictions.Allowed)
	}

	return nil
}
