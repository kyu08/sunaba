#!/bin/bash
set -e

# リポジトリとカテゴリをスペース区切りで定義
# フォーマット: "リポジトリ名 カテゴリ"
REPOS=(
  "bbf-kubernetes go"
  "go-system-programming computer_science"
  "go-static-analysis go"
  "codecrafters-git-go go"
  "codecrafters-redis-go go"
  "jitsuyo-golang go"
  "distributed-services-with-go go"
  "tinet network"
  "TS-official typescript"
  "tour-of-rust rust"
  "s99 scala"
  "cop-bon scala"
  "play-api-sample scala"
  "play-next-beat-2 scala"
  "play-next-beat-1 scala"
  "play-hands-on-by-bizreach scala"
  "play-scala-text scala"
  "scala-qiita scala"
  "scala-text scala"
  "tour-of-scala scala"
  "python-playground python"
  "cs50 computer_science"
  "c-playground c"
  "data-structure-and-algorithm computer_science"
  "kisokara-elm elm"
  "twitter-clone-api webapp"
  "twitter-clone-ui webapp"
  "crwn-clothing react"
  "monsters-rolodex react"
  "git-study git"
  "devchallenges-weather-app react"
  "nand2tetris rust"
  "rust-playground_ rust"
  "syntect-tui-sample rust"
  "rust-json-decode-with-line-number rust"
  "learning-rust-from-concepts rust"
  "the-rust-programming-language rust"
  "minigrep rust"
  "rusty-journal rust"
)

OWNER="kyu08"
UNIFIED_REPO="sunaba"

# エラーが発生したリポジトリを記録
FAILED_REPOS=()

# 各リポジトリを統合
for entry in "${REPOS[@]}"; do
  repo=$(echo "$entry" | cut -d' ' -f1)
  category=$(echo "$entry" | cut -d' ' -f2)
  
  echo "Processing $repo..."
  
  subdir="$category/$repo"
  
  # エラーが発生しても続行
  if ! (
    set -e
    # remoteを追加してfetch
    git remote add "$repo" "https://github.com/$OWNER/$repo.git" 2>/dev/null || true
    git fetch "$repo"
    
    # デフォルトブランチを取得（main or master）
    default_branch=$(git remote show "$repo" | grep 'HEAD branch' | cut -d' ' -f5)
    
    # サブディレクトリにマージ（履歴を保持）
    git subtree add --prefix="$subdir" "$repo/$default_branch" --message "Add $repo to $subdir"
    
    # remote削除
    git remote remove "$repo"
  ); then
    echo "✗ Failed to merge $repo"
    FAILED_REPOS+=("$repo")
    # エラー時もremoteをクリーンアップ
    git remote remove "$repo" 2>/dev/null || true
    git merge --abort 2>/dev/null || true
    continue
  fi
  
  echo "✓ $repo merged into $subdir/"
done

echo ""
echo "======================================"
if [ ${#FAILED_REPOS[@]} -eq 0 ]; then
  echo "All repositories merged successfully!"
else
  echo "Failed repositories:"
  for failed_repo in "${FAILED_REPOS[@]}"; do
    echo "  - $failed_repo"
  done
  exit 1
fi
