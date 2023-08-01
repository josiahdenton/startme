# startme

```
Use at your own risk, this is still very new
```

`startme` is a CLI tool for saving templates or snippets of code for later use. The tool supports

- adding a template 
- searching for a template and copying to system clipboard

and currently only supposts `nvim` as the default editor.


## Usage

Once you have `startme` on your path, you can run

- For help
```bashrc
startme -h
```

- For adding a new template
```
startme -n <title> -e <extension> # -e is optional and defaults to md (markdown)
```

- For searching existing templates
```
startme 
```

## Future Work

I would like to be able to...
- delete/edit existing templates
- have more flexible cmd options
- make the CLI more "Unix" like
- set a global config to set the defualt editor (i.e. nvim, vim, vi, emacs, nano)
- improve the install process by hosting the binary
    - possibly use [goreleaser](https://goreleaser.com/)
