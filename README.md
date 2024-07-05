
# GOHealthChecker

## Purpose

The purpose of the project is to create a command-line interface (CLI) application in Go. This application is called "Healthchecker," and it is designed to check whether a given domain is reachable on a specified port.

## Purpose of the Code

The primary purpose of the code is to:

1. Accept Command-Line Arguments:
   - The application accepts two flags:
     - `--domain` (or `-d`): A required flag that specifies the domain name to check.
     - `--port` (or `-p`): An optional flag that specifies the port number to check (defaults to port 80 if not provided).
2. Check Domain Reachability:
   - The application checks if the specified domain is reachable on the given port by performing a network check. This logic is implemented in the `Check` function (defined in another file, `check.go`).
3. Print the Status:
   - The application prints the result of the reachability check to the console, indicating whether the domain is "UP" or "DOWN".

## Output of the Code


```markdown
![Alt text](https://github.com/nihar-15/GoHealthCheck/blob/main/sampleOutput.png)
```



Explanation of the Output:

- `[UP] google.com is reachable at google.com:80`
  - This indicates that the domain google.com is reachable on port 80 (default HTTP port).
- `DNS: 172.217.174.238 2404:6800:4009:801::200e`
  - These are the IP addresses resolved from the domain name google.com. The first IP address (172.217.174.238) is an IPv4 address, and the second (2404:6800:4009:801::200e) is an IPv6 address.
- `Response Time: 263.917818ms`
  - This is the time it took to establish a TCP connection to google.com on port 80. The connection was established in approximately 263.92 milliseconds.
- `From: 10.0.2.15:36004`
  - This indicates the local IP address and port used for the outgoing connection. In this case, the local machine's IP address is 10.0.2.15 and the port is 36004.
- `To: 172.217.174.238:80`
  - This indicates the remote IP address and port to which the connection was made. Here, the connection was made to the IP address 172.217.174.238 on port 80.
- `[HTTP] http://google.com responded with status 200 OK`
  - This indicates that an HTTP GET request to http://google.com was successful and the server responded with an HTTP status code of 200 OK, which means the request was successful.

## Replicating the Checks Using Standard Network Tools

To replicate the checks performed by the Go program using standard network tools, you can use a combination of commands like `nslookup` or `dig` for DNS resolution, `telnet` or `nc` (netcat) for TCP connectivity, and `curl` for HTTP status checking. Here’s how you can perform each check:

### DNS Resolution

To resolve the domain to its IP addresses:
```sh
nslookup google.com
```
or
```sh
dig google.com +short
```

### TCP Connectivity

To check TCP connectivity to port 80:
```sh
telnet google.com 80
```
or, using `nc` (netcat):
```sh
nc -zv google.com 80
```

### Measure Response Time

Measuring the exact response time for TCP connectivity using simple network tools can be tricky, but you can approximate it using `time` with `nc`:
```sh
time nc -zv google.com 80
```

### HTTP Response

To check the HTTP status:
```sh
curl -I http://google.com
```

## Brief Introduction of `cli.App`

The `cli.App` struct and its method in Go are part of the `urfave/cli` package, a popular library for building command-line interfaces (CLI) in Go. The `cli.App` struct is the central part of this package, representing an application with commands, flags, and other configuration options.

### `cli.App` Struct

The `cli.App` struct defines an application with the following fields:

1. **Name**: A string representing the name of the application.
2. **Usage**: A string describing the application, often used to provide a brief summary of what the application does.
3. **Version**: A string specifying the version of the application.
4. **Commands**: A slice of `cli.Command` structs, each representing a command that the application can execute.
5. **Flags**: A slice of `cli.Flag` interfaces, representing global flags for the application.
6. **Action**: A `cli.ActionFunc` (or function) to be executed if no command is specified.
7. **Before**: A `cli.BeforeFunc` (or function) to be executed before any other actions.
8. **After**: A `cli.AfterFunc` (or function) to be executed after all actions.
9. **Authors**: A slice of `cli.Author` structs, representing the authors of the application.
10. **HelpName**: A string that can override the help message name for the application.
11. **UsageText**: A string that can provide more detailed usage text for the application.
12. **Description**: A string that can offer a detailed description of the application.
13. **Compiled**: A `time.Time` value indicating when the application was compiled.
14. **Copyright**: A string for copyright information.
15. **ArgsUsage**: A string specifying how to use the command-line arguments.
16. **CommandNotFound**: A function to be executed if a specified command is not found.
17. **OnUsageError**: A function to be executed if there is an error with the usage.
18. **EnableBashCompletion**: A boolean enabling bash completion.
19. **HideHelp**: A boolean to hide the help command.
20. **HideVersion**: A boolean to hide the version command.
```

![Screenshot](images/screenshot.png)
```

Ensure that the image is in the correct path relative to your `README.md` file.
