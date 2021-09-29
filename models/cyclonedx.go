package models

import (
	"encoding/json"
	"os"
)

type CycloneDxSwidTag struct {
	TagId      string              `json:"tagId,omitempty"`
	Name       string              `json:"name,omitempty"`
	Version    string              `json:"version,omitempty"`
	TagVersion int                 `json:"tagVersion,omitempty"`
	Patch      bool                `json:"patch,omitempty"`
	Text       CycloneDxAttachment `json:"text,omitempty"`
	Url        string              `json:"url,omitempty"`
}

type CycloneDxIdentifiableAction struct {
	Timestamp string `json:"timestamp,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type CycloneDxCommit struct {
	Uid       string                      `json:"uid,omitempty"`
	Url       string                      `json:"url,omitempty"`
	Author    CycloneDxIdentifiableAction `json:"author,omitempty"`
	Committer CycloneDxIdentifiableAction `json:"committer,omitempty"`
	Message   string                      `json:"message,omitempty"`
}

type CycloneDxDiff struct {
	Text CycloneDxAttachment `json:"text,omitempty"`
	Url  string              `json:"string,omitempty"`
}

type CycloneDxPatch struct {
	Type     string          `json:"type,omitempty"`
	Diff     CycloneDxDiff   `json:"diff,omitempty"`
	Resolves []CycloneDxDiff `json:"resolves,omitempty"`
}

type CycloneDxComponentPedigree struct {
	Ancestors   []CycloneDxComponent `json:"ancestors,omitempty"`
	Descendants []CycloneDxComponent `json:"descendants,omitempty"`
	Variants    []CycloneDxComponent `json:"variants,omitempty"`
	Commits     []CycloneDxCommit    `json:"commits,omitempty"`
	Patches     []CycloneDxPatch     `json:"patches,omitempty"`
	Notes       string               `json:"notes,omitempty"`
}

type CycloneDxExternalReference struct {
	Url     string        `json:"url"`
	Comment string        `json:"comment,omitempty"`
	Type    string        `json:"type"`
	Hashes  CycloneDxHash `json:"hashes,omitempty"`
}

type CycloneDxCopyright struct {
	Text string `json:"text,omitempty"`
}

type CycloneDxEvidence struct {
	Licenses  []CycloneDxLicenseEntry `json:"licenses,omitempty"`
	Copyright []CycloneDxCopyright    `json:"copyright,omitempty"`
}

type CycloneDxService struct {
	BomRef             string                        `json:"bom-ref,omitempty"`
	Provider           CycloneDxOrganizationalEntity `json:"provider,omitempty"`
	Group              string                        `json:"group,omitempty"`
	Name               string                        `json:"name"`
	Version            string                        `json:"version,omitempty"`
	Description        string                        `json:"description,omitempty"`
	Endpoints          []string                      `json:"endpoints,omitempty"`
	Authenticated      bool                          `json:"authenticated,omitempty"`
	XTrustBoundary     bool                          `json:"x-trust-boundary,omitempty"`
	Data               []CycloneDxHash               `json:"data,omitempty"`
	Licenses           []CycloneDxLicenseEntry       `json:"licenses,omitempty"`
	ExternalReferences []CycloneDxExternalReference  `json:"externalReferences,omitempty"`
	Services           []CycloneDxService            `json:"services,omitempty"`
	Properties         []CycloneDxNameValuePair      `json:"properties,omitempty"`
}

type CycloneDxComponent struct {
	Type               string                        `json:"type"`
	MimeType           string                        `json:"mime-type,omitempty"`
	BomRef             string                        `json:"bom-ref,omitempty"`
	Supplier           CycloneDxOrganizationalEntity `json:"supplier,omitempty"`
	Author             string                        `json:"author,omitempty"`
	Publisher          string                        `json:"publisher,omitempty"`
	Group              string                        `json:"group,omitempty"`
	Name               string                        `json:"name"`
	Version            string                        `json:"version"`
	Description        string                        `json:"description,omitempty"`
	Scope              string                        `json:"scope,omitempty"`
	Hashes             []CycloneDxHash               `json:"hashes,omitempty"`
	Licenses           []CycloneDxLicenseEntry       `json:"licenses,omitempty"`
	Copyright          string                        `json:"copyright,omitempty"`
	Cpe                string                        `json:"cpe,omitempty"`
	Purl               string                        `json:"purl,omitempty"`
	Swid               CycloneDxSwidTag              `json:"swid,omitempty"`
	Modified           bool                          `json:"modified,omitempty"`
	Pedigree           CycloneDxComponentPedigree    `json:"pedigree,omitempty"`
	ExternalReferences []CycloneDxExternalReference  `json:"externalReferences,omitempty"`
	Components         []CycloneDxComponent          `json:"components,omitempty"`
	Evidence           CycloneDxEvidence             `json:"evidence,omitempty"`
	Properties         []CycloneDxNameValuePair      `json:"properties,omitempty"`
}

type CycloneDxOrganizationalContact struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type CycloneDxOrganizationalEntity struct {
	Name    string                           `json:"name,omitempty"`
	Url     []string                         `json:"url,omitempty"`
	Contact []CycloneDxOrganizationalContact `json:"contact,omitempty"`
}

type CycloneDxHash struct {
	Alg     string `json:"alg"`
	Content string `json:"content"`
}

type CycloneDxTools struct {
	Vendor  string          `json:"vendor,omitempty"`
	Name    string          `json:"name,omitempty"`
	Version string          `json:"version,omitempty"`
	Hashes  []CycloneDxHash `json:"hashes,omitempty"`
}

type CycloneDxAttachment struct {
	ContentType string `json:"contentType,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Content     string `json:"content"`
}

type CycloneDxLicense struct {
	Id   string              `json:"id,omitempty"`
	Name string              `json:"name,omitempty"`
	Text CycloneDxAttachment `json:"text,omitempty"`
	Url  string              `json:"url,omitempty"`
}

type CycloneDxLicenseEntry struct {
	License    CycloneDxLicense `json:"license,omitempty"`
	Expression string           `json:"expression,omitempty"`
}

type CycloneDxNameValuePair struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type CycloneDxMetadata struct {
	Timestamp   string                           `json:"timestamp,omitempty"`
	Tools       []CycloneDxTools                 `json:"tools,omitempty"`
	Authors     []CycloneDxOrganizationalContact `json:"authors,omitempty"`
	Component   CycloneDxComponent               `json:"component,omitempty"`
	Manufacture CycloneDxOrganizationalEntity    `json:"manufacture,omitempty"`
	Supplier    CycloneDxOrganizationalEntity    `json:"supplier,omitempty"`
	Licenses    []CycloneDxLicenseEntry          `json:"licenses,omitempty"`
	Properties  []CycloneDxNameValuePair         `json:"properties,omitempty"`
}

type CycloneDxDependency struct {
	Ref       string   `json:"ref"`
	DependsOn []string `json:"dependsOn,omitempty"`
}

type CycloneDxComposition struct {
	Aggregate    string                `json:"aggregate"`
	Assemblies   []string              `json:"assemblies,omitempty"`
	Dependencies []CycloneDxDependency `json:"dependencies,omitempty"`
}

type CycloneDxBom struct {
	BomFormat          string                       `json:"bomFormat"`
	SpecVersion        string                       `json:"specVersion"`
	SerialNumber       string                       `json:"serialNumber,omitempty"`
	Version            int                          `json:"version"`
	Metadata           CycloneDxMetadata            `json:"metadata,omitempty"`
	Components         []CycloneDxComponent         `json:"components,omitempty"`
	Services           []CycloneDxService           `json:"services,omitempty"`
	ExternalReferences []CycloneDxExternalReference `json:"externalReferences,omitempty"`
	Dependencies       []CycloneDxDependency        `json:"dependencies,omitempty"`
	Compositions       []CycloneDxComposition       `json:"compositions,omitempty"`
}

func ParseCycloneDxBom(path string) CycloneDxBom {
	bom := CycloneDxBom{}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &bom)
	if err != nil {
		panic(err)
	}

	//TODO: get rid of this, it's currently only here for testing purposes
	marshalled, err := json.MarshalIndent(bom, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("bom-output.json", marshalled, 0777)
	if err != nil {
		panic(err)
	}

	return bom
}
