# gitmon
Monitor git repos for corresponding upstream changes

### Requirements:
 - `git` installed locally on the system. Project uses raw `git` commands to get the latest updates

### Running:

The project binary can be installed as a Windows service using something like `nssm` or windows command `installutil` or
`SC` Ref: [StackOverflow](https://stackoverflow.com/questions/8164859/install-a-windows-service-using-a-windows-command-prompt). 

On linux `systemd` can be used to package it and then run as a background service. I do not have a Linux machine so the
support is not added out of the box

#### Note:
When running as a `service`, program should run under the user which as access to the `git project`


### Standalone local run
`gitmon` can be run directly in `terminal` as well.


### Build locally

Requirements:
- go 1.20
- Make

```bash
make all
```
