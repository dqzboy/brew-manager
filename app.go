package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Package represents a Homebrew package
type Package struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Type        string `json:"type"` // "formula" or "cask"
	Installed   bool   `json:"installed"`
}

// OperationResult represents the result of a brew operation
type OperationResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Output  string `json:"output"`
}

// GetInstalledPackages returns all installed packages
func (a *App) GetInstalledPackages() ([]Package, error) {
	packages := []Package{}

	// Get all versions at once for better performance
	allVersions := make(map[string]string)
	versionsOutput, err := a.execBrewCommand("list", "--versions")
	if err == nil {
		for _, line := range strings.Split(strings.TrimSpace(versionsOutput), "\n") {
			line = strings.TrimSpace(line)
			// 跳过空行和错误信息
			if line == "" ||
				strings.HasPrefix(line, "Error:") ||
				strings.Contains(line, "Cannot") ||
				strings.Contains(line, "✓") ||
				strings.Contains(line, "✗") {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				allVersions[parts[0]] = parts[1]
			}
		}
	}

	// Get installed formulae
	formulae, err := a.execBrewCommand("list", "--formula")
	if err == nil {
		for _, name := range strings.Split(strings.TrimSpace(formulae), "\n") {
			name = strings.TrimSpace(name)
			// 跳过空行和错误信息
			if name == "" ||
				strings.HasPrefix(name, "Error:") ||
				strings.Contains(name, "Cannot") ||
				strings.Contains(name, "✓") ||
				strings.Contains(name, "✗") {
				continue
			}
			version := allVersions[name]
			if version == "" {
				version = "unknown"
			}
			packages = append(packages, Package{
				Name:        name,
				Version:     version,
				Description: "", // 不预加载描述，提升速度
				Type:        "formula",
				Installed:   true,
			})
		}
	}

	// Get installed casks
	casks, err := a.execBrewCommand("list", "--cask")
	if err == nil {
		for _, name := range strings.Split(strings.TrimSpace(casks), "\n") {
			name = strings.TrimSpace(name)
			// 跳过空行和错误信息
			if name == "" ||
				strings.HasPrefix(name, "Error:") ||
				strings.Contains(name, "Cannot") ||
				strings.Contains(name, "✓") ||
				strings.Contains(name, "✗") {
				continue
			}
			version := allVersions[name]
			if version == "" {
				version = "unknown"
			}
			packages = append(packages, Package{
				Name:        name,
				Version:     version,
				Description: "", // 不预加载描述，提升速度
				Type:        "cask",
				Installed:   true,
			})
		}
	}

	return packages, nil
}

// SearchPackages searches for packages
func (a *App) SearchPackages(query string) ([]Package, error) {
	if query == "" {
		return []Package{}, nil
	}

	packages := []Package{}
	installedMap := make(map[string]bool)

	// 优化：并行获取已安装包列表，仅获取名称
	installedOutput, err := a.execBrewCommand("list", "--formula", "-1")
	if err == nil {
		for _, name := range strings.Split(strings.TrimSpace(installedOutput), "\n") {
			if name != "" {
				installedMap[strings.TrimSpace(name)] = true
			}
		}
	}

	// 也检查已安装的 cask
	installedCasks, err := a.execBrewCommand("list", "--cask", "-1")
	if err == nil {
		for _, name := range strings.Split(strings.TrimSpace(installedCasks), "\n") {
			if name != "" {
				installedMap[strings.TrimSpace(name)] = true
			}
		}
	}

	// Search formulae and casks
	output, err := a.execBrewCommand("search", query)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 跳过空行、标题行、错误信息和其他无效行
		if line == "" ||
			strings.HasPrefix(line, "==>") ||
			strings.HasPrefix(line, "Error:") ||
			strings.Contains(line, "Cannot") ||
			strings.Contains(line, "exit") ||
			strings.Contains(line, "[") ||
			strings.Contains(line, "✓") ||
			strings.Contains(line, "✗") {
			continue
		}

		installed := installedMap[line]

		packages = append(packages, Package{
			Name:        line,
			Description: "", // 不预加载描述以提高速度
			Type:        "formula",
			Installed:   installed,
		})
	}

	return packages, nil
}

// InstallPackage installs a package
func (a *App) InstallPackage(name string, isCask bool) OperationResult {
	args := []string{"install"}
	if isCask {
		args = append(args, "--cask")
	}
	args = append(args, name)

	output, err := a.execBrewCommand(args...)
	if err != nil {
		// 检查是否是锁文件冲突
		if strings.Contains(output, "has already locked") || strings.Contains(output, "Please wait for it to finish") {
			return OperationResult{
				Success: false,
				Message: "安装失败：另一个 Homebrew 进程正在运行，请稍后重试",
				Output:  output,
			}
		}
		return OperationResult{
			Success: false,
			Message: fmt.Sprintf("安装失败: %v", err),
			Output:  output,
		}
	}

	return OperationResult{
		Success: true,
		Message: fmt.Sprintf("成功安装 %s", name),
		Output:  output,
	}
}

// UninstallPackage uninstalls a package
func (a *App) UninstallPackage(name string, isCask bool) OperationResult {
	return a.UninstallPackageWithForce(name, isCask, false)
}

// UninstallPackageWithForce uninstalls a package with option to ignore dependencies
func (a *App) UninstallPackageWithForce(name string, isCask bool, ignoreDependencies bool) OperationResult {
	args := []string{"uninstall"}
	if isCask {
		args = append(args, "--cask")
	}
	if ignoreDependencies {
		args = append(args, "--ignore-dependencies")
	}
	args = append(args, name)

	output, err := a.execBrewCommand(args...)
	if err != nil {
		// 检查是否是依赖冲突
		if strings.Contains(output, "Refusing to uninstall") && strings.Contains(output, "because it is required by") {
			return OperationResult{
				Success: false,
				Message: "卸载失败：该软件被其他软件依赖",
				Output:  output,
			}
		}
		return OperationResult{
			Success: false,
			Message: fmt.Sprintf("卸载失败: %v", err),
			Output:  output,
		}
	}

	return OperationResult{
		Success: true,
		Message: fmt.Sprintf("成功卸载 %s", name),
		Output:  output,
	}
}

// ReinstallPackage reinstalls a package
func (a *App) ReinstallPackage(name string, isCask bool) OperationResult {
	args := []string{"reinstall"}
	if isCask {
		args = append(args, "--cask")
	}
	args = append(args, name)

	output, err := a.execBrewCommand(args...)
	if err != nil {
		// 检查是否是锁文件冲突
		if strings.Contains(output, "has already locked") || strings.Contains(output, "Please wait for it to finish") {
			return OperationResult{
				Success: false,
				Message: "重装失败：另一个 Homebrew 进程正在运行，请稍后重试",
				Output:  output,
			}
		}
		return OperationResult{
			Success: false,
			Message: fmt.Sprintf("重新安装失败: %v", err),
			Output:  output,
		}
	}

	return OperationResult{
		Success: true,
		Message: fmt.Sprintf("成功重新安装 %s", name),
		Output:  output,
	}
}

// UpdatePackage updates a package
func (a *App) UpdatePackage(name string) OperationResult {
	output, err := a.execBrewCommand("upgrade", name)
	if err != nil {
		// 检查是否是锁文件冲突
		if strings.Contains(output, "has already locked") || strings.Contains(output, "Please wait for it to finish") {
			return OperationResult{
				Success: false,
				Message: "更新失败：另一个 Homebrew 进程正在运行，请稍后重试",
				Output:  output,
			}
		}
		return OperationResult{
			Success: false,
			Message: fmt.Sprintf("更新失败: %v", err),
			Output:  output,
		}
	}

	return OperationResult{
		Success: true,
		Message: fmt.Sprintf("成功更新 %s", name),
		Output:  output,
	}
}

// UpdateAllPackages updates all packages
func (a *App) UpdateAllPackages() OperationResult {
	output, err := a.execBrewCommand("upgrade")
	if err != nil {
		// 检查是否是锁文件冲突
		if strings.Contains(output, "has already locked") || strings.Contains(output, "Please wait for it to finish") {
			return OperationResult{
				Success: false,
				Message: "更新失败：另一个 Homebrew 进程正在运行，请稍后重试",
				Output:  output,
			}
		}
		return OperationResult{
			Success: false,
			Message: fmt.Sprintf("更新全部失败: %v", err),
			Output:  output,
		}
	}

	return OperationResult{
		Success: true,
		Message: "成功更新所有软件包",
		Output:  output,
	}
}

// UpdateBrew updates Homebrew itself
func (a *App) UpdateBrew() OperationResult {
	output, err := a.execBrewCommand("update")
	if err != nil {
		return OperationResult{
			Success: false,
			Message: fmt.Sprintf("更新 Homebrew 失败: %v", err),
			Output:  output,
		}
	}

	return OperationResult{
		Success: true,
		Message: "成功更新 Homebrew",
		Output:  output,
	}
}

// GetOutdatedPackages returns packages that can be updated
func (a *App) GetOutdatedPackages() ([]Package, error) {
	output, err := a.execBrewCommand("outdated")
	if err != nil && !strings.Contains(err.Error(), "exit status") {
		return nil, err
	}

	packages := []Package{}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 跳过空行、错误信息和其他无效行
		if line == "" ||
			strings.HasPrefix(line, "Error:") ||
			strings.HasPrefix(line, "Warning:") ||
			strings.Contains(line, "Cannot") ||
			strings.Contains(line, "exit") ||
			strings.Contains(line, "[") ||
			strings.Contains(line, "✓") ||
			strings.Contains(line, "✗") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) > 0 {
			name := parts[0]
			version := ""
			if len(parts) > 1 {
				version = parts[1]
			}
			packages = append(packages, Package{
				Name:      name,
				Version:   version,
				Installed: true,
			})
		}
	}

	return packages, nil
}

// CleanupBrew cleans up old versions
func (a *App) CleanupBrew() OperationResult {
	output, err := a.execBrewCommand("cleanup")
	if err != nil {
		return OperationResult{
			Success: false,
			Message: fmt.Sprintf("清理失败: %v", err),
			Output:  output,
		}
	}

	return OperationResult{
		Success: true,
		Message: "成功清理旧版本和缓存",
		Output:  output,
	}
}

// GetPackageInfo returns detailed information about a package
func (a *App) GetPackageInfo(name string) (string, error) {
	return a.execBrewCommand("info", name)
}

// BrewConfig represents Homebrew configuration
type BrewConfig struct {
	CoreURL   string `json:"coreUrl"`
	BottleURL string `json:"bottleUrl"`
}

// GetBrewConfig returns current Homebrew configuration
func (a *App) GetBrewConfig() (BrewConfig, error) {
	config := BrewConfig{}

	// Get Homebrew repository path first
	repoPath, err := a.execBrewCommand("--repository")
	if err == nil && len(repoPath) > 0 {
		repoPath = strings.TrimSpace(repoPath)

		// Get the remote URL from git
		gitCmd := exec.Command("git", "-C", repoPath, "remote", "get-url", "origin")
		gitOutput, gitErr := gitCmd.Output()
		if gitErr == nil && len(gitOutput) > 0 {
			config.CoreURL = strings.TrimSpace(string(gitOutput))
		} else {
			config.CoreURL = repoPath
		}
	} else {
		config.CoreURL = "无法获取"
	}

	// Get Homebrew Bottle domain from environment
	bottleCmd := exec.Command("sh", "-c", "echo $HOMEBREW_BOTTLE_DOMAIN")
	bottleOutput, err := bottleCmd.Output()
	if err == nil && len(bottleOutput) > 0 && strings.TrimSpace(string(bottleOutput)) != "" {
		config.BottleURL = strings.TrimSpace(string(bottleOutput))
	} else {
		config.BottleURL = "默认源 (GitHub)"
	}

	return config, nil
}

// SetBrewMirror sets Homebrew mirror source
func (a *App) SetBrewMirror(mirrorType string) OperationResult {
	var commands [][]string

	switch mirrorType {
	case "tsinghua":
		commands = [][]string{
			{"git", "-C", "$(brew --repo)", "remote", "set-url", "origin", "https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git"},
			{"git", "-C", "$(brew --repo homebrew/core)", "remote", "set-url", "origin", "https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git"},
			{"git", "-C", "$(brew --repo homebrew/cask)", "remote", "set-url", "origin", "https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-cask.git"},
		}
	case "ustc":
		commands = [][]string{
			{"git", "-C", "$(brew --repo)", "remote", "set-url", "origin", "https://mirrors.ustc.edu.cn/brew.git"},
			{"git", "-C", "$(brew --repo homebrew/core)", "remote", "set-url", "origin", "https://mirrors.ustc.edu.cn/homebrew-core.git"},
			{"git", "-C", "$(brew --repo homebrew/cask)", "remote", "set-url", "origin", "https://mirrors.ustc.edu.cn/homebrew-cask.git"},
		}
	case "aliyun":
		commands = [][]string{
			{"git", "-C", "$(brew --repo)", "remote", "set-url", "origin", "https://mirrors.aliyun.com/homebrew/brew.git"},
			{"git", "-C", "$(brew --repo homebrew/core)", "remote", "set-url", "origin", "https://mirrors.aliyun.com/homebrew/homebrew-core.git"},
			{"git", "-C", "$(brew --repo homebrew/cask)", "remote", "set-url", "origin", "https://mirrors.aliyun.com/homebrew/homebrew-cask.git"},
		}
	case "official":
		commands = [][]string{
			{"git", "-C", "$(brew --repo)", "remote", "set-url", "origin", "https://github.com/Homebrew/brew.git"},
			{"git", "-C", "$(brew --repo homebrew/core)", "remote", "set-url", "origin", "https://github.com/Homebrew/homebrew-core.git"},
			{"git", "-C", "$(brew --repo homebrew/cask)", "remote", "set-url", "origin", "https://github.com/Homebrew/homebrew-cask.git"},
		}
	default:
		return OperationResult{
			Success: false,
			Message: "不支持的镜像源类型",
			Output:  "",
		}
	}

	// Execute commands using shell to expand $(brew --repo)
	var output strings.Builder
	for _, cmd := range commands {
		cmdStr := strings.Join(cmd, " ")
		shellCmd := exec.Command("sh", "-c", cmdStr)
		out, err := shellCmd.CombinedOutput()
		output.WriteString(string(out))
		if err != nil {
			return OperationResult{
				Success: false,
				Message: fmt.Sprintf("设置镜像源失败: %v", err),
				Output:  output.String(),
			}
		}
	}

	// Update brew
	_, _ = a.execBrewCommand("update")

	return OperationResult{
		Success: true,
		Message: "成功设置镜像源，建议重启终端以应用环境变量",
		Output:  output.String(),
	}
}

// Helper functions

func (a *App) getBrewPath() string {
	// Try to find brew in common locations
	brewPaths := []string{
		"/opt/homebrew/bin/brew",              // Apple Silicon Mac
		"/usr/local/bin/brew",                 // Intel Mac
		"/home/linuxbrew/.linuxbrew/bin/brew", // Linux
	}

	for _, path := range brewPaths {
		if _, err := exec.LookPath(path); err == nil {
			return path
		}
	}

	// Fallback to "brew" and hope it's in PATH
	return "brew"
}

func (a *App) execBrewCommand(args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	brewPath := a.getBrewPath()
	cmd := exec.CommandContext(ctx, brewPath, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func (a *App) getPackageVersion(name string) string {
	output, err := a.execBrewCommand("list", "--versions", name)
	if err != nil {
		return "unknown"
	}
	parts := strings.Fields(strings.TrimSpace(output))
	if len(parts) > 1 {
		return parts[1]
	}
	return "unknown"
}

func (a *App) getPackageDescription(name string) string {
	output, err := a.execBrewCommand("info", name)
	if err != nil {
		return ""
	}
	lines := strings.Split(output, "\n")
	if len(lines) > 0 {
		// Usually the first line contains the description
		firstLine := strings.TrimSpace(lines[0])
		if !strings.HasPrefix(firstLine, "==>") {
			return firstLine
		}
		if len(lines) > 1 {
			return strings.TrimSpace(lines[1])
		}
	}
	return ""
}

func (a *App) isPackageInstalled(name string) bool {
	_, err := a.execBrewCommand("list", name)
	return err == nil
}
