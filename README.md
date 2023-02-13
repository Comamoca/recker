<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Comamoca/recker?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Comamoca/recker?style=flat-square)
![Issues](https://img.shields.io/github/issues/Comamoca/recker?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Comamoca/recker?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Comamoca/recker/bug?style=flat-square)

# 🦊 recker

The repository picker

</div>

<table>
  <thead>
    <tr>
      <th style="text-align:center">🍔English</th>
      <th style="text-align:center"><a href="README.ja.md">🍡日本語</a></th>
    </tr>
  </thead>
</table>

<div align="center">

</div>

## 🚀 How to use

> **New Feature!**
> if you select `sandbox` repository, recker launch with
> fuzzy-finder in `sandbox` repository again.

require [ghq](https://github.com/x-motemen/ghq)

```
# out put selected repo path
recker

# move to selected repo path
cd $(recker)
```

## ⬇️ Install

build binary

```sh
go build
./recker
```

and move to on your path.

## ⛏️ Development

```sh
go run .
```

## 📝 Todo

- [ ] fix TODOs

## 📜 License

MIT

### 🧩 Modules

- [go-fuzzyfinder](http://github.com/ktr0731/go-fuzzyfinder)

## 💕 Special Thanks

- [fzf](https://github.com/junegunn/fzf)
