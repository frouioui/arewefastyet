## arewefastyet exec

Execute a task

```
arewefastyet exec [flags]
```

### Options

```
      --ansible-inventory-files strings   List of inventory files used by Ansible
      --ansible-playbook-files strings    List of playbook files used by Ansible
      --ansible-root-directory string     Root directory of Ansible
      --db-database string                Database to use.
      --db-host string                    Hostname of the database
      --db-password string                Password to authenticate the database.
      --db-user string                    User used to connect to the database
      --equinix-instance-type string      Instance type to use for the creation of a new node
      --equinix-project-id string         Project ID to use for Equinix Metal
      --equinix-token string              Auth Token for Equinix Metal
      --exec-git-ref string               Git reference on which the benchmarks will run
      --exec-root-dir string              Path to the root directory of exec
      --exec-source string                Name of the source that triggered the execution
  -h, --help                              help for exec
      --infra-path string                 Path to the infra directory
      --stats-remote-db-database string   Name of the stats remote database.
      --stats-remote-db-host string       Hostname of the stats remote database.
      --stats-remote-db-password string   Password to authenticate the stats remote database.
      --stats-remote-db-port string       Port of the stats remote database.
      --stats-remote-db-user string       User used to connect to the stats remote database
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/arewefastyet/config.yaml)
```

### SEE ALSO

* [arewefastyet](arewefastyet.md)	 - 

