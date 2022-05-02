package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	. "github.com/blad3mak3r/ghrd/structs"
)

var (
	owner    string
	repo     string
	token    string
	artifact string

	help    bool
	version bool
)

const (
	major = 0
	minor = 1
	path  = 1

	apiUrl               = "https://api.github.com"
	acceptHeader         = "application/vnd.github.v3+json"
	acceptDownloadHeader = "application/octet-stream"
)

func getRevision() string {
	out, err := exec.Command("sh", "-c", "git rev-parse --short HEAD").Output()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", out)
}

func getVersion() string {
	version := fmt.Sprintf("%d.%d.%d", major, minor, path)
	revision := getRevision()

	if len(revision) > 0 {
		return fmt.Sprintf("%s_%s", version, revision)
	} else {
		return version
	}
}

func error(str string) {
	fmt.Println("\033[31m" + str + "\033[0m\n")
	os.Exit(1)
}

func main() {

	flag.BoolVar(&help, "help", false, "Show this help menu")
	flag.BoolVar(&version, "version", false, "Print GHRD version.")

	flag.StringVar(&owner, "owner", "", "The GitHub repository owner.")
	flag.StringVar(&repo, "repo", "", "The GitHub repository name.")
	flag.StringVar(&token, "token", "", "The GitHub Personal Access Token.")
	flag.StringVar(&artifact, "artifact", "", "The artifact to download (with .extension).")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Printf("✅ GitHub Release Downloader v%s\n", getVersion())
		os.Exit(0)
	}

	if len(owner) < 1 {
		error("⚠️ Flag '--owner' is not present, use --help for more info.")
	}

	if len(repo) < 1 {
		error("⚠️ Flag '--repo' is not present, use --help for more info.")
	}

	if len(token) < 1 {
		error("⚠️ Flag '--token' is not present, use --help for more info.")
	}

	if len(artifact) < 1 {
		error("⚠️ Flag '--artifact' is not present, use --help for more info.")
	}

	fmt.Printf("⏳ Initializing GitHub Release Downloader (GHRD) %s\n", getVersion())

	httpClient := http.Client{}

	release := getLatestRealease(&httpClient)

	fmt.Printf("• Found latest release with name '%s' and tag '%s'.\n", release.Name, release.TagName)

	asset, exists := getAssetFromRelease(release)
	if !exists {
		error("⚠️ Asset with name " + artifact + " wasn't found on release " + release.TagName)
	}

	downloadAsset(&httpClient, asset)
}

func getLatestRealease(httpClient *http.Client) GithubRelease {
	release := GithubRelease{}

	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", apiUrl, owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error(err.Error())
	}

	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	res, err := httpClient.Do(req)
	if err != nil {
		error(err.Error())
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&release)
	if err != nil {
		error(err.Error())
	}

	return release
}

func getAssetFromRelease(release GithubRelease) (GithubAsset, bool) {
	for _, asset := range release.Assets {
		if asset.Name == artifact {
			return asset, true
		}
	}
	return GithubAsset{}, false
}

func downloadAsset(httpClient *http.Client, asset GithubAsset) {
	url := asset.Url

	fmt.Printf("• Downloading asset %s from url %s\n", asset.Name, url)

	file, err := os.Create(artifact)
	if err != nil {
		error(err.Error())
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error(err.Error())
	}

	req.Header.Add("Accept", acceptDownloadHeader)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	res, err := httpClient.Do(req)
	if err != nil {
		error(err.Error())
	}

	defer res.Body.Close()

	size, err := io.Copy(file, res.Body)
	if err != nil {
		error(err.Error())
	}

	defer file.Close()

	fmt.Printf("• Downloaded %s with size %d from %s\n", artifact, size, url)

	fmt.Printf("• Applying required execution permissions to %s\n", artifact)

	err = file.Chmod(0770)
	if err != nil {
		error(err.Error())
	}

	fmt.Printf("• Applied Execution permissions to user and group.\n")
	fmt.Print("\n\n\033[32m🎉 Done! 🎉\033[0m\n\n")
}
