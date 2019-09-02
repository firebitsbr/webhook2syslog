# webhook2syslog

This simple HTTP server listens for JSON POST messages and writes them to syslog.

### Usage

```bash
Usage of ./webhook2syslog:
  -host string
    	HTTP server address (IPv4 or hostname). (default "localhost")
  -message string
    	Syslog message prefix. (default "[TheHive Webhook]")
  -port int
    	HTTP listener port. (default 5001)
```

### License

This software is licensed under GNU Affero General Public License version 3

    Copyright (C) 2019 Roger Johnston