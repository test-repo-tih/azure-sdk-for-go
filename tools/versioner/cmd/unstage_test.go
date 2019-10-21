// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/tools/apidiff/report"
	"github.com/Azure/azure-sdk-for-go/tools/internal/modinfo"
)

const versionGoFormat = `package foo

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/%s foo/2019-04-01"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return "%s"
}

// tag: %s
`

func Test_getTags(t *testing.T) {
	if os.Getenv("TRAVIS") == "true" {
		// travis does a shallow clone so tag count is not consistent
		t.SkipNow()
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	tags, err := getTags(cwd, "v10")
	if err != nil {
		t.Fatalf("failed to get tags: %v", err)
	}
	if l := len(tags); l != 11 {
		t.Fatalf("expected 11 tags, got %d", l)
	}
	found := false
	for _, tag := range tags {
		if tag == "v10.1.0-beta" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("didn't find tag v10.1.0-beta")
	}
}

func Test_getTagPrefix(t *testing.T) {
	p, err := getTagPrefix(filepath.Join("work", "src", "github.com", "Azure", "azure-sdk-for-go", "services", "redis", "mgmt", "2018-03-01", "redis"))
	if err != nil {
		t.Fatal("failed to get tag prefix")
	}
	if p != "services/redis/mgmt/2018-03-01/redis" {
		t.Fatalf("wrong value '%s' for tag prefix", p)
	}
	p, err = getTagPrefix("/work/src/github.com/something/else")
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if p != "" {
		t.Fatalf("unexpected tag '%s'", p)
	}
}

type mockModInfo struct {
	dir     string
	exports bool
	breaks  bool
}

func (mock mockModInfo) DestDir() string {
	return mock.dir
}

func (mock mockModInfo) NewExports() bool {
	return mock.exports
}

func (mock mockModInfo) BreakingChanges() bool {
	return mock.breaks
}

func (mock mockModInfo) VersionSuffix() bool {
	return modinfo.HasVersionSuffix(mock.dir)
}

func (mock mockModInfo) NewModule() bool {
	// not needed by tests
	return false
}

func (mock mockModInfo) GenerateReport() report.Package {
	// not needed by tests
	return report.Package{}
}

func Test_calculateModuleTagMajorV1(t *testing.T) {
	pkg := mockModInfo{
		dir: filepath.Join("work", "src", "github.com", "Azure", "azure-sdk-for-go", "services", "foo"),
	}
	tag, err := calculateModuleTag([]string{}, pkg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tag != "services/foo/v1.0.0" {
		t.Fatalf("bad tag '%s", tag)
	}
}

func Test_calculateModuleTagMajorV2(t *testing.T) {
	tags := []string{
		"services/foo/v1.0.0",
		"services/foo/v1.1.0",
	}
	pkg := mockModInfo{
		dir:    filepath.Join("work", "src", "github.com", "Azure", "azure-sdk-for-go", "services", "foo", "v2"),
		breaks: true,
	}
	tag, err := calculateModuleTag(tags, pkg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tag != "services/foo/v2.0.0" {
		t.Fatalf("bad tag '%s", tag)
	}
}

func Test_calculateModuleTagMajorV2Invalid(t *testing.T) {
	tags := []string{
		"services/foo/v1.0.0",
		"services/foo/v1.1.0",
	}
	pkg := mockModInfo{
		dir:    filepath.Join("work", "src", "github.com", "Azure", "azure-sdk-for-go", "services", "foo"), // missing /v2 suffix
		breaks: true,
	}
	tag, err := calculateModuleTag(tags, pkg)
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if tag != "" {
		t.Fatal("expected no tag")
	}
}

func Test_calculateModuleTagMinorV1(t *testing.T) {
	tags := []string{
		"services/foo/v1.0.0",
		"services/foo/v1.0.1",
	}
	pkg := mockModInfo{
		dir:     filepath.Join("work", "src", "github.com", "Azure", "azure-sdk-for-go", "services", "foo"),
		exports: true,
	}
	tag, err := calculateModuleTag(tags, pkg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tag != "services/foo/v1.1.0" {
		t.Fatalf("bad tag '%s", tag)
	}
}

func Test_calculateModuleTagMinorV2(t *testing.T) {
	tags := []string{
		"services/foo/v1.0.0",
		"services/foo/v1.0.1",
		"services/foo/v2.0.0",
	}
	pkg := mockModInfo{
		dir: filepath.Join("work", "src", "github.com", "Azure", "azure-sdk-for-go", "services", "foo", "v2"),
	}
	tag, err := calculateModuleTag(tags, pkg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tag != "services/foo/v2.0.1" {
		t.Fatalf("bad tag '%s", tag)
	}
}

func Test_findLatestMajorVersion(t *testing.T) {
	ver, err := findLatestMajorVersion("../../testdata/scenarioa/foo/stage")
	if err != nil {
		t.Fatalf("failed to find LMV: %v", err)
	}
	if ver != filepath.Join("..", "..", "testdata", "scenarioa", "foo") {
		t.Fatalf("bad LMV %s", ver)
	}
	ver, err = findLatestMajorVersion("../../testdata/scenariod/foo/stage")
	if err != nil {
		t.Fatalf("failed to find LMV: %v", err)
	}
	if ver != filepath.Join("..", "..", "testdata", "scenariod", "foo", "v2") {
		t.Fatalf("bad LMV %s", ver)
	}
	ver, err = findLatestMajorVersion("../../testdata/scenariof/foo/stage")
	if err != nil {
		t.Fatalf("failed to find LMV: %v", err)
	}
	if ver != filepath.Join("..", "..", "testdata", "scenariof", "foo") {
		t.Fatalf("bad LMV %s", ver)
	}
}

type byteBufferSeeker struct {
	buf *bytes.Buffer
}

func (b byteBufferSeeker) Read(p []byte) (int, error) {
	return b.buf.Read(p)
}

func (b byteBufferSeeker) Write(p []byte) (int, error) {
	return b.buf.Write(p)
}

func (b byteBufferSeeker) Seek(offset int64, whence int) (int64, error) {
	if offset != 0 && whence != 0 {
		panic("seek only supports 0, 0")
	}
	b.buf.Reset()
	return 0, nil
}

func Test_updateGoModVerA(t *testing.T) {
	// updates from v1 to v2
	before := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo\n\n%s\n", goVersion)
	buf := byteBufferSeeker{
		buf: bytes.NewBuffer([]byte(before)),
	}
	err := updateGoModVer(buf, "v2")
	if err != nil {
		t.Fatalf("updateGoModVerA failed: %v", err)
	}
	after := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v2\n\n%s\n", goVersion)
	if !strings.EqualFold(buf.buf.String(), after) {
		t.Fatalf("bad go.mod update, epected %s got %s", after, buf.buf.String())
	}
}

func Test_updateGoModVerB(t *testing.T) {
	// updates from v2 to v3
	before := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v2\n\n%s\n", goVersion)
	buf := byteBufferSeeker{
		buf: bytes.NewBuffer([]byte(before)),
	}
	err := updateGoModVer(buf, "v3")
	if err != nil {
		t.Fatalf("updateGoModVerA failed: %v", err)
	}
	after := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/services/foo/mgmt/2019-05-01/foo/v3\n\n%s\n", goVersion)
	if !strings.EqualFold(buf.buf.String(), after) {
		t.Fatalf("bad go.mod update, epected %s got %s", after, buf.buf.String())
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func verifyVersion(t *testing.T, path, version, tag string) {
	b, err := ioutil.ReadFile(filepath.Join(path, "version.go"))
	if err != nil {
		t.Fatalf("failed to read version.go file: %v", err)
	}
	expected := fmt.Sprintf(versionGoFormat, version, version, tag)
	if !fileContentEquals(expected, string(b)) {
		t.Fatalf("bad version.go file, expected '%s' got '%s'", expected, string(b))
	}
}

func verifyGoMod(t *testing.T, path, expected string) {
	b, err := ioutil.ReadFile(filepath.Join(path, "go.mod"))
	if err != nil {
		t.Fatalf("failed to read go.mod file: %v", err)
	}
	if !fileContentEquals(expected, string(b)) {
		t.Fatalf("bad go.mod file, expected '%s' got '%s'", expected, string(b))
	}
}

func verifyChangelog(t *testing.T, path string) {
	if !fileExists(filepath.Join(path, "CHANGELOG.md")) {
		t.Fatalf("expected changelog in %s", path)
	}
}

func verifyNoChangelog(t *testing.T, path string) {
	if fileExists(filepath.Join(path, "CHANGELOG.md")) {
		t.Fatalf("unexpected changelog in %s", path)
	}
}

func verifyGoVet(t *testing.T, root string) {
	root, err := filepath.Abs(root)
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	folders := make([]string, 0)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			pkg := path[strings.Index(path, "github.com"):]
			folders = append(folders, pkg)
			return nil
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed to list all sub-folders in root '%s': %v", root, err)
	}
	for _, folder := range folders {
		c := exec.Command("go", "vet", folder)
		if output, err := c.CombinedOutput(); err != nil {
			t.Fatalf("vet failed: %s", string(output))
		}
	}
}

func fileContentEquals(expected, content string) bool {
	replacedContent := strings.ReplaceAll(content, "\r\n", "\n")
	return strings.EqualFold(replacedContent, expected)
}

// scenariob
func Test_theCommandImplMajor(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root, prefix string) ([]string, error) {
		// root doesn't matter
		if !strings.HasSuffix(prefix, "/testdata/scenariob/foo") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{
			"tools/testdata/scenariob/foo/v1.0.0",
			"tools/testdata/scenariob/foo/v1.1.0",
		}, nil
	}
	pkg, err := filepath.Abs("../../testdata/scenariob/foo/stage")
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	tag, err := theUnstageCommand([]string{pkg})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expectedTag = "tools/testdata/scenariob/foo/v2.0.0"
	if tag != expectedTag {
		t.Fatalf("bad tag, expected '%s' got '%s'", expectedTag, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariob/foo/v2\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenariob/foo/v2", expectedMod)
	verifyVersion(t, "../../testdata/scenariob/foo/v2", "2.0.0", tag)
	verifyChangelog(t, "../../testdata/scenariob/foo/v2")
	verifyGoVet(t, "../../testdata/scenariob/foo")
}

// scenarioa
func Test_theCommandImplMinor(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root, prefix string) ([]string, error) {
		// root doesn't matter
		if !strings.HasSuffix(prefix, "/testdata/scenarioa/foo") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{
			"tools/testdata/scenarioa/foo/v1.0.0",
		}, nil
	}
	pkg, err := filepath.Abs("../../testdata/scenarioa/foo/stage")
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	tag, err := theUnstageCommand([]string{pkg})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expectedTag = "tools/testdata/scenarioa/foo/v1.1.0"
	if tag != expectedTag {
		t.Fatalf("bad tag, expected '%s' got '%s'", expectedTag, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioa/foo\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenarioa/foo", expectedMod)
	verifyVersion(t, "../../testdata/scenarioa/foo", "1.1.0", tag)
	verifyChangelog(t, "../../testdata/scenarioa/foo")
	verifyGoVet(t, "../../testdata/scenarioa/foo")
}

// scenarioc
func Test_theCommandImplPatch(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root, prefix string) ([]string, error) {
		// root doesn't matter
		if !strings.HasSuffix(prefix, "/testdata/scenarioc/foo") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{
			"tools/testdata/scenarioc/foo/v1.0.0",
		}, nil
	}
	pkg, err := filepath.Abs("../../testdata/scenarioc/foo/stage")
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	tag, err := theUnstageCommand([]string{pkg})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expected = "tools/testdata/scenarioc/foo/v1.0.1"
	if tag != expected {
		t.Fatalf("bad tag, expected '%s' got '%s'", expected, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioc/foo\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenarioc/foo", expectedMod)
	verifyVersion(t, "../../testdata/scenarioc/foo", "1.0.1", tag)
	verifyChangelog(t, "../../testdata/scenarioc/foo")
	verifyGoVet(t, "../../testdata/scenarioc/foo")
}

// scenariod
func Test_theCommandImplMajorV3(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root, prefix string) ([]string, error) {
		if !strings.HasSuffix(prefix, "/testdata/scenariod/foo/v2") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{
			"tools/testdata/scenariod/foo/v1.0.0",
			"tools/testdata/scenariod/foo/v1.0.1",
			"tools/testdata/scenariod/foo/v1.1.0",
			"tools/testdata/scenariod/foo/v1.2.0",
			"tools/testdata/scenariod/foo/v2.0.0",
			"tools/testdata/scenariod/foo/v2.1.0",
			"tools/testdata/scenariod/foo/v2.1.1",
		}, nil
	}
	pkg, err := filepath.Abs("../../testdata/scenariod/foo/stage")
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	tag, err := theUnstageCommand([]string{pkg})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expectedTag = "tools/testdata/scenariod/foo/v3.0.0"
	if tag != expectedTag {
		t.Fatalf("bad tag, expected '%s' got '%s'", expectedTag, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariod/foo/v3\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenariod/foo/v3", expectedMod)
	verifyVersion(t, "../../testdata/scenariod/foo/v3", "3.0.0", tag)
	verifyChangelog(t, "../../testdata/scenariod/foo/v3")
	verifyGoVet(t, "../../testdata/scenariod/foo")
}

// scenarioe
func Test_theCommandImplMajorMinor(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root, prefix string) ([]string, error) {
		// root doesn't matter
		if !strings.HasSuffix(prefix, "/testdata/scenarioe/foo/v2") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{
			"tools/testdata/scenarioe/foo/v1.0.0",
			"tools/testdata/scenarioe/foo/v1.1.0",
			"tools/testdata/scenarioe/foo/v2.0.0",
			"tools/testdata/scenarioe/foo/v2.1.0",
		}, nil
	}
	pkg, err := filepath.Abs("../../testdata/scenarioe/foo/stage")
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	tag, err := theUnstageCommand([]string{pkg})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expectedTag = "tools/testdata/scenarioe/foo/v2.2.0"
	if tag != expectedTag {
		t.Fatalf("bad tag, expected '%s' got '%s'", expectedTag, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioe/foo/v2\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenarioe/foo/v2", expectedMod)
	verifyVersion(t, "../../testdata/scenarioe/foo/v2", "2.2.0", tag)
	verifyChangelog(t, "../../testdata/scenarioe/foo/v2")
	verifyGoVet(t, "../../testdata/scenarioe/foo")
}

// scenariof
func Test_theCommandImplNewMod(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root, prefix string) ([]string, error) {
		// root doesn't matter
		if !strings.HasSuffix(prefix, "/testdata/scenariof/foo") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{}, nil
	}
	pkg, err := filepath.Abs("../../testdata/scenariof/foo/stage")
	if err != nil {
		t.Fatalf("failed to get absolute path: %v", err)
	}
	tag, err := theUnstageCommand([]string{pkg})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expectedTag = "tools/testdata/scenariof/foo/v1.0.0"
	if tag != expectedTag {
		t.Fatalf("bad tag, expected '%s' got '%s'", expectedTag, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariof/foo\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenariof/foo", expectedMod)
	verifyVersion(t, "../../testdata/scenariof/foo", "1.0.0", tag)
	verifyNoChangelog(t, "../../testdata/scenariof/foo")
	verifyGoVet(t, "../../testdata/scenariof/foo")
}

// scenariog
func Test_theCommandNewMgmtMajorV2(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root string, prefix string) ([]string, error) {
		if !strings.HasPrefix(prefix, "/testdata/scenariog/foo") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{
			"/testdata/scenariog/foo/mgmt/2019-10-11/foo/v1.0.0",
		}, nil
	}
	stage, err := filepath.Abs("../../testdata/scenariog/foo/mgmt/2019-10-11/foo/stage")
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	tag, err := theUnstageCommand([]string{stage})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expectedTag = "tools/testdata/scenariog/foo/mgmt/2019-10-11/foo/v2.0.0"
	if tag != expectedTag {
		t.Fatalf("bad tag, expected '%s' got '%s'", expectedTag, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenariog/foo/v2\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenariog/foo/mgmt/2019-10-11/foo/v2", expectedMod)
	verifyVersion(t, "../../testdata/scenariog/foo/mgmt/2019-10-11/foo/v2", "2.0.0", tag)
	verifyChangelog(t, "../../testdata/scenariog/foo/mgmt/2019-10-11/foo/v2")
	verifyGoVet(t, "../../testdata/scenariog/foo/mgmt/2019-10-11/foo")
}

// scenarioh
func Test_theCommandNewMgmtMajorV2WithOneLineImport(t *testing.T) {
	cleanTestData()
	defer cleanTestData()
	getTagsHook = func(root string, prefix string) ([]string, error) {
		if !strings.HasPrefix(prefix, "/testdata/scenarioh/foo") {
			return nil, fmt.Errorf("bad prefix '%s'", prefix)
		}
		return []string{
			"/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v1.0.0",
			"/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v1.1.0",
			"/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v1.2.0",
		}, nil
	}
	stage, err := filepath.Abs("../../testdata/scenarioh/foo/mgmt/2019-10-11/foo/stage")
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	tag, err := theUnstageCommand([]string{stage})
	if err != nil {
		t.Fatalf("failed: %v", err)
	}
	const expectedTag = "tools/testdata/scenarioh/foo/mgmt/2019-10-11/foo/v2.0.0"
	if tag != expectedTag {
		t.Fatalf("bad tag, expected '%s' got '%s'", expectedTag, tag)
	}
	expectedMod := fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/tools/testdata/scenarioh/foo/v2\n\n%s\n", goVersion)
	verifyGoMod(t, "../../testdata/scenarioh/foo/mgmt/2019-10-11/foo/v2", expectedMod)
	verifyVersion(t, "../../testdata/scenarioh/foo/mgmt/2019-10-11/foo/v2", "2.0.0", tag)
	verifyChangelog(t, "../../testdata/scenarioh/foo/mgmt/2019-10-11/foo/v2")
	verifyGoVet(t, "../../testdata/scenarioh/foo/mgmt/2019-10-11/foo")
}
