<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Comamoca/recker?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Comamoca/recker?style=flat-square)
![Issues](https://img.shields.io/github/issues/Comamoca/recker?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Comamoca/recker?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Comamoca/recker/bug?style=flat-square)

# π¦ recker

The repository picker

</div>

<table>
  <thead>
    <tr>
      <th style="text-align:center">πEnglish</th>
      <th style="text-align:center"><a href="README.ja.md">π‘ζ₯ζ¬θͺ</a></th>
    </tr>
  </thead>
</table>

<div align="center">

</div>

## π How to use

> **Note**
> New Feature!!
> if you select `sandbox` repository, recker launch with
> fuzzy-finder in `sandbox` repository again.

require [ghq](https://github.com/x-motemen/ghq)

```
# out put selected repo path
recker

# move to selected repo path
cd $(recker)
```

## β¬οΈ Install

build binary

```sh
go build
./recker
```

and move to on your path.

## βοΈ Development

```sh
go run .
```

## π Todo

- [ ] fix TODOs

## π License

MIT

### π§© Modules

- [go-fuzzyfinder](http://github.com/ktr0731/go-fuzzyfinder)

## π Special Thanks

- [fzf](https://github.com/junegunn/fzf)
