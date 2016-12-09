# dotmgmt

The stupid-simple way to manage your dotfiles.

## Use cases

Say you have a `dotfiles` directory in your home folder, containing folders and/or files that you want symlinked
in your home directory:

```
dotmgmt
```

Done. Your files have been symlinked.

Say you want to remove the files you have just symlinked:

```
dotmgmt -park
```

Done. Your symlinks have been removed, and the original files have not been touched.


![https://i.imgur.com/CFaiApr.png](https://i.imgur.com/CFaiApr.png)
