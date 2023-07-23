# Try Golang

## Preparation

- Install [ms-vscode-remote.remote-containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension
- **If using SSH to clone**: Add ssh key to SSH agent <https://code.visualstudio.com/remote/advancedcontainers/sharing-git-credentials>
  ```
  ssh-add ~/.ssh/id_rsa
  ```
- Open repo in vscode devcontainer with:

  - CTRL+SHIFT+P -> Dev-Containers: Clone Repository in Named Volume
  - Paste repo HTTPS or SSH url
  - Choose volume name
  - Open terminal inside vscode to get a shell inside devcontainer

## Development

```bash
make codegen
nf start
```

## Test

```bash
make test
```

## Links

- <https://github.com/golang/tools/blob/master/gopls/doc/workspace.md>
- <https://protobuf.dev/reference/go/go-generated/>
