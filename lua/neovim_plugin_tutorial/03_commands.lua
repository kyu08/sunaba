-- コマンドの作成

-- ========================================
-- vim.api.nvim_create_user_command
-- ========================================
-- :Hello で呼び出せるコマンドを作成
-- NOTE: :Hello foo barで呼び出したとき、optsには"foo bar"が入る。
vim.api.nvim_create_user_command("Hello", function(opts)
	print("Hello, " .. (opts.args or "World"))
end, {
	nargs = "?",
	desc = "Say hello",
})

-- 基本形

-- opts の中身
-- opts.args   : 引数全体の文字列
-- opts.fargs  : 引数をスペースで分割した配列
-- opts.bang   : ! が付いたか (bool)
-- opts.range  : 範囲指定された場合 (line1, line2)

-- nargs オプション
-- "0" : 引数なし
-- "1" : 引数1つ（必須）
-- "?" : 引数0か1
-- "*" : 引数0以上
-- "+" : 引数1以上

-- 補完を追加
vim.api.nvim_create_user_command("SetColor", function(opts)
	vim.opt.background = opts.args
end, {
	nargs = 1,
	complete = function()
		return { "dark", "light" }
	end,
})

-- バッファローカルなコマンド
-- vim.api.nvim_buf_create_user_command(0, "BufHello", ...)

--[[
練習問題（Neovim内で :luafile % で実行）:
1. :Greet コマンドを作成し、引数の名前で挨拶を表示してください
2. opts.fargs を使って複数の引数を受け取るコマンドを作ってください
3. complete オプションで補完候補を追加してください
4. :Greet! で ! 付きの場合に挙動を変えてみてください
]]
