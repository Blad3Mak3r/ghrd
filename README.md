# GitHub Release Dowloader
GHRD is a CLI that allows you to download an artifact of the latest release from a repository on GitHub.

## How to use
Download the latest release from [here](#downloads) into your computer.

To download releases from GitHub Releases you need the following things:
- **Owner:** the owner of the repository (``--owner Blad3Mak3r``).
- **Repository:** the name of the repository (``--repo ghrd``).
- **Artifact:** the name of the artifact to download with extension (``--artifact ghrd-x86_64.exe`` or ``--artifact ghrd-i386-linux``).
- **GPA Token:** the GitHub Personal Access Token to authenticate the necesary requests to download the artifact (``--token ghp_xxxxx``).

Now you have to modes to execute the CLI:
### Direct mode
```bash
./ghrd --owner Blad3Mak3r --repo ghrd --artifact ghrd-x86_64-darwin --token ghp_xxxxx
```
Put arguments directly on the same line.

### Prompt mode
```bash
./ghrd --prompt
```
The script will ask for arguments that where not defined.
```bash
./ghrd --token ghp_xxxxx --prompt
```
Only will ask for --owner, --repo and --artifact.

## Help
```bash
./ghrd --help
```

## Downloads
- [Windows i386][win-i386]
- [Windows amd64][win-amd64]
- Windows ARM arch64(``comming soon``)
- [Linux i386][linux-i386]
- [Linux amd64][linux-amd64]
- Linux ARM arch64(``comming soon``)
- [MacOS amd64][darwin-amd64]
- MacOS ARM arch64(``comming soon``)


[win-i386]: https://github.com/Blad3Mak3r/ghrd/releases/download/0.2.0/ghrd-i386.exe
[win-amd64]: https://github.com/Blad3Mak3r/ghrd/releases/download/0.2.0/ghrd-x86_64.exe
[linux-i386]: https://github.com/Blad3Mak3r/ghrd/releases/download/0.2.0/ghrd-i386-linux
[linux-amd64]: https://github.com/Blad3Mak3r/ghrd/releases/download/0.2.0/ghrd-x86_64-linux
[darwin-amd64]: https://github.com/Blad3Mak3r/ghrd/releases/download/0.2.0/ghrd-x86_64-darwin
