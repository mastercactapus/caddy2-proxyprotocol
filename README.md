# Add PROXY protocol support to Caddy 2

`proxy_protocol` is a listener wrapper for Caddy 2 that adds support for PROXY headers on new connections.

## Configuration

### Options

|Name|Type|Default|Description|
|---|---|---|---|
|`timeout`|duration|`5s`|Specifies the maximum time for the PROXY header to be received. If zero, timeout is disabled.|
|`allow`|[]string|`0.0.0.0\0`|A list of CIDR ranges to allow/require PROXY headers from.|

### JSON

The wrapper needs to be loaded BEFORE the `tls` wrapper.

```js
{
  "apps": {
    "http": {
      "servers": {
        "myserver": {
          // ...
          "listener_wrappers":[
            {"wrapper": "proxy_protocol", "timeout": "5s", "allow": ["192.168.86/24"]},
            {"wrapper":"tls"}
          ]
          // ...
        }
      }
    }
  }
}
```

### Caddyfile

The wrapper may be configured via [global options](https://caddyserver.com/docs/caddyfile/options#listener-wrappers) in the Caddyfile.

```
{
  servers {
    listener_wrappers {
      proxy_protocol {
        timeout <duration>
        allow <cidrs...>
      }
      tls
    }
  }
}
```
