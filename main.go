package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/cavaliercoder/go-rpm"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/sontags/env"
	"github.com/unprofession-al/pkgpile/yum"
)

var version string
var commitId string

var config = &Configuration{
	Version:  version,
	CommitId: commitId,
}

var l = NewLogger()
var filenameTemplate *template.Template

func init() {
	env.Var(&config.Port, "PORT", "8080", "Port to bind")
	env.Var(&config.Base, "BASE", "/tmp/", "Base directories of the repos")
	env.Var(&config.DebugStr, "DEBUG", "false", "Turn debugging on (only print commands to be run)")
	env.Var(&config.FilenameTemplate, "FILENAME_TEMPLATE", "{{.Name}}-{{.Version}}-{{.Release}}.{{.Architecture}}.rpm", "Turn debugging on (only print commands to be run)")
}

var store = map[string]map[string]rpm.PackageFile{}
var repodata = map[string]yum.RepoData{}

func main() {
	env.Parse("PKGPILE", false)

	var err error
	filenameTemplate, err = template.New("filename").Parse(config.FilenameTemplate)
	if err != nil {
		panic(err)
	}

	config.Debug = !strings.Contains(config.DebugStr, "false")

	r := mux.NewRouter()
	r.HandleFunc("/{repo}/", UploadPackage).Methods("POST")
	r.HandleFunc("/{repo}/repodata/filelists.xml.gz", GetFilelists).Methods("GET")
	r.HandleFunc("/{repo}/repodata/other.xml.gz", GetOther).Methods("GET")
	r.HandleFunc("/{repo}/repodata/primary.xml.gz", GetPrimary).Methods("GET")
	r.HandleFunc("/{repo}/repodata/repomd.xml", GetRepomd).Methods("GET")
	r.HandleFunc("/config.json", GetConfig).Methods("GET")
	chain := alice.New().Then(r)

	l.l("starting...", "pkgpile should be ready")

	log.Fatal(http.ListenAndServe(":"+config.Port, chain))
}
