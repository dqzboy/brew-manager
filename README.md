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


## 💻 软件截图

<br/>
<table>
    <tr>
      <td width="50%" align="center"><b>已安装软件</b></td>
      <td width="50%" align="center"><b>软件包更新</b></td>
    </tr>
    <tr>
        <td width="50%" align="center"><img src="https://cdn.jsdelivr.net/gh/dqzboy/Images@main/picture/brewmanager01.png?raw=true"></td>
        <td width="50%" align="center"><img src="https://cdn.jsdelivr.net/gh/dqzboy/Images@main/picture/brewmanager02.png?raw=true"></td>
    </tr>
    <tr>
      <td width="50%" align="center"><b>软件搜索</b></td>
      <td width="50%" align="center"><b>软件详情</b></td>
    </tr>
    <tr>
        <td width="50%" align="center"><img src="https://cdn.jsdelivr.net/gh/dqzboy/Images@main/picture/brewmanager05.png?raw=true"></td>
        <td width="50%" align="center"><img src="https://cdn.jsdelivr.net/gh/dqzboy/Images@main/picture/brewmanager06.png?raw=true"></td>
    </tr>
    <tr>
      <td width="50%" align="center"><b>软件包更新</b></td>
      <td width="50%" align="center"><b>软件包卸载</b></td>
    </tr>
    <tr>
        <td width="50%" align="center"><img src="https://cdn.jsdelivr.net/gh/dqzboy/Images@main/picture/%E6%9B%B4%E6%96%B0.png?raw=true"></td>
        <td width="50%" align="center"><img src="https://cdn.jsdelivr.net/gh/dqzboy/Images@main/picture/%E5%8D%B8%E8%BD%BD.png?raw=true"></td>
    </tr>
</table>

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

## 🙏 致谢

- [Wails](https://wails.io/) - 优秀的 Go + Web 桌面应用框架
- [Homebrew](https://brew.sh/) - macOS 包管理器
- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架