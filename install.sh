#!/bin/bash

# aicommit 一鍵安裝腳本
# 使用方式: curl -sSL https://raw.githubusercontent.com/jacky_li/aicommit/main/install.sh | bash

set -e

echo "🚀 開始安裝 aicommit..."

# 檢查 Go 是否已安裝
if ! command -v go &> /dev/null; then
    echo "❌ 錯誤: 請先安裝 Go (https://golang.org/dl/)"
    exit 1
fi

# 檢查 Git 是否已安裝
if ! command -v git &> /dev/null; then
    echo "❌ 錯誤: 請先安裝 Git"
    exit 1
fi

# 建立暫存目錄
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

echo "📥 下載 aicommit..."
git clone --depth 1 https://github.com/jacky_li/aicommit.git .

echo "🔨 編譯安裝..."
go install ./cmd/aicommit/

# 偵測 shell 設定檔
SHELL_RC=""
if [ -n "$ZSH_VERSION" ] || [ "$SHELL" = "/bin/zsh" ]; then
    SHELL_RC="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ] || [ "$SHELL" = "/bin/bash" ]; then
    if [ -f "$HOME/.bash_profile" ]; then
        SHELL_RC="$HOME/.bash_profile"
    else
        SHELL_RC="$HOME/.bashrc"
    fi
fi

# 加入 PATH
EXPORT_LINE='export PATH="$PATH:$HOME/go/bin"'
if [ -n "$SHELL_RC" ]; then
    if ! grep -q "$HOME/go/bin" "$SHELL_RC" 2>/dev/null; then
        echo "" >> "$SHELL_RC"
        echo "# Added by aicommit installer" >> "$SHELL_RC"
        echo "$EXPORT_LINE" >> "$SHELL_RC"
        echo "📝 已將 PATH 設定加入 $SHELL_RC"
    else
        echo "✅ PATH 已設定"
    fi
fi

# 清理暫存目錄
cd /
rm -rf "$TEMP_DIR"

echo ""
echo "✅ aicommit 安裝完成！"
echo ""
echo "📋 下一步："
echo "   1. 重新開啟終端機，或執行: source $SHELL_RC"
echo "   2. 設定 API Key:"
echo "      aicommit set anthropic_key=你的_API_KEY"
echo "      aicommit set provider=anthropic"
echo ""
echo "   或使用 OpenAI:"
echo "      aicommit set openai_key=你的_API_KEY"
echo "      aicommit set provider=openai"
echo ""
echo "🎉 使用方式: git add . && aicommit"
