.PHONY: help install dev build clean run

help: ## 显示帮助信息
	@echo "可用的命令："
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## 安装依赖
	@echo "安装 Go 依赖..."
	@go mod download
	@echo "安装前端依赖..."
	@cd frontend && npm install
	@echo "✓ 依赖安装完成"

dev: ## 启动开发模式
	@echo "启动开发模式..."
	@wails dev

build: ## 构建应用
	@echo "构建应用..."
	@wails build
	@echo "✓ 构建完成，输出目录：build/bin/"

build-prod: ## 构建生产版本（压缩）
	@echo "构建生产版本..."
	@wails build -clean -upx
	@echo "✓ 构建完成，输出目录：build/bin/"

clean: ## 清理构建文件
	@echo "清理构建文件..."
	@rm -rf build/bin/
	@rm -rf frontend/dist/
	@echo "✓ 清理完成"

clean-all: clean ## 清理所有文件（包括依赖）
	@echo "清理依赖..."
	@rm -rf frontend/node_modules/
	@echo "✓ 清理完成"

run: build ## 构建并运行应用
	@echo "运行应用..."
	@open build/bin/Brew\ Manager.app

test: ## 测试 Go 代码
	@echo "运行测试..."
	@go test -v ./...

lint: ## 代码检查
	@echo "检查 Go 代码..."
	@go vet ./...
	@echo "检查前端代码..."
	@cd frontend && npm run lint || echo "请在 package.json 中添加 lint 脚本"

update-deps: ## 更新依赖
	@echo "更新 Go 依赖..."
	@go get -u ./...
	@go mod tidy
	@echo "更新前端依赖..."
	@cd frontend && npm update
	@echo "✓ 依赖更新完成"

doctor: ## 检查开发环境
	@echo "检查开发环境..."
	@echo -n "Go: "
	@go version || echo "❌ 未安装"
	@echo -n "Node: "
	@node --version || echo "❌ 未安装"
	@echo -n "NPM: "
	@npm --version || echo "❌ 未安装"
	@echo -n "Wails: "
	@wails version || echo "❌ 未安装"
	@echo -n "Homebrew: "
	@brew --version || echo "❌ 未安装"

.DEFAULT_GOAL := help

