# aicommit

使用 AI 自動生成 Git commit message 的 CLI 工具。支援 OpenAI 和 Anthropic API。

## 功能特色

- 自動生成 [Conventional Commits](https://www.conventionalcommits.org/) 格式的 commit message
- 支援 OpenAI 和 Anthropic 兩種 AI 提供者
- 支援多國語言 commit message（中文、日文、韓文等）
- 簡單的命令列介面
- 安全的設定檔儲存

## Installation 安裝

### 一鍵安裝（推薦）

複製以下指令到終端機執行，自動下載預編譯的執行檔並安裝：

**macOS (Apple Silicon M1/M2/M3):**
```bash
curl -sSL https://github.com/meowalien/aicommit/releases/latest/download/aicommit-darwin-arm64 -o /usr/local/bin/aicommit && \
chmod +x /usr/local/bin/aicommit && \
echo "✅ 安裝完成！" && \
aicommit --help
```

**macOS (Intel):**
```bash
curl -sSL https://github.com/meowalien/aicommit/releases/latest/download/aicommit-darwin-amd64 -o /usr/local/bin/aicommit && \
chmod +x /usr/local/bin/aicommit && \
echo "✅ 安裝完成！" && \
aicommit --help
```

**Linux (x86_64):**
```bash
curl -sSL https://github.com/meowalien/aicommit/releases/latest/download/aicommit-linux-amd64 -o /usr/local/bin/aicommit && \
chmod +x /usr/local/bin/aicommit && \
echo "✅ 安裝完成！" && \
aicommit --help
```

**Linux (ARM64):**
```bash
curl -sSL https://github.com/meowalien/aicommit/releases/latest/download/aicommit-linux-arm64 -o /usr/local/bin/aicommit && \
chmod +x /usr/local/bin/aicommit && \
echo "✅ 安裝完成！" && \
aicommit --help
```

> 如果 `/usr/local/bin` 需要權限，請在指令前加上 `sudo`

### 手動下載安裝

1. 前往 [Releases](https://github.com/meowalien/aicommit/releases) 頁面
2. 下載對應你系統的執行檔：
   - `aicommit-darwin-arm64` - macOS Apple Silicon
   - `aicommit-darwin-amd64` - macOS Intel
   - `aicommit-linux-amd64` - Linux x86_64
   - `aicommit-linux-arm64` - Linux ARM64
3. 重新命名為 `aicommit` 並移動到 PATH 目錄中
4. 賦予執行權限：`chmod +x aicommit`

### 從原始碼編譯安裝

如果你想從原始碼編譯，需要先安裝 Go 1.21 或更高版本。

```bash
# 1. Clone 專案
git clone https://github.com/meowalien/aicommit.git
cd aicommit

# 2. 編譯並安裝到 GOPATH/bin
go install ./cmd/aicommit/

# 3. 設定 PATH（加入 ~/.zshrc）
grep -q 'export PATH="$PATH:$HOME/go/bin"' ~/.zshrc || echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.zshrc
source ~/.zshrc

# 4. 驗證安裝
aicommit --help
```

## 設定 API Key

在使用之前，需要先設定 AI 提供者的 API Key：

### 使用 Anthropic（預設）

```bash
aicommit set anthropic_key=你的_ANTHROPIC_API_KEY
aicommit set provider=anthropic
```

### 使用 OpenAI

```bash
aicommit set openai_key=你的_OPENAI_API_KEY
aicommit set provider=openai
```

### 可選：指定模型

```bash
# Anthropic 模型
aicommit set anthropic_model=claude-sonnet-4-20250514

# OpenAI 模型
aicommit set openai_model=gpt-4o
```

### 設定 Commit Message 語言

預設為英文 (`en`)，可設定為其他語言：

```bash
# 繁體中文
aicommit set language=zh-TW

# 簡體中文
aicommit set language=zh-CN

# 日文
aicommit set language=ja

# 韓文
aicommit set language=ko

# 英文（預設）
aicommit set language=en
```

也可以使用簡寫 `lang`：

```bash
aicommit set lang=zh-TW
```

設定檔儲存於 `~/.aicommit/config.yaml`

## 使用方式

### 基本用法

```bash
# 1. 先 stage 你的修改
git add .

# 2. 執行 aicommit 自動生成 commit message 並 commit
aicommit
```

### 進階選項

```bash
# 預覽模式：只顯示生成的 message，不執行 commit
aicommit --dry-run
aicommit -d

# 詳細輸出模式
aicommit --verbose
aicommit -v

# 顯示說明
aicommit --help
```

## 指令說明

| 指令 | 說明 |
|------|------|
| `aicommit` | 自動生成 commit message 並 commit |
| `aicommit set KEY=VALUE` | 設定 API key 或其他選項 |
| `aicommit --dry-run` | 預覽模式，不執行 commit |
| `aicommit --verbose` | 詳細輸出模式 |
| `aicommit --help` | 顯示說明 |

## 設定選項

| Key | 說明 | 預設值 |
|-----|------|--------|
| `provider` | AI 提供者 (`openai` 或 `anthropic`) | `anthropic` |
| `language` / `lang` | Commit message 語言 | `en` |
| `anthropic_key` | Anthropic API Key | - |
| `anthropic_model` | Anthropic 模型 | `claude-sonnet-4-20250514` |
| `openai_key` | OpenAI API Key | - |
| `openai_model` | OpenAI 模型 | `gpt-4o` |

## 使用範例

```bash
$ git add .
$ aicommit
Generating commit message using anthropic...

Generated commit message:
feat(auth): add user login validation

Commit successful!
```

```bash
$ aicommit --dry-run
Generating commit message using anthropic...

Generated commit message:
fix(api): resolve null pointer exception in user service

(dry-run mode - not committing)
```

### 中文 Commit Message 範例

```bash
$ aicommit set language=zh-TW
$ aicommit --dry-run
Generating commit message using anthropic...

Generated commit message:
feat(auth): 新增使用者登入驗證功能

(dry-run mode - not committing)
```

## 生成的 Commit Message 格式

工具會自動生成符合 [Conventional Commits](https://www.conventionalcommits.org/) 規範的 commit message：

- `feat`: 新功能
- `fix`: 修復 bug
- `docs`: 文件修改
- `style`: 程式碼格式修改（不影響功能）
- `refactor`: 重構
- `perf`: 效能優化
- `test`: 測試相關
- `build`: 建置系統或外部依賴修改
- `ci`: CI 設定修改
- `chore`: 其他雜項修改

## 故障排除

### 找不到 aicommit 指令

確保 `$HOME/go/bin` 在你的 PATH 中：

```bash
echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.zshrc
source ~/.zshrc
```

### API Key 錯誤

確認你的 API Key 設定正確：

```bash
cat ~/.aicommit/config.yaml
```

### 沒有 staged 的檔案

在執行 `aicommit` 之前，需要先 stage 檔案：

```bash
git add .
# 或
git add <specific-files>
```

## 解除安裝

```bash
# 刪除執行檔
rm $(which aicommit)

# 刪除設定檔
rm -rf ~/.aicommit
```

## License

MIT
