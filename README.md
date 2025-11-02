# 🍺 Brew Manager

Homebrew 桌面管理工具，提供直观的图形界面来管理 macOS 上的 Homebrew 软件包。

## ✨ 功能特性

- 📦 **查看已安装软件**：一目了然地查看所有通过 Homebrew 安装的软件包
- 🔍 **搜索软件包**：快速搜索 Homebrew 仓库中的软件包
- ⬇️ **安装软件**：一键安装 Formulae 或 Cask 应用
- 🗑️ **卸载软件**：轻松卸载不需要的软件包
- 🔄 **重新安装**：重装有问题的软件包
- ⬆️ **更新软件**：单个或批量更新软件包
- 🧹 **清理缓存**：清理旧版本和缓存文件
- 📊 **可更新列表**：查看所有可以更新的软件包


## 🚀 快速开始

### 开发环境设置

1. **克隆或下载项目**

2. **安装 Wails CLI**
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

3. **安装依赖**
```bash
# 安装 Go 依赖
go mod download

# 安装前端依赖
cd frontend
npm install
cd ..
```

4. **运行开发模式**
```bash
wails dev
```

### 构建应用

```bash
# 构建生产版本
wails build

# 构建后的应用位于 build/bin/ 目录
```

## 📖 使用说明

### 已安装软件

- 查看所有已安装的 Formulae 和 Cask 应用
- 使用搜索框快速过滤
- 对每个软件包可以：
  - 查看详细信息
  - 重新安装
  - 卸载

### 搜索软件

- 输入关键词搜索 Homebrew 仓库
- 查看搜索结果并直接安装
- 显示软件是否已安装

### 可更新软件

- 查看所有有新版本的软件包
- 单个更新或一键更新全部
- 自动刷新已安装列表

### 工具栏功能

- **更新 Homebrew**：更新 Homebrew 本身到最新版本
- **清理缓存**：清理旧版本文件和下载缓存


### Brew 命令封装

应用封装了常用的 Brew 命令：
- `brew list` - 列出已安装软件
- `brew search` - 搜索软件包
- `brew install` - 安装软件
- `brew uninstall` - 卸载软件
- `brew reinstall` - 重新安装
- `brew upgrade` - 更新软件
- `brew update` - 更新 Homebrew
- `brew outdated` - 查看可更新软件
- `brew cleanup` - 清理缓存
- `brew info` - 查看软件详情


## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

## 🙏 致谢

- [Wails](https://wails.io/) - 优秀的 Go + Web 桌面应用框架
- [Homebrew](https://brew.sh/) - macOS 包管理器
- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架