" neovimで読み込まれたく無いのでファイル名を変えている
set nocompatible

"-------------------------------------------------------------------------------
" Settings
"-------------------------------------------------------------------------------
"文字コードをUFT-8に設定
set fenc=utf-8
set encoding=utf-8
set nospell
set mouse=n
map <LeftMouse> <Nop>

" 読み込みのタイミング的に？init.vimの途中移行に書かないとcmdheight指定が効かないので
" init.vimに書いている(プラグインが上書きしてるとかかもしれない)
" set cmdheight=1
" バックアップファイルを作らない
set nobackup
" スワップファイルを作らない
set noswapfile
" 編集中のファイルが変更されたら自動で読み直す
set autoread
" バッファが編集中でもその他のファイルを開けるように
set hidden
" 入力中のコマンドをステータスに表示する
set showcmd
" クリップボードをOSと共有
set clipboard=unnamed 
set nocursorline
set ttimeoutlen=50
" 上下n行の視界を確保
set scrolloff=10
" solarizedをちゃんと表示するために必要
set termguicolors 
" Backspaceキーの影響範囲に制限を設けない
set backspace=indent,eol,start 
" コマンドラインモードでTABキーによるファイル名補完を有効にする
set wildmenu wildmode=list:longest,full
" カーソルの左右移動で行末から次の行の行頭への移動が可能になる
set whichwrap=b,s,h,l,<,>,[,],~ 
" timeout を無効に(https://stackoverflow.com/questions/14737429/how-to-disable-the-timeout-on-the-vim-leader-key)
set notimeout nottimeout
" vimdiff を左右分割に
set diffopt=vertical
highlight DiffAdd gui=NONE guibg=#0020a0 guifg=#a0d0ff
highlight DiffChange gui=NONE guibg=#401830 guifg=#e03870
highlight DiffDelete gui=NONE guibg=#0020a0 guifg=#a0d0ff
highlight DiffText gui=NONE guibg=#802860 guifg=#ff78f0

"-------------------------------------------------------------------------------
" 折りたたみ
"-------------------------------------------------------------------------------
"インデント単位で折り畳む
set foldmethod=indent
" デフォルトで折り畳みOFF
set nofoldenable

"-------------------------------------------------------------------------------
" 見た目系
"-------------------------------------------------------------------------------
" 行番号を表示
set number
" 行末の1文字先までカーソルを移動できるように
set virtualedit=onemore
" スマートインデント
set smartindent
" 括弧入力時の対応する括弧を表示
set showmatch
set matchtime=3
" コマンドラインの補完
set wildmode=list:longest
" 折り返し時に表示行単位での移動できるようにする
nnoremap j gj
nnoremap k gk
" floating window を透明にする
" set winblend=10

if has('vim_starting')
    " 挿入モード時に非点滅の縦棒タイプのカーソル
    let &t_SI .= "\e[6 q"
    " ノーマルモード時に非点滅のブロックタイプのカーソル
    let &t_EI .= "\e[2 q"
    " 置換モード時に非点滅の下線タイプのカーソル
    let &t_SR .= "\e[4 q"
endif

"-------------------------------------------------------------------------------
" Tab系
"-------------------------------------------------------------------------------
" 不可視文字を可視化(タブが「▸-」と表示される)
set list listchars=tab:\▸\-
" Tab文字を半角スペースにする
set expandtab
" 行頭以外のTab文字の表示幅（スペースいくつ分）
set tabstop=2
" 行頭でのTab文字の表示幅
set shiftwidth=2

"-------------------------------------------------------------------------------
" 検索系
"-------------------------------------------------------------------------------
" 検索文字列が小文字の場合は大文字小文字を区別なく検索する
set ignorecase
" 検索文字列に大文字が含まれている場合は区別して検索する
set smartcase
" 検索文字列入力時に順次対象文字列にヒットさせる
set incsearch
" 検索時に最後まで行ったら最初に戻る
" set wrapscan
" 検索語をハイライト表示
set hlsearch
" ESC連打でハイライト解除
nmap <Esc><Esc> :nohlsearch<CR><Esc>
" 右端で行を折り返さない
" set nowrap

"-------------------------------------------------------------------------------
" Alias
"-------------------------------------------------------------------------------
 " <Leader> を <Space> に
let mapleader = "\<Space>"
 " jjで INSERT MODE からでる + 保存
inoremap <silent> jj <ESC>:w<CR>
inoremap <silent> <ESC> <ESC>:w<CR>
" 保存
nnoremap <Leader>w :w<CR>
" バッファを全て保存
nnoremap <Leader><Leader>w :wa<CR>
" 左右に分割
nnoremap <Leader>v <C-w>v
" 左のタブに移動
nnoremap <Leader>h <C-w>h
" 右のタブに移動
nnoremap <Leader>l <C-w>l
" 下のタブに移動
nnoremap <Leader>j <C-w>j
" 上のタブに移動
nnoremap <Leader>k <C-w>k
" git 画面表示
" nnoremap <Leader><Leader>gg :Gstatus<CR>
" git push
" nnoremap <Leader><Leader>gp :Gpush -u origin HEAD<CR>
" git log
" nnoremap <Leader><Leader>gl :Glog<CR>
" 段落移動を remap
nnoremap <C-j> }
nnoremap <C-k> {
vnoremap <C-j> }
vnoremap <C-k> {

nnoremap <Leader>c cgn
" 補完候補選択をC-yにremap
inoremap <C-o> <C-y>
nnoremap <Leader>z ZZ
nnoremap <Leader><Leader>f :Format<CR>
nnoremap /  /\v
nnoremap <Leader><Leader>z :wa<CR><C-w><C-o>ZZ
nnoremap <Leader><Leader>c :let @* = expand('%:p')[56:-1]<CR>
tnoremap <c-\> <c-\><c-n>
nnoremap <c-t> <Nop>
nnoremap gh :Commentary<CR>
vnoremap gh :Commentary<CR>
vnoremap [ <
vnoremap ] >

inoremap <C-a> <C-o>^
inoremap <C-e> <C-o>$
inoremap <C-f> <C-o>w
inoremap <C-b> <C-o>b
inoremap <C-d> <C-o>x
inoremap <C-j> <C-o>j
inoremap <C-l> <C-o>l

nnoremap <Leader><Leader>r :source ~/.config/nvim/init.vim<CR>

set laststatus=2

nnoremap <Leader>y y$
nnoremap <C-h> /\v
nnoremap x "zx
vnoremap x "zx
vnoremap v "zx
" 置換
nnoremap <Leader>x :%s/\v//g<Left><Left><Left>

nnoremap <c-q> <c-z>

" VSCode Vim Pluginだとなんか挙動がおかしいのでsettings.json側で設定する
" nnoremap <c-y> 4<c-y>
" nnoremap <c-e> 4<c-e>
nnoremap <Leader><Leader><Leader><Leader>g :!gofumpt -w %:p<CR>

" netrwをスペース+fで開く/閉じるスクリプト

" netrwの設定
let g:netrw_banner = 0        " バナーを非表示
let g:netrw_liststyle = 3     " ツリー表示
let g:netrw_browse_split = 4  " 前のウィンドウで開く
let g:netrw_altv = 1          " 右側に開く
let g:netrw_winsize = 25      " ウィンドウサイズ

" netrwが開いているかどうかを追跡する変数
let g:netrw_is_open = 0

" netrwのトグル関数
function! ToggleNetrw()
    if g:netrw_is_open
        " netrwを閉じる
        for i in range(1, bufnr('$'))
            if buflisted(i)
                if getbufvar(i, '&filetype') == "netrw"
                    silent exe 'bdelete ' . i
                endif
            endif
        endfor
        let g:netrw_is_open = 0
    else
        " netrwを開く
        let g:netrw_is_open = 1
        silent Lexplore
    endif
endfunction

" netrw用のキーマッピング設定
augroup netrw_mapping
    autocmd!
    autocmd filetype netrw call NetrwMapping()
augroup END

function! NetrwMapping()
    " netrwバッファでqを押すとネットワを閉じる
    nmap <buffer> q :call ToggleNetrw()<CR>
endfunction

" キーマッピング
nnoremap <silent> <Space>f :call ToggleNetrw()<CR>
