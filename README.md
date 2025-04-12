# mcptee

A simple tool that proxies input from MCP client to MCP server and returns the output back to the client, while also logging both to the file line by line indicating the direction of the data with a prefixes "in:" and "out:", resulting in a file that is normally valid YAML.

Before every "in:", tool would print `---` and endline that would make file multi-document YAML, so normally every combination of request/respons would be a separate document.

This could be pretty handy for testing and troubleshooting MCP clients and servers.

## Installation

## npm

```bash
npm install -g @strowk/mcptee
```

## Github Releases

Download prebulit binaries from the [releases](https://github.com/strowk/mcptee/releases) page and put in your PATH

## Build from source

```bash
go install github.com/strowk/mcptee@latest
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

If you have for example a postgres server running in docker:

```bash
docker run --rm --name postgres-mcp-test -e POSTGRES_PASSWORD=thesecret -p 7777:5432 postgres:latest
```

, and you want to test it with mcp-server-postgres (run it in another terminal):

```bash
mcptee log.yaml npx @modelcontextprotocol/server-postgres postgres://postgres:thesecret@localhost:7777
```

Now send these one by one:

```json
{"jsonrpc":"2.0","id":0,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"testing","version":"0.0.1"}}}

{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}

{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"query", "arguments": {"sql": "SELECT datname FROM pg_database"} }}
```

You should see the output in the `log.yaml` file like this:

```yaml
---
in: {"jsonrpc":"2.0","id":0,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"testing","version":"0.0.1"}}}
out: {"result":{"protocolVersion":"2024-11-05","capabilities":{"resources":{},"tools":{}},"serverInfo":{"name":"example-servers/postgres","version":"0.1.0"}},"jsonrpc":"2.0","id":0}
---
in: {"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}
out: {"result":{"tools":[{"name":"query","description":"Run a read-only SQL query","inputSchema":{"type":"object","properties":{"sql":{"type":"string"}}}}]},"jsonrpc":"2.0","id":1}
---
in: {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"query", "arguments": {"sql": "SELECT datname FROM pg_database"} }}
out: {"result":{"content":[{"type":"text","text":"[\n  {\n    \"datname\": \"postgres\"\n  },\n  {\n    \"datname\": \"template1\"\n  },\n  {\n    \"datname\": \"template0\"\n  }\n]"}],"isError":false},"jsonrpc":"2.0","id":1}
```

