# Neovimプラグインの構造

## ディレクトリ構成

```
myplugin/
├── lua/
│   └── myplugin/
│       └── init.lua      -- require("myplugin") で読み込まれる
├── plugin/
│   └── myplugin.lua      -- Neovim起動時に自動実行される
└── README.md
```

## lua/ ディレクトリ

`require()` で読み込むモジュールを置く。

| require | 読み込まれるファイル |
|---------|---------------------|
| `require("myplugin")` | `lua/myplugin/init.lua` または `lua/myplugin.lua` |
| `require("myplugin.utils")` | `lua/myplugin/utils.lua` |
| `require("myplugin.core.parser")` | `lua/myplugin/core/parser.lua` |

## plugin/ ディレクトリ

- Neovim起動時に**自動で実行される**
- コマンドやキーマップの登録に使う
- 重い処理は書かない（起動が遅くなる）

## 開発時のテスト方法

```vim
:luafile %          " 現在のファイルを実行
:lua print("test")  " Luaコードを直接実行
:lua =vim.opt.number:get()  " = で値を表示（Neovim 0.9+）
```

## runtimepath

プラグインが読み込まれるには `runtimepath` に追加されている必要がある。

```lua
-- 確認
:lua =vim.opt.runtimepath:get()

-- 手動で追加（開発時）
vim.opt.runtimepath:append("/path/to/myplugin")
```

パッケージマネージャ（lazy.nvim等）を使えば自動で管理される。
