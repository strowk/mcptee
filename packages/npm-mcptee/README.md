# mcptee

A simple tool that proxies input from MCP client to MCP server and returns the output back to the client, while also logging both to the file line by line indicating the direction of the data with a prefixes "in:" and "out:", resulting in a file that is normally valid YAML.

Before every "in:", tool would print `---` and endline that would make file multi-document YAML, so normally every combination of request/respons would be a separate document.

This could be pretty handy for testing and troubleshooting MCP clients and servers.

## Usage

```bash
mcptee <out-file> <command> [args...]
```

Some examples:
```bash
mcptee ./server_log.yaml mcp-k8s-go

mcptee ./server_log.yaml npx @strowk/mcp-k8s
```

