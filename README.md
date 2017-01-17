# dotmgmt

The stupid-simple way to manage your dotfiles.

## Use cases

Say you have a `dotfiles` directory in your home folder, containing folders and/or files that you want symlinked
in your home directory:

```
dotmgmt
```

Done. Your files have been symlinked.

A record of it is then kept under `~/.dotmgmt/mgmt` in `json` format.

This file is updated everytime `dotmgmt` adds or removes a symlink, so you always have an up-to-date record of which files you are managing.



Say you want to remove the files you have just symlinked:

```
dotmgmt -reset
```

Done. Your symlinks have been removed, and the original files have not been touched.


## USAGE:

- Remove a specific symlink:
```
dotmgmt -del ~/symlink/path
```

- Create a symlink for a specific file:
```
dotmgmt -make /path/to/dir
```

- Path to the folder you want to create symlinks from:
```
dotmgmt -p /path/to/dir (default ".")
```

- Remove all the symlinks for the files in the current directory:
```
dotmgmt -reset
```
