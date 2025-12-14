#!/bin/bash

# aicommit 一鍵安裝腳本
# 使用方式: curl -sSL https://raw.githubusercontent.com/meowalien/aicommit/main/install.sh | bash

set -e

echo "🚀 開始安裝 aicommit..."

# 偵測系統架構
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
    darwin)
        OS="darwin"
        ;;
    linux)
        OS="linux"
        ;;
    *)
        echo "❌ 不支援的作業系統: $OS"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64|amd64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "❌ 不支援的系統架構: $ARCH"
        exit 1
        ;;
esac

BINARY_NAME="aicommit-${OS}-${ARCH}"
DOWNLOAD_URL="https://github.com/meowalien/aicommit/releases/latest/download/${BINARY_NAME}"
INSTALL_DIR="/usr/local/bin"
INSTALL_PATH="${INSTALL_DIR}/aicommit"

echo "📥 下載 ${BINARY_NAME}..."
echo "   URL: ${DOWNLOAD_URL}"

# 檢查是否需要 sudo
if [ -w "$INSTALL_DIR" ]; then
    curl -sSL "$DOWNLOAD_URL" -o "$INSTALL_PATH"
    chmod +x "$INSTALL_PATH"
else
    echo "⚠️  需要管理員權限安裝到 ${INSTALL_DIR}"
    sudo curl -sSL "$DOWNLOAD_URL" -o "$INSTALL_PATH"
    sudo chmod +x "$INSTALL_PATH"
fi

# 驗證安裝
if command -v aicommit &> /dev/null; then
    echo ""
    echo "✅ aicommit 安裝完成！"
    echo ""
    echo "📋 下一步："
    echo "   1. 設定 API Key:"
    echo "      aicommit set anthropic_key=你的_API_KEY"
    echo "      aicommit set provider=anthropic"
    echo ""
    echo "   或使用 OpenAI:"
    echo "      aicommit set openai_key=你的_API_KEY"
    echo "      aicommit set provider=openai"
    echo ""
    echo "   2. 設定語言（可選）:"
    echo "      aicommit set language=zh-TW"
    echo ""
    echo "🎉 使用方式: git add . && aicommit"
else
    echo "❌ 安裝失敗，請手動下載安裝"
    exit 1
fi
