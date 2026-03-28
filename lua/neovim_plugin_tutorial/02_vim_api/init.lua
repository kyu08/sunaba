-- 基本的なVim API

-- ========================================
-- vim.api - Neovim API (nvim_xxx関数)
-- ========================================
-- 低レベルなNeovim操作

-- 現在のバッファの全行を取得
local current_buf_lines = vim.api.nvim_buf_get_lines(0, 0, -1, false)
for i, line in ipairs(current_buf_lines) do
	print("[" .. i .. "]: " .. line)
end

-- 現在行を変更
vim.api.nvim_set_current_line("new text")

-- メッセージを表示
vim.api.nvim_echo({ { "Hello", "WarningMsg" } }, true, {})

-- ========================================
-- vim.fn - Vimscript関数を呼ぶ
-- ========================================
-- Vimscriptの関数をそのまま使える

-- ========================================
-- vim.opt - オプション設定
-- ========================================
-- :set xxx に相当

-- ========================================
-- vim.g / vim.b / vim.w - 変数
-- ========================================
-- vim.g.xxx  -- グローバル変数 (let g:xxx)
-- vim.b.xxx  -- バッファローカル (let b:xxx)
-- vim.w.xxx  -- ウィンドウローカル (let w:xxx)

--[[
練習問題（Neovim内で :luafile % で実行）:
1. vim.fn.expand("%:p") で現在ファイルのパスを表示してください
2. vim.opt.number = true で行番号を表示してください
3. vim.api.nvim_echo で色付きメッセージを表示してください
4. vim.g.mapleader を確認・変更してみてください
]]
