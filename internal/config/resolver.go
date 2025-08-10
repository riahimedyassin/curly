package config

import (
	"fmt"
	"slices"
)

// Resolve creates the global config by applying team config dominance
func (r *ConfigLoader) Resolve() (*Config, error) {
	resolvedVars, err := r.resolveVariables()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve variables: %w", err)
	}

	return &Config{
		Team:         r.primaryConfig.Team.Team,
		Organization: r.primaryConfig.Team.Organization,
		Template:     r.primaryConfig.Template.TemplateInfo,
		Variables:    *resolvedVars,
		Files:        r.primaryConfig.Template.Files,
	}, nil
}

func (r *ConfigLoader) resolveVariables() (*resolvedVariables, error) {
	vars := &resolvedVariables{}

	vars.ComponentName = r.resolveStringVariable(
		r.primaryConfig.Template.Variables.ComponentName,
		r.primaryConfig.Team.Defaults.ComponentPath,
		r.primaryConfig.Team.Enforcement.ComponentPath != "",
	)

	vars.ComponentPath = r.resolveStringVariable(
		r.primaryConfig.Template.Variables.ComponentPath,
		r.primaryConfig.Team.Defaults.ComponentPath,
		r.primaryConfig.Team.Enforcement.ComponentPath != "",
	)

	vars.IncludeTests = r.resolveBoolVariable(
		r.primaryConfig.Template.Variables.IncludeTests,
		r.primaryConfig.Team.Defaults.IncludeTests,
		r.primaryConfig.Team.Enforcement.IncludeTests,
	)

	vars.IncludeProps = r.resolveBoolVariable(
		r.primaryConfig.Template.Variables.IncludeProps,
		r.primaryConfig.Team.Defaults.IncludeProps,
		r.primaryConfig.Team.Enforcement.IncludeProps,
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

func (r *ConfigLoader) resolveStringVariable(templateVar Variable[string], teamDefault string, enforced bool) resolvedVariable[string] {
	value := templateVar.Default
	if teamDefault != "" {
		value = teamDefault
	}

	return resolvedVariable[string]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    enforced,
	}
}

func (r *ConfigLoader) resolveBoolVariable(templateVar Variable[bool], teamDefault bool, enforced bool) resolvedVariable[bool] {
	value := templateVar.Default
	if enforced {
		value = teamDefault
	}

	return resolvedVariable[bool]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    enforced,
	}
}

func (r *ConfigLoader) resolveStylingVariable() (resolvedVariable[string], error) {
	templateVar := r.primaryConfig.Template.Variables.Styling
	value := templateVar.Default

	// Apply team default if set
	if r.primaryConfig.Team.Defaults.Styling != "" {
		value = r.primaryConfig.Team.Defaults.Styling
	}

	// Validate against restrictions
	if err := r.validateStyling(value); err != nil {
		return resolvedVariable[string]{}, err
	}

	return resolvedVariable[string]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    r.primaryConfig.Team.Defaults.Styling != "",
	}, nil
}

func (r *ConfigLoader) resolveExportTypeVariable() (resolvedVariable[string], error) {
	templateVar := r.primaryConfig.Template.Variables.ExportType
	value := templateVar.Default

	// Apply team default if set
	if r.primaryConfig.Team.Defaults.ExportType != "" {
		value = r.primaryConfig.Team.Defaults.ExportType
	}

	// Validate against restrictions
	if err := r.validateExportType(value); err != nil {
		return resolvedVariable[string]{}, err
	}

	return resolvedVariable[string]{
		Value:       value,
		Required:    templateVar.Required,
		Description: templateVar.Description,
		Enforced:    r.primaryConfig.Team.Defaults.ExportType != "",
	}, nil
}

// Validation methods
func (r *ConfigLoader) validateStyling(value string) error {
	restrictions := r.primaryConfig.Team.Restrictions.Styling

	// Check forbidden list
	if slices.Contains(restrictions.Forbidden, value) {
		return fmt.Errorf("styling '%s' is forbidden by team primaryConfig", value)
	}

	// Check allowed list (if specified)
	if len(restrictions.Allowed) > 0 && !slices.Contains(restrictions.Allowed, value) {
		return fmt.Errorf("styling '%s' is not in allowed list: %v", value, restrictions.Allowed)
	}

	return nil
}

func (r *ConfigLoader) validateExportType(value string) error {
	restrictions := r.primaryConfig.Team.Restrictions.ExportType

	// Check forbidden list
	if slices.Contains(restrictions.Forbidden, value) {
		return fmt.Errorf("export type '%s' is forbidden by team primaryConfig", value)
	}

	// Check allowed list (if specified)
	if len(restrictions.Allowed) > 0 && !slices.Contains(restrictions.Allowed, value) {
		return fmt.Errorf("export type '%s' is not in allowed list: %v", value, restrictions.Allowed)
	}

	return nil
}
