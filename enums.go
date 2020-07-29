/*
Copyright 2020 The Flux CD contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gitprovider

import "github.com/fluxcd/go-git-providers/validation"

// TransportType is an enum specifying the transport type used when cloning a repository
type TransportType string

const (
	// TransportTypeHTTPS specifies a clone URL of the form:
	// https://<domain>/<org>/[<sub-orgs...>/]<repo>.git
	TransportTypeHTTPS = TransportType("https")
	// TransportTypeGit specifies a clone URL of the form:
	// git@<domain>:<org>/[<sub-orgs...>/]<repo>.git
	TransportTypeGit = TransportType("git")
	// TransportTypeSSH specifies a clone URL of the form:
	// ssh://git@<domain>/<org>/[<sub-orgs...>/]<repo>
	TransportTypeSSH = TransportType("ssh")
)

// RepositoryCredentialType is an enum specifying the type of a certain repository credential
type RepositoryCredentialType string

const (
	// RepositoryCredentialTypeDeployKey is the Deploy Key credential type
	// The struct used is DeployKey in this package
	RepositoryCredentialTypeDeployKey = RepositoryCredentialType("deploykey")
)

// RepoVisibility is an enum specifying the visibility of a repository
type RepoVisibility string

const (
	// RepoVisibilityPublic specifies that the repository should be publicly accessible
	RepoVisibilityPublic = RepoVisibility("public")
	// RepoVisibilityInternal specifies that the repository should accessible within the
	// own organization
	RepoVisibilityInternal = RepoVisibility("internal")
	// RepoVisibilityPrivate specifies that the repository should only be accessible by
	// specifically added team members
	RepoVisibilityPrivate = RepoVisibility("private")
)

// knownRepoVisibilityValues is a map of known RepoVisibility values, used for validation
var knownRepoVisibilityValues = map[RepoVisibility]struct{}{
	RepoVisibilityPublic:   {},
	RepoVisibilityInternal: {},
	RepoVisibilityPrivate:  {},
}

// validateRepoVisibility validates a given RepoVisibility.
// Use as errs.Append(validateRepoVisibility(visibility), visibility, "FieldName")
func validateRepoVisibility(r RepoVisibility) error {
	_, ok := knownRepoVisibilityValues[r]
	if !ok {
		return validation.ErrFieldEnumInvalid
	}
	return nil
}

// repoVisibilityVar returns a pointer to a RepoVisibility
func repoVisibilityVar(r RepoVisibility) *RepoVisibility {
	return &r
}

// RepositoryPermission is an enum specifying the access level for a certain team or person
// for a given repository
type RepositoryPermission string

const (
	// RepositoryPermissionPull ("pull") - team members can pull, but not push to or administer this repository
	// This is called "guest" in GitLab
	RepositoryPermissionPull = RepositoryPermission("pull")

	// RepositoryPermissionTriage ("triage") - team members can proactively manage issues and pull requests without write access.
	// This is called "reporter" in GitLab
	RepositoryPermissionTriage = RepositoryPermission("triage")

	// RepositoryPermissionPush ("push") - team members can pull and push, but not administer this repository
	// This is called "developer" in GitLab
	RepositoryPermissionPush = RepositoryPermission("push")

	// RepositoryPermissionMaintain ("maintain") - team members can manage the repository without access to sensitive or destructive actions.
	// This is called "maintainer" in GitLab
	RepositoryPermissionMaintain = RepositoryPermission("maintain")

	// RepositoryPermissionAdmin ("admin") - team members can pull, push and administer this repository
	// This is called "admin" or "owner" in GitLab
	RepositoryPermissionAdmin = RepositoryPermission("admin")
)

// knownRepoVisibilityValues is a map of known RepositoryPermission values, used for validation
var knownRepositoryPermissionValues = map[RepositoryPermission]struct{}{
	RepositoryPermissionPull:     {},
	RepositoryPermissionTriage:   {},
	RepositoryPermissionPush:     {},
	RepositoryPermissionMaintain: {},
	RepositoryPermissionAdmin:    {},
}

// validateRepositoryPermission validates a given RepositoryPermission.
// Use as errs.Append(validateRepositoryPermission(permission), permission, "FieldName")
func validateRepositoryPermission(p RepositoryPermission) error {
	_, ok := knownRepositoryPermissionValues[p]
	if !ok {
		return validation.ErrFieldEnumInvalid
	}
	return nil
}

// repositoryPermissionVar returns a pointer to a RepositoryPermission
func repositoryPermissionVar(p RepositoryPermission) *RepositoryPermission {
	return &p
}

// LicenseTemplate is an enum specifying a license template that can be used when creating a
// repository. Examples of available licenses are here:
// https://docs.github.com/en/github/creating-cloning-and-archiving-repositories/licensing-a-repository#searching-github-by-license-type
type LicenseTemplate string

const (
	// LicenseTemplateApache2 specifies use of the Apache 2.0 license, see
	// https://choosealicense.com/licenses/apache-2.0/
	LicenseTemplateApache2 = LicenseTemplate("apache-2.0")
	// LicenseTemplateMIT specifies use of the MIT license, see
	// https://choosealicense.com/licenses/mit/
	LicenseTemplateMIT = LicenseTemplate("mit")
	// LicenseTemplateGPL3 specifies use of the GNU General Public License v3.0, see
	// https://choosealicense.com/licenses/gpl-3.0/
	LicenseTemplateGPL3 = LicenseTemplate("gpl-3.0")
)

// knownLicenseTemplateValues is a map of known LicenseTemplate values, used for validation
var knownLicenseTemplateValues = map[LicenseTemplate]struct{}{
	LicenseTemplateApache2: {},
	LicenseTemplateMIT:     {},
	LicenseTemplateGPL3:    {},
}

// validateLicenseTemplate validates a given LicenseTemplate.
// Use as errs.Append(validateLicenseTemplate(template), template, "FieldName")
func validateLicenseTemplate(t LicenseTemplate) error {
	_, ok := knownLicenseTemplateValues[t]
	if !ok {
		return validation.ErrFieldEnumInvalid
	}
	return nil
}

// licenseTemplateVar returns a pointer to a LicenseTemplate
func licenseTemplateVar(t LicenseTemplate) *LicenseTemplate {
	return &t
}
