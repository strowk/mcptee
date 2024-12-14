# mcptee

A simple tool that proxies input from MCP client to MCP server and returns the output back to the client, while also logging both to the file line by line indicating the direction of the data with a prefixes "in:" and "out:", resulting in a file that is normally valid YAML.

Before every "in:", tool would print `---` and endline that would make file multi-document YAML, so normally every combination of request/respons would be a separate document.

This could be pretty handy for testing and troubleshooting MCP clients and servers.

## Installation

Download prebulit binaries from the [releases](https://github.com/strowk/mcptee/releases) page and put in your PATH, or build from source:

```bash
go get github.com/strowk/mcptee
go install github.com/strowk/mcptee
```

## Usage

```bash
mcptee <out-file> <command> [args...]
```

Some examples:
```bash
mcptee ./server_log.yaml mcp-k8s-go

mcptee ./server_log.yaml npx @strowk/mcp-k8s
```

You would probably configure it with corresponding client. For example, with Claude:

```json
{
    "mcpServers": {
        "my_server": {
            "command": "mcptee",
            "args": ["/path/to/log_file.yaml", "npx", "my_mcp_server"]
        }
    }
}
```

## Example

This is taken from actual server log session between Zed client and mcp-k8s-go server:

```yaml
---
in: {"jsonrpc":"2.0","id":0,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"Zed","version":"0.1.0"}}}
out: {"jsonrpc":"2.0","result":{"capabilities":{"prompts":{"listChanged":false},"resources":{"listChanged":false,"subscribe":false}},"protocolVersion":"2024-11-05","serverInfo":{"name":"mcp-k8s-go","version":"0.0.1"}},"id":0}
---
in: {"jsonrpc":"2.0","method":"notifications/initialized","params":{}}
---
in: {"jsonrpc":"2.0","id":1,"method":"prompts/list","params":{}}
out: {"jsonrpc":"2.0","result":{"prompts":[{"arguments":[{"description":"Namespace to list Pods from, defaults to all namespaces","name":"namespace","required":false}],"description":"List Kubernetes Pods with name and namespace in the current context","name":"list-k8s-pods"}]},"id":1}
---
in: {"jsonrpc":"2.0","id":2,"method":"completion/complete","params":{"ref":{"type":"ref/prompt","name":"list-k8s-pods"},"argument":{"name":"namespace","value":""}}}
out: {"jsonrpc":"2.0","result":{"completion":{"hasMore":false,"total":5,"values":["default","kube-node-lease","kube-public","kube-system","test"]}},"id":2}
---
in: {"jsonrpc":"2.0","id":3,"method":"completion/complete","params":{"ref":{"type":"ref/prompt","name":"list-k8s-pods"},"argument":{"name":"namespace","value":"k"}}}
out: {"jsonrpc":"2.0","result":{"completion":{"hasMore":false,"total":3,"values":["kube-node-lease","kube-public","kube-system"]}},"id":3}
---
in: {"jsonrpc":"2.0","id":4,"method":"completion/complete","params":{"ref":{"type":"ref/prompt","name":"list-k8s-pods"},"argument":{"name":"namespace","value":"ku"}}}
out: {"jsonrpc":"2.0","result":{"completion":{"hasMore":false,"total":3,"values":["kube-node-lease","kube-public","kube-system"]}},"id":4}
---
in: {"jsonrpc":"2.0","id":5,"method":"completion/complete","params":{"ref":{"type":"ref/prompt","name":"list-k8s-pods"},"argument":{"name":"namespace","value":"kub"}}}
out: {"jsonrpc":"2.0","result":{"completion":{"hasMore":false,"total":3,"values":["kube-node-lease","kube-public","kube-system"]}},"id":5}
---
in: {"jsonrpc":"2.0","id":6,"method":"completion/complete","params":{"ref":{"type":"ref/prompt","name":"list-k8s-pods"},"argument":{"name":"namespace","value":"kube"}}}
out: {"jsonrpc":"2.0","result":{"completion":{"hasMore":false,"total":3,"values":["kube-node-lease","kube-public","kube-system"]}},"id":6}
---

