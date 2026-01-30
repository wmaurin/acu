# Acu

A cross-platform CLI toolkit for common terminal tasks.

## Installation

```bash
go install github.com/wmaurin/acu@latest    # Install
acu --help                                  # Verify
```

## Commands

### Docker

```bash
acu docker clean [--all]       # Remove stopped containers, dangling images, unused networks
                               # --all: Stop and remove all containers, images, networks, cache
```

### Port

```bash
acu port find <port>           # Find what's using a port (pid, process name, user)
acu port kill <port>           # Kill process using a port
acu port scan <host>           # Scan common ports on a host
```

### Disk & Files

```bash
acu disk largest [path]        # Find largest files/directories (sorted, human-readable)
acu disk clean [path]          # Find and remove common junk (node_modules, __pycache__, .DS_Store, etc.)
```

### Network

```bash
acu net speed                      # Network speed test
acu net http <url>                 # HTTP request with timing breakdown (DNS, connect, TLS, TTFB)
```

### SSH

```bash
acu ssh list                                         # List saved SSH connections
acu ssh add <alias> <user@host> [-i key] [-p port]   # Save SSH connection
acu ssh rm <alias>                                   # Remove saved connection
acu ssh tunnel <alias> -L <local>:<remote>           # Create SSH tunnel with saved connection
```

### Generators

```bash
acu gen password [length]      # Generate secure password (default: 16 chars)
acu gen uuid                   # Generate UUID v4
acu gen secret [length]        # Generate base64 secret for configs
```

### Encoding

```bash
acu codec encode <format> <string>   # Encode string (format: base64, url, html)
acu codec decode <format> <string>   # Decode string (format: base64, url, html)
acu codec hash <file>                # Generate md5, sha1, sha256 checksums for a file
```

### Terminal-Agent (maybe? Llama?)

```bash
acu ask <question>             # Ask any questions, can also execute on approval
```
