-- モジュールの定義
-- テーブルを作成し、関数を追加して、最後にreturnする

local M = {}

-- プライベート関数（外部からアクセスできない）

-- パブリック関数（M.xxx で外部に公開）
function M.greet(name)
	return "Hello, " .. name
end

return M
