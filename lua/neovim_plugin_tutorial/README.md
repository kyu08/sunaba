# Neovimプラグイン基礎チュートリアル

簡単なコマンド・キーマップを作れるようになることを目標としたチュートリアル。

## 構成

| ファイル | 内容 |
|----------|------|
| `01_plugin_structure.md` | プラグインのディレクトリ構成、`require`の仕組み |
| `02_vim_api/init.lua` | `vim.api`, `vim.fn`, `vim.opt`, `vim.g` |
| `03_commands.lua` | `nvim_create_user_command` でコマンド作成 |
| `04_keymaps.lua` | `vim.keymap.set` でキーマップ作成 |
| `05_autocmd.lua` | `nvim_create_autocmd` でイベント処理 |

## 進め方

1. 各ファイルを順番に開く
2. 解説を読む
3. Neovim内で `:luafile %` を実行して動作確認
4. 練習問題を解く

## 前提

- Lua基礎文法（`lua_basics/` で学習済み）
- Neovimの基本操作
