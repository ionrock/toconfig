Toconfig
========

Toconfig takes variables in your environment and plugs them into a
config file template before running a process.

For example here is a simple template.

```yaml
Hello: {{ Get "USER" }}
```

And here is a command that uses the template to write the config.

```
$ toconfig --template example.conf.tmpl --config example.conf cat example.conf
Hello: eric
```

Why?
====

When your goal is (12 Factor)[http://12factor.net/], yet you have many
apps that don't accept environment variables as configuration, you can
use toconfig to help bridge that gap. Also, because `toconfig` runs
when the process starts and exits when the process exits, it is well
suited for containers where the environment is passed to something
like the docker runtime.
