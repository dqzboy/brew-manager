<template>
  <div class="app-container">
    <!-- Header -->
    <div class="header">
      <div class="header-left">
        <div class="app-icon">🍺</div>
        <h1 class="app-title">Brew Manager</h1>
      </div>
      <div class="header-actions">
        <div class="theme-toggle-group">
          <button 
            class="btn btn-icon-only theme-toggle-btn" 
            @click="cycleTheme" 
            :title="themeTooltip"
          >
            <i v-if="currentTheme === 'light'" class="fas fa-sun"></i>
            <i v-else-if="currentTheme === 'dark'" class="fas fa-moon"></i>
            <i v-else class="fas fa-circle-half-stroke"></i>
          </button>
        </div>
        <button class="btn btn-icon" @click="updateBrew" :disabled="loading" title="更新 Homebrew">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/>
          </svg>
          更新 Homebrew
        </button>
        <button class="btn btn-icon" @click="cleanupBrew" :disabled="loading" title="清理缓存">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M3 6h18M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
          </svg>
          清理
        </button>
        <button class="btn btn-icon" @click="switchTab('settings')" :disabled="loading" title="设置">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z"/>
            <circle cx="12" cy="12" r="3"/>
          </svg>
          设置
        </button>
      </div>
    </div>

    <!-- Tabs -->
    <div class="tabs">
      <button 
        class="tab" 
        :class="{ active: currentTab === 'installed' }"
        @click="switchTab('installed')"
      >
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
          <circle cx="12" cy="7" r="4"/>
        </svg>
        已安装 <span class="tab-badge">{{ installedPackages.length }}</span>
      </button>
      <button 
        class="tab" 
        :class="{ active: currentTab === 'search' }"
        @click="switchTab('search')"
      >
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="11" cy="11" r="8"/>
          <path d="m21 21-4.35-4.35"/>
        </svg>
        搜索
      </button>
      <button 
        class="tab" 
        :class="{ active: currentTab === 'outdated' }"
        @click="switchTab('outdated')"
      >
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>
        </svg>
        可更新 <span class="tab-badge" v-if="outdatedPackages.length > 0">{{ outdatedPackages.length }}</span>
      </button>
      <button 
        class="tab" 
        :class="{ active: currentTab === 'settings' }"
        @click="switchTab('settings')"
      >
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z"/>
          <circle cx="12" cy="12" r="3"/>
        </svg>
        设置
      </button>
    </div>

    <!-- Main Content -->
    <div class="main-content">
      <!-- Loading Progress Bar -->
      <div v-if="loading" class="progress-bar-container">
        <div class="progress-bar">
          <div class="progress-bar-fill"></div>
        </div>
        <div class="progress-text">{{ loadingMessage }}</div>
      </div>

      <!-- Settings Tab -->
      <div v-if="currentTab === 'settings'" class="settings-content">
        <div class="settings-section">
          <h2 class="settings-title">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4M17 8l-5-5-5 5M12 3v12"/>
            </svg>
            应用更新
          </h2>
          <div class="settings-card">
            <div class="setting-item">
              <div class="setting-label">
                <span class="setting-name">当前版本</span>
                <span class="setting-description">Brew Manager 应用版本</span>
              </div>
              <div class="setting-value">
                <code>{{ appVersion }}</code>
              </div>
            </div>
            <div class="setting-item">
              <div class="setting-label">
                <span class="setting-name">自动检查更新</span>
                <span class="setting-description">启动时自动检查是否有新版本</span>
              </div>
              <div class="setting-toggle">
                <label class="switch">
                  <input type="checkbox" v-model="autoCheckUpdate" @change="saveUpdateSettings">
                  <span class="slider"></span>
                </label>
                <span class="toggle-label">{{ autoCheckUpdate ? '已启用' : '已禁用' }}</span>
              </div>
            </div>
            <div class="setting-item">
              <div class="setting-label">
                <span class="setting-name">更新模式</span>
                <span class="setting-description">检测到新版本时的处理方式</span>
              </div>
              <div class="setting-radio-group">
                <label class="radio-label">
                  <input type="radio" value="notify" v-model="updateMode" @change="saveUpdateSettings">
                  <span class="radio-text">
                    <span class="radio-title">仅通知</span>
                    <span class="radio-desc">提示更新但不自动下载</span>
                  </span>
                </label>
                <label class="radio-label">
                  <input type="radio" value="download" v-model="updateMode" @change="saveUpdateSettings">
                  <span class="radio-text">
                    <span class="radio-title">自动下载</span>
                    <span class="radio-desc">后台下载新版本，等待安装</span>
                  </span>
                </label>
              </div>
            </div>
            <div class="setting-actions">
              <button class="btn btn-primary" @click="checkForUpdates(false)" :disabled="loading">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/>
                </svg>
                立即检查更新
              </button>
            </div>
          </div>
        </div>

        <div class="settings-section">
          <h2 class="settings-title">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 2l-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.777-7.777zm0 0L15.5 7.5m0 0l3 3L22 7l-3-3m-3.5 3.5L19 4"/>
            </svg>
            Homebrew 镜像源
          </h2>
          <div class="settings-card">
            <div class="setting-item">
              <div class="setting-label">
                <span class="setting-name">当前 Homebrew 核心源</span>
                <span class="setting-description">brew.git 仓库地址</span>
              </div>
              <div class="setting-value">
                <code>{{ brewCoreUrl || '获取中...' }}</code>
              </div>
            </div>
            <div class="setting-item">
              <div class="setting-label">
                <span class="setting-name">当前 Homebrew Bottles 源</span>
                <span class="setting-description">预编译二进制包下载地址</span>
              </div>
              <div class="setting-value">
                <code>{{ brewBottleUrl || '获取中...' }}</code>
              </div>
            </div>
            <div class="setting-actions">
              <button class="btn btn-primary" @click="loadBrewSources" :disabled="loading">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/>
                </svg>
                刷新源信息
              </button>
            </div>
          </div>
          
          <h3 class="settings-subtitle">切换镜像源</h3>
          <div class="mirror-grid">
            <div class="mirror-card" @click="confirmSetMirror('official')" :class="{ disabled: loading }">
              <div class="mirror-icon">🌐</div>
              <div class="mirror-name">官方源</div>
              <div class="mirror-desc">GitHub 官方源</div>
              <div class="mirror-url">github.com</div>
            </div>
            <div class="mirror-card" @click="confirmSetMirror('tsinghua')" :class="{ disabled: loading }">
              <div class="mirror-icon">🎓</div>
              <div class="mirror-name">清华大学</div>
              <div class="mirror-desc">TUNA 镜像站</div>
              <div class="mirror-url">mirrors.tuna.tsinghua.edu.cn</div>
            </div>
            <div class="mirror-card" @click="confirmSetMirror('ustc')" :class="{ disabled: loading }">
              <div class="mirror-icon">🔬</div>
              <div class="mirror-name">中科大</div>
              <div class="mirror-desc">USTC 镜像站</div>
              <div class="mirror-url">mirrors.ustc.edu.cn</div>
            </div>
            <div class="mirror-card" @click="confirmSetMirror('aliyun')" :class="{ disabled: loading }">
              <div class="mirror-icon">☁️</div>
              <div class="mirror-name">阿里云</div>
              <div class="mirror-desc">阿里云镜像</div>
              <div class="mirror-url">mirrors.aliyun.com</div>
            </div>
          </div>
          
          <div class="settings-info">
            <p class="info-note">
              💡 切换镜像源后会自动执行 brew update。国内用户建议使用国内镜像源以提升下载速度。
            </p>
          </div>
        </div>

        <div class="settings-section">
          <h2 class="settings-title">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="10"/>
              <path d="M12 16v-4M12 8h.01"/>
            </svg>
            关于
          </h2>
          <div class="settings-card">
            <div class="about-info">
              <h3>Brew Manager</h3>
              <p class="about-desc">一个现代化的 Homebrew 图形界面管理工具</p>
              <div class="about-links-horizontal">
                <a @click.prevent="openURL('https://github.com/dqzboy/brew-manager')" class="about-link-compact">
                  <i class="fab fa-github"></i>
                  <span>项目</span>
                </a>
                <a @click.prevent="openURL('https://github.com/dqzboy')" class="about-link-compact">
                  <i class="fas fa-user"></i>
                  <span>作者</span>
                </a>
                <a @click.prevent="openURL('https://github.com/dqzboy/brew-manager/issues')" class="about-link-compact">
                  <i class="fas fa-bug"></i>
                  <span>反馈</span>
                </a>
              </div>
            </div>
          </div>
        </div>

        <div class="settings-section">
          <h2 class="settings-title">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M18 8h1a4 4 0 0 1 0 8h-1M2 8h16v9a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V8zM6 1v3M10 1v3M14 1v3"/>
            </svg>
            推广合作
          </h2>
          <div class="promo-carousel-container">
            <button class="carousel-control carousel-prev" @click="scrollPromoLeft" :disabled="promoScrollPosition <= 0">
              <i class="fas fa-chevron-left"></i>
            </button>
            
            <div class="promo-carousel" ref="promoCarousel">
              <div class="promo-track-wrapper">
                <div class="ad-card ad-mojie" @click="openURL('https://mojie.co/register?aff=CG6h8Irm')">
                  <div class="ad-header">
                    <div class="ad-icon">✈️</div>
                    <div class="ad-category">机场推广</div>
                  </div>
                  <h4 class="ad-title">魔戒</h4>
                  <p class="ad-desc">流量不过期 不限速 不限设备数、低延迟、高网速</p>
                  <button class="ad-button" @click.stop="openURL('https://mojie.co/register?aff=CG6h8Irm')">
                    访问官网 <i class="fas fa-arrow-right"></i>
                  </button>
                </div>

                <div class="ad-card ad-cloudcone" @click="openURL('https://dqzboy.github.io/proxyui/CloudCone')">
                  <div class="ad-header">
                    <div class="ad-icon">☁️</div>
                    <div class="ad-category">云服务商</div>
                  </div>
                  <h4 class="ad-title">CloudCone</h4>
                  <p class="ad-desc">提供稳定、超高性价比云计算基础设施支持</p>
                  <button class="ad-button" @click.stop="openURL('https://dqzboy.github.io/proxyui/CloudCone')">
                    访问官网 <i class="fas fa-arrow-right"></i>
                  </button>
                </div>

                <div class="ad-card ad-racknerd" @click="openURL('https://dqzboy.github.io/proxyui/racknerd')">
                  <div class="ad-header">
                    <div class="ad-icon">☁️</div>
                    <div class="ad-category">云服务商</div>
                  </div>
                  <h4 class="ad-title">RackNerd</h4>
                  <p class="ad-desc">高性能美国云服务器，多机房可选，续费同价</p>
                  <button class="ad-button" @click.stop="openURL('https://dqzboy.github.io/proxyui/racknerd')">
                    访问官网 <i class="fas fa-arrow-right"></i>
                  </button>
                </div>

                <div class="ad-card ad-teacat" @click="openURL('https://teacat999.com/#/register?code=ps4sZcDa')">
                  <div class="ad-header">
                    <div class="ad-icon">✈️</div>
                    <div class="ad-category">机场推广</div>
                  </div>
                  <h4 class="ad-title">Teacat</h4>
                  <p class="ad-desc">不限速、不限设备、IEPL专线，解锁ChatGPT、速度快、稳定</p>
                  <button class="ad-button" @click.stop="openURL('https://teacat999.com/#/register?code=ps4sZcDa')">
                    访问官网 <i class="fas fa-arrow-right"></i>
                  </button>
                </div>
              </div>
            </div>

            <button class="carousel-control carousel-next" @click="scrollPromoRight" :disabled="promoAtEnd">
              <i class="fas fa-chevron-right"></i>
            </button>
          </div>
        </div>
      </div>

      <!-- Other Tabs Content -->
      <div v-else class="tab-content-wrapper">
        <!-- Search Bar - Fixed -->
        <div class="search-bar-fixed">
        <input 
          v-if="currentTab === 'search'"
          type="text" 
          class="search-input" 
          v-model="searchQuery"
          @keyup.enter="searchPackages"
          placeholder="输入软件名称搜索..."
        >
        <input 
          v-else
          type="text" 
          class="search-input" 
          v-model="filterQuery"
          placeholder="过滤列表..."
        >
        <button 
          v-if="currentTab === 'search'" 
          class="btn btn-primary" 
          @click="searchPackages"
          :disabled="loading || !searchQuery"
        >
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="8"/>
              <path d="m21 21-4.35-4.35"/>
            </svg>
          搜索
        </button>
        <button 
          v-if="currentTab === 'outdated'" 
          class="btn btn-success" 
          @click="updateAll"
          :disabled="loading || outdatedPackages.length === 0"
        >
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/>
            </svg>
          全部更新
        </button>
      </div>

      <!-- Package List -->
      <div class="package-list">
          <div v-if="currentPackages.length === 0 && !loading" class="empty-state">
            <div class="empty-icon">
              <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <circle cx="12" cy="12" r="10"/>
                <path d="M12 8v4m0 4h.01"/>
              </svg>
        </div>
          <p v-if="currentTab === 'search'">输入关键词搜索软件包</p>
          <p v-else-if="currentTab === 'outdated'">🎉 所有软件都是最新版本</p>
          <p v-else>暂无已安装的软件包</p>
        </div>

        <div v-else>
          <div 
            v-for="pkg in currentPackages" 
            :key="pkg.name" 
            class="package-item"
          >
            <div class="package-info">
              <div class="package-name">{{ pkg.name }}</div>
              <div class="package-meta">
                  <span class="package-version" v-if="pkg.version && pkg.version !== 'unknown'">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M20 6L9 17l-5-5"/>
                    </svg>
                    {{ pkg.version }}
                </span>
                  <span class="package-type" :class="`type-${pkg.type}`">
                    {{ pkg.type === 'formula' ? '📦 Formula' : '🎯 Cask' }}
                  </span>
                <span 
                  class="status-badge status-installed" 
                  v-if="pkg.installed"
                >
                  已安装
                </span>
              </div>
                <div class="package-description" v-if="pkg.description && pkg.description.trim()">
                {{ pkg.description }}
              </div>
            </div>
            <div class="package-actions">
              <button 
                v-if="!pkg.installed" 
                class="btn btn-success btn-small"
                @click="installPackage(pkg)"
                :disabled="loading"
                  title="安装此软件包"
              >
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4M7 10l5 5 5-5M12 15V3"/>
                  </svg>
                安装
              </button>
              <button 
                v-if="pkg.installed && currentTab !== 'outdated'" 
                class="btn btn-secondary btn-small"
                @click="showPackageInfo(pkg)"
                :disabled="loading"
                  title="查看详细信息"
              >
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="12" cy="12" r="10"/>
                    <path d="M12 16v-4M12 8h.01"/>
                  </svg>
                详情
              </button>
              <button 
                v-if="pkg.installed && currentTab === 'outdated'" 
                class="btn btn-success btn-small"
                @click="updatePackage(pkg)"
                :disabled="loading"
                  title="更新到最新版本"
              >
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/>
                  </svg>
                更新
              </button>
              <button 
                v-if="pkg.installed" 
                  class="btn btn-warning btn-small"
                  @click="confirmReinstall(pkg)"
                :disabled="loading"
                  title="重新安装"
              >
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M21.5 2v6h-6M2.5 22v-6h6M2 11.5a10 10 0 0 1 18.8-4.3M22 12.5a10 10 0 0 1-18.8 4.2"/>
                  </svg>
                重装
              </button>
              <button 
                v-if="pkg.installed" 
                class="btn btn-danger btn-small"
                  @click="confirmUninstall(pkg)"
                :disabled="loading"
                  title="卸载此软件包"
              >
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M3 6h18M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
                  </svg>
                卸载
              </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Operation Progress Modal -->
    <div v-for="operation in operationModals" :key="operation.id">
      <transition name="modal">
        <div v-if="operation.show && !operation.isMinimized" class="modal-overlay operation-overlay" @click.self="closeOperationModal(operation.id)">
          <div class="modal operation-modal" @click.stop>
            <div class="operation-header">
              <div class="operation-icon">
                <div v-if="!operation.isCompleted" class="spinner-modern"></div>
                <svg v-else-if="operation.title === '操作完成'" class="icon-success" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20 6L9 17l-5-5"/>
                </svg>
                <svg v-else class="icon-error" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"/>
                  <path d="m15 9-6 6m0-6 6 6"/>
                </svg>
              </div>
              <div class="operation-info">
                <h3 class="operation-title">{{ operation.title }}</h3>
                <p class="operation-subtitle">{{ operation.subtitle }}</p>
              </div>
              <button class="operation-minimize" @click="toggleOperationMinimize(operation.id)" title="最小化">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M19 12H5"/>
                </svg>
              </button>
              <button class="operation-close" @click="closeOperationModal(operation.id)" title="关闭">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="m18 6-12 12M6 6l12 12"/>
                </svg>
              </button>
            </div>
            
            <div v-if="!operation.isCompleted" class="operation-progress">
              <div class="progress-bar-modern">
                <div class="progress-bar-fill-modern"></div>
              </div>
            </div>

            <div class="operation-logs">
              <div class="logs-header">
                <span class="logs-title">
                  <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                    <path d="M14 2v6h6M16 13H8m8 4H8m8 4H8"/>
                  </svg>
                  操作日志
                </span>
                <span class="logs-count">{{ operation.logs.length }} 条</span>
              </div>
              <div class="logs-content" ref="logsContainer">
                <div v-if="operation.logs.length === 0" class="logs-empty">
                  等待操作输出...
                </div>
                <div v-else>
                  <div 
                    v-for="(log, index) in operation.logs" 
                    :key="index" 
                    class="log-line"
                    :class="log.type"
                  >
                    <span class="log-time">{{ log.time }}</span>
                    <span class="log-text">{{ log.text }}</span>
                  </div>
                </div>
              </div>
            </div>

            <div class="operation-footer" v-if="operation.canClose">
              <button class="btn btn-primary" @click="closeOperationModal(operation.id)">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M20 6L9 17l-5-5"/>
                </svg>
                完成
              </button>
            </div>
          </div>
        </div>
      </transition>
    </div>

    <!-- Operation Toast (Minimized) - Multiple Support with stacking -->
    <div class="operation-toasts-container">
      <transition-group name="toast-list">
        <div 
          v-for="(operation, index) in operationModals.filter(op => op.show && op.isMinimized)" 
          :key="operation.id" 
          class="operation-toast"
          :style="{ bottom: (24 + index * 90) + 'px' }"
          @click="toggleOperationMinimize(operation.id)"
        >
          <div class="operation-header">
            <div class="operation-icon">
              <div v-if="!operation.isCompleted" class="spinner-modern"></div>
              <svg v-else-if="operation.title === '操作完成'" class="icon-success" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                <path d="M20 6L9 17l-5-5"/>
              </svg>
              <svg v-else class="icon-error" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                <circle cx="12" cy="12" r="10"/>
                <path d="m15 9-6 6m0-6 6 6"/>
              </svg>
            </div>
            <div class="operation-info">
              <h3 class="operation-title">{{ operation.title }}</h3>
              <p class="operation-subtitle">{{ operation.subtitle }}</p>
            </div>
            <button class="operation-close" @click.stop="closeOperationModal(operation.id)" title="关闭">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="m18 6-12 12M6 6l12 12"/>
              </svg>
            </button>
          </div>
            
          <div v-if="!operation.isCompleted" class="operation-progress">
            <div class="progress-bar-modern">
              <div class="progress-bar-fill-modern"></div>
            </div>
          </div>
          <div v-if="operation.logs.length > 0" class="operation-latest-log">
            {{ operation.logs[operation.logs.length - 1].text }}
          </div>
        </div>
      </transition-group>
    </div>

    <!-- Footer -->
    <div class="footer">
      <div class="footer-content">
        <div class="footer-info">
          <a @click.prevent="openURL('https://github.com/dqzboy/brew-manager')" class="footer-title">Brew Manager</a>
          <span class="footer-separator">•</span>
          <span class="footer-version">{{ appVersion }}</span>
        </div>
        <div class="footer-author">
          <span>Made with</span>
          <svg width="14" height="14" viewBox="0 0 24 24" fill="currentColor" class="heart-icon">
            <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
          </svg>
          <span>by <a @click.prevent="openURL('https://github.com/dqzboy')" class="author-link">dqzboy</a></span>
        </div>
      </div>
    </div>

    <!-- Back to Top Button (只在已安装/搜索/可更新页面显示) -->
    <transition name="fade">
      <button 
        v-if="showBackToTop && currentTab !== 'settings'" 
        class="back-to-top-btn" 
        @click="scrollToTop"
        title="返回顶部"
      >
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M12 19V5M5 12l7-7 7 7"/>
        </svg>
      </button>
    </transition>

    <!-- Toast Notification -->
    <transition name="toast">
      <div v-if="toast.show" class="toast" :class="`toast-${toast.type}`">
        <div class="toast-icon">
          <svg v-if="toast.type === 'success'" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M20 6L9 17l-5-5"/>
          </svg>
          <svg v-else-if="toast.type === 'error'" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="10"/>
            <path d="m15 9-6 6m0-6 6 6"/>
          </svg>
          <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="10"/>
            <path d="M12 16v-4M12 8h.01"/>
          </svg>
    </div>
        <div class="toast-message">{{ toast.message }}</div>
      </div>
    </transition>

    <!-- Confirmation Modal -->
    <transition name="modal">
      <div v-if="confirmDialog.show" class="modal-overlay" @click="closeConfirmDialog">
        <div class="modal confirm-modal" @click.stop>
          <div class="modal-header">
            <div class="modal-icon" :class="`icon-${confirmDialog.type}`">
              <svg v-if="confirmDialog.type === 'danger'" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
                <path d="M12 9v4m0 4h.01"/>
              </svg>
              <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"/>
                <path d="M12 16v-4M12 8h.01"/>
              </svg>
            </div>
            <h3 class="modal-title">{{ confirmDialog.title }}</h3>
          </div>
          <div class="modal-content">
            <p>{{ confirmDialog.message }}</p>
          </div>
          <div class="modal-actions">
            <button class="btn btn-secondary" @click="closeConfirmDialog">取消</button>
            <button 
              class="btn" 
              :class="confirmDialog.type === 'danger' ? 'btn-danger' : 'btn-primary'"
              @click="confirmAction"
            >
              确认
            </button>
          </div>
        </div>
      </div>
    </transition>

    <!-- Info Modal -->
    <transition name="modal">
    <div v-if="modal.show" class="modal-overlay" @click="closeModal">
        <div class="modal info-modal" @click.stop>
          <div class="modal-header">
            <h3 class="modal-title">{{ modal.title }}</h3>
            <button class="modal-close" @click="closeModal">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="m18 6-12 12M6 6l12 12"/>
              </svg>
            </button>
          </div>
          <div class="modal-content">
            <pre>{{ modal.content }}</pre>
          </div>
        <div class="modal-actions">
            <button class="btn btn-primary" @click="closeModal">关闭</button>
        </div>
      </div>
    </div>
    </transition>

    <!-- Update Available Modal -->
    <transition name="modal">
      <div v-if="updateDialog.show" class="modal-overlay" @click="closeUpdateDialog">
        <div class="modal update-modal" @click.stop>
          <div class="update-header">
            <div class="update-icon-wrapper">
              <i class="fas fa-download update-icon"></i>
            </div>
            <h3 class="update-title">发现新版本</h3>
            <p class="update-version">{{ updateDialog.version }}</p>
          </div>
          <div class="update-content">
            <h4 class="update-section-title">更新内容</h4>
            <div class="update-description">{{ updateDialog.description }}</div>
            <div class="update-info">
              <div class="info-row">
                <span class="info-label">发布时间：</span>
                <span class="info-value">{{ updateDialog.publishedAt }}</span>
              </div>
              <div class="info-row">
                <span class="info-label">文件大小：</span>
                <span class="info-value">{{ updateDialog.fileSize }}</span>
              </div>
            </div>
          </div>
          <div class="modal-actions">
            <button class="btn btn-secondary" @click="closeUpdateDialog">稍后提醒</button>
            <button class="btn btn-success" @click="downloadUpdate">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4M7 10l5 5 5-5M12 15V3"/>
              </svg>
              立即更新
            </button>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { GetInstalledPackages, SearchPackages, InstallPackage, UninstallPackage, 
         UninstallPackageWithForce, ReinstallPackage, UpdatePackage, UpdateAllPackages, UpdateBrew, 
         CleanupBrew, GetOutdatedPackages, GetPackageInfo, GetBrewConfig, SetBrewMirror } from '../wailsjs/go/main/App'
import { BrowserOpenURL } from '../wailsjs/runtime/runtime'

export default {
  name: 'App',
  setup() {
    const currentTab = ref('installed')
    const loading = ref(false)
    const loadingMessage = ref('加载中...')
    const installedPackages = ref([])
    const searchResults = ref([])
    const outdatedPackages = ref([])
    const searchQuery = ref('')
    const filterQuery = ref('')
    const brewCoreUrl = ref('')
    const brewBottleUrl = ref('')
    const currentTheme = ref('dark') // 'light', 'dark', 'auto'
    const showBackToTop = ref(false)
    const autoCheckUpdate = ref(true)
    const updateMode = ref('notify') // 'notify' or 'download'
    const appVersion = ref('v1.0.0') // 从GitHub获取的实际版本号
    const promoCarousel = ref(null)
    const promoScrollPosition = ref(0)
    const promoAtEnd = ref(false)
    
    const toast = ref({
      show: false,
      message: '',
      type: 'success'
    })

    const updateDialog = ref({
      show: false,
      version: '',
      description: '',
      publishedAt: '',
      fileSize: '',
      downloadUrl: ''
    })

    const modal = ref({
      show: false,
      title: '',
      content: ''
    })

    const confirmDialog = ref({
      show: false,
      title: '',
      message: '',
      type: 'warning',
      onConfirm: null
    })

    const operationModals = ref([])
    const nextOperationId = ref(0)

    const logsContainer = ref(null)

    const themeTooltip = computed(() => {
      if (currentTheme.value === 'light') return '当前：浅色主题 (点击切换)'
      if (currentTheme.value === 'dark') return '当前：深色主题 (点击切换)'
      return '当前：跟随系统 (点击切换)'
    })

    const currentPackages = computed(() => {
      let packages = []
      
      if (currentTab.value === 'installed') {
        packages = installedPackages.value
      } else if (currentTab.value === 'search') {
        packages = searchResults.value
      } else if (currentTab.value === 'outdated') {
        packages = outdatedPackages.value
      }

      if (filterQuery.value && currentTab.value !== 'search') {
        const query = filterQuery.value.toLowerCase()
        packages = packages.filter(pkg => 
          pkg.name.toLowerCase().includes(query) ||
          (pkg.description && pkg.description.toLowerCase().includes(query))
        )
      }

      return packages
    })

    const showToast = (message, type = 'success') => {
      toast.value = { show: true, message, type }
      setTimeout(() => {
        toast.value.show = false
      }, 3000)
    }

    const showModal = (title, content) => {
      console.log('showModal called:', title)
      modal.value = { 
        show: true, 
        title, 
        content 
      }
      console.log('modal.value:', modal.value)
    }

    const closeModal = () => {
      modal.value.show = false
    }

    const showConfirmDialog = (title, message, type, onConfirm) => {
      confirmDialog.value = {
        show: true,
        title,
        message,
        type,
        onConfirm
      }
    }

    const closeConfirmDialog = () => {
      confirmDialog.value.show = false
      confirmDialog.value.onConfirm = null
    }

    const confirmAction = () => {
      if (confirmDialog.value.onConfirm) {
        confirmDialog.value.onConfirm()
      }
      closeConfirmDialog()
    }

    const showOperationModal = (title, subtitle) => {
      const id = nextOperationId.value++
      const newOperation = {
        id,
        show: true,
        title,
        subtitle,
        logs: [],
        canClose: false,
        isCompleted: false,
        isMinimized: false
      }
      operationModals.value.push(newOperation)
      return id
    }

    const addOperationLog = (text, type = 'info', operationId = null) => {
      const now = new Date()
      const time = `${String(now.getHours()).padStart(2, '0')}:${String(now.getMinutes()).padStart(2, '0')}:${String(now.getSeconds()).padStart(2, '0')}`
      
      const operation = operationId !== null 
        ? operationModals.value.find(op => op.id === operationId)
        : operationModals.value[operationModals.value.length - 1]
      
      if (operation) {
        operation.logs.push({
          time,
          text,
          type
        })

        // 自动滚动到底部
        setTimeout(() => {
          if (logsContainer.value) {
            logsContainer.value.scrollTop = logsContainer.value.scrollHeight
          }
        }, 10)
      }
    }

    const closeOperationModal = (operationId) => {
      const index = operationModals.value.findIndex(op => op.id === operationId)
      if (index !== -1) {
        operationModals.value.splice(index, 1)
      }
    }

    const toggleOperationMinimize = (operationId) => {
      const operation = operationModals.value.find(op => op.id === operationId)
      if (operation) {
        operation.isMinimized = !operation.isMinimized
      }
    }

    const completeOperation = (success, message, operationId = null) => {
      const operation = operationId !== null 
        ? operationModals.value.find(op => op.id === operationId)
        : operationModals.value[operationModals.value.length - 1]
      
      if (operation) {
        operation.canClose = true
        operation.isCompleted = true
        if (success) {
          operation.title = '操作完成'
          operation.subtitle = message
          addOperationLog('✓ ' + message, 'success', operation.id)
        } else {
          operation.title = '操作失败'
          operation.subtitle = message
          addOperationLog('✗ ' + message, 'error', operation.id)
        }
        
        // 如果是最小化模式，操作完成后3秒自动关闭
        if (operation.isMinimized) {
          setTimeout(() => {
            if (operation.isCompleted && operation.isMinimized) {
              closeOperationModal(operation.id)
            }
          }, 3000)
        }
      }
    }

    const switchTab = async (tab) => {
      currentTab.value = tab
      filterQuery.value = ''
      showBackToTop.value = false // 切换标签时隐藏返回顶部按钮
      
      if (tab === 'installed' && installedPackages.value.length === 0) {
        await loadInstalledPackages()
      } else if (tab === 'outdated') {
        await loadOutdatedPackages()
      } else if (tab === 'settings') {
        await loadBrewSources()
      }

      // 切换标签后重新设置滚动监听
      setTimeout(() => {
        setupScrollListeners()
      }, 100)
    }

    const loadInstalledPackages = async () => {
      loading.value = true
      loadingMessage.value = '正在加载已安装的软件包...'
      try {
        installedPackages.value = await GetInstalledPackages()
      } catch (error) {
        showToast('加载已安装软件失败: ' + error, 'error')
      } finally {
        loading.value = false
        // 加载完成后重新设置滚动监听
        setTimeout(() => {
          setupScrollListeners()
        }, 100)
      }
    }

    const loadOutdatedPackages = async () => {
      try {
        outdatedPackages.value = await GetOutdatedPackages()
      } catch (error) {
        showToast('加载可更新软件失败: ' + error, 'error')
      } finally {
        // 加载完成后重新设置滚动监听
        setTimeout(() => {
          setupScrollListeners()
        }, 100)
      }
    }

    const searchPackages = async () => {
      if (!searchQuery.value) return
      
      loading.value = true
      loadingMessage.value = '正在搜索软件包...'
      try {
        // 优化搜索：使用 brew search --formula 和 --cask 分别搜索会更快
        searchResults.value = await SearchPackages(searchQuery.value)
        loading.value = false
        
        if (searchResults.value.length === 0) {
          showToast('未找到相关软件包', 'info')
        } else {
          showToast(`找到 ${searchResults.value.length} 个软件包`, 'success')
        }

        // 搜索完成后重新设置滚动监听
        setTimeout(() => {
          setupScrollListeners()
        }, 100)
      } catch (error) {
        showToast('搜索失败: ' + error, 'error')
        loading.value = false
      }
    }

    const loadBrewSources = async () => {
      loading.value = true
      loadingMessage.value = '正在获取 Homebrew 源信息...'
      try {
        const config = await GetBrewConfig()
        brewCoreUrl.value = config.coreUrl || '未知'
        brewBottleUrl.value = config.bottleUrl || '默认源 (GitHub)'
      } catch (error) {
        showToast('获取源信息失败: ' + error, 'error')
        brewCoreUrl.value = '获取失败'
        brewBottleUrl.value = '获取失败'
      } finally {
        loading.value = false
      }
    }

    const confirmSetMirror = (mirrorType) => {
      if (loading.value) return
      
      const mirrorNames = {
        'official': 'GitHub 官方源',
        'tsinghua': '清华大学镜像',
        'ustc': '中科大镜像',
        'aliyun': '阿里云镜像'
      }
      
      showConfirmDialog(
        '切换镜像源',
        `确定要切换到 ${mirrorNames[mirrorType]} 吗？此操作会修改 Homebrew 的远程仓库地址并执行更新。`,
        'warning',
        () => setBrewMirror(mirrorType)
      )
    }

    const setBrewMirror = async (mirrorType) => {
      loading.value = true
      loadingMessage.value = '正在切换镜像源，这可能需要几分钟...'
      try {
        const result = await SetBrewMirror(mirrorType)
        if (result.success) {
          showToast(result.message, 'success')
          await loadBrewSources()
        } else {
          showToast(result.message, 'error')
        }
      } catch (error) {
        showToast('切换镜像源失败: ' + error, 'error')
      } finally {
        loading.value = false
      }
    }

    const installPackage = async (pkg) => {
      showOperationModal('安装软件包', `正在安装 ${pkg.name}`)
      addOperationLog(`开始安装 ${pkg.name} (${pkg.type})`)
      
      try {
        addOperationLog('正在下载软件包...')
        await new Promise(resolve => setTimeout(resolve, 500))
        
        addOperationLog('正在解析依赖关系...')
        await new Promise(resolve => setTimeout(resolve, 300))
        
        addOperationLog('正在执行安装命令...')
        const result = await InstallPackage(pkg.name, pkg.type === 'cask')
        
        if (result.output) {
          const lines = result.output.split('\n').filter(line => line.trim())
          lines.forEach(line => addOperationLog(line))
        }
        
        if (result.success) {
          completeOperation(true, result.message)
          showToast(result.message, 'success')
          await loadInstalledPackages()
          // 更新搜索结果中的状态
          if (searchResults.value.length > 0) {
            const index = searchResults.value.findIndex(p => p.name === pkg.name)
            if (index !== -1) {
              searchResults.value[index].installed = true
            }
          }
        } else {
          if (result.message.includes('另一个 Homebrew 进程正在运行')) {
            completeOperation(false, result.message + '\n\n💡 提示：请等待其他 Homebrew 操作完成，或在终端执行 "killall brew" 结束所有 brew 进程后重试。')
          } else {
            completeOperation(false, result.message)
          }
          showToast(result.message, 'error')
        }
      } catch (error) {
        completeOperation(false, '安装失败: ' + error)
        showToast('安装失败: ' + error, 'error')
      }
    }

    const confirmUninstall = (pkg) => {
      showConfirmDialog(
        '确认卸载',
        `确定要卸载 ${pkg.name} 吗？此操作无法撤销。`,
        'danger',
        () => uninstallPackage(pkg)
      )
    }

    const uninstallPackage = async (pkg, force = false) => {
      showOperationModal('卸载软件包', `正在卸载 ${pkg.name}`)
      addOperationLog(`开始卸载 ${pkg.name} (${pkg.type})`)
      
      try {
        if (!force) {
          addOperationLog('正在检查依赖关系...')
          await new Promise(resolve => setTimeout(resolve, 300))
        } else {
          addOperationLog('正在强制卸载（忽略依赖）...')
        }
        
        addOperationLog('正在执行卸载命令...')
        const result = force 
          ? await UninstallPackageWithForce(pkg.name, pkg.type === 'cask', true)
          : await UninstallPackage(pkg.name, pkg.type === 'cask')
        
        if (result.output) {
          const lines = result.output.split('\n').filter(line => line.trim())
          lines.forEach(line => addOperationLog(line))
        }
        
        if (result.success) {
          completeOperation(true, result.message)
          showToast(result.message, 'success')
          await loadInstalledPackages()
          // 更新搜索结果中的状态
          if (searchResults.value.length > 0) {
            const index = searchResults.value.findIndex(p => p.name === pkg.name)
            if (index !== -1) {
              searchResults.value[index].installed = false
            }
          }
        } else {
          // 检查是否是依赖冲突
          if (result.message.includes('被其他软件依赖') && !force) {
            closeOperationModal()
            // 从输出中提取依赖信息
            const dependencyMatch = result.output.match(/because it is required by ([^,]+)/)
            const dependency = dependencyMatch ? dependencyMatch[1].trim() : '其他软件'
            
            showConfirmDialog(
              '检测到依赖关系',
              `${pkg.name} 被 ${dependency} 依赖。\n\n强制卸载可能会导致 ${dependency} 无法正常工作。\n\n是否强制卸载？`,
              'danger',
              () => uninstallPackage(pkg, true)
            )
          } else {
            completeOperation(false, result.message)
            showToast(result.message, 'error')
          }
        }
      } catch (error) {
        completeOperation(false, '卸载失败: ' + error)
        showToast('卸载失败: ' + error, 'error')
      }
    }

    const confirmReinstall = (pkg) => {
      showConfirmDialog(
        '确认重装',
        `确定要重新安装 ${pkg.name} 吗？`,
        'warning',
        () => reinstallPackage(pkg)
      )
    }

    const reinstallPackage = async (pkg) => {
      showOperationModal('重新安装', `正在重新安装 ${pkg.name}`)
      addOperationLog(`开始重新安装 ${pkg.name} (${pkg.type})`)
      
      try {
        addOperationLog('正在卸载旧版本...')
        await new Promise(resolve => setTimeout(resolve, 500))
        
        addOperationLog('正在下载最新版本...')
        await new Promise(resolve => setTimeout(resolve, 500))
        
        addOperationLog('正在执行安装命令...')
        const result = await ReinstallPackage(pkg.name, pkg.type === 'cask')
        
        if (result.output) {
          const lines = result.output.split('\n').filter(line => line.trim())
          lines.forEach(line => addOperationLog(line))
        }
        
        if (result.success) {
          completeOperation(true, result.message)
          showToast(result.message, 'success')
          await loadInstalledPackages()
        } else {
          if (result.message.includes('另一个 Homebrew 进程正在运行')) {
            completeOperation(false, result.message + '\n\n💡 提示：请等待其他 Homebrew 操作完成，或在终端执行 "killall brew" 结束所有 brew 进程后重试。')
          } else {
            completeOperation(false, result.message)
          }
          showToast(result.message, 'error')
        }
      } catch (error) {
        completeOperation(false, '重新安装失败: ' + error)
        showToast('重新安装失败: ' + error, 'error')
      }
    }

    const updatePackage = async (pkg) => {
      showOperationModal('更新软件包', `正在更新 ${pkg.name}`)
      addOperationLog(`开始更新 ${pkg.name}`)
      
      try {
        addOperationLog('正在检查可用更新...')
        await new Promise(resolve => setTimeout(resolve, 300))
        
        addOperationLog('正在下载新版本...')
        await new Promise(resolve => setTimeout(resolve, 500))
        
        addOperationLog('正在执行更新命令...')
        const result = await UpdatePackage(pkg.name)
        
        if (result.output) {
          const lines = result.output.split('\n').filter(line => line.trim())
          lines.forEach(line => addOperationLog(line))
        }
        
        if (result.success) {
          completeOperation(true, result.message)
          showToast(result.message, 'success')
          await loadInstalledPackages()
          await loadOutdatedPackages()
        } else {
          // 检查是否是锁文件冲突
          if (result.message.includes('另一个 Homebrew 进程正在运行')) {
            completeOperation(false, result.message + '\n\n💡 提示：请等待其他 Homebrew 操作完成，或在终端执行 "killall brew" 结束所有 brew 进程后重试。')
          } else {
            completeOperation(false, result.message)
          }
          showToast(result.message, 'error')
        }
      } catch (error) {
        completeOperation(false, '更新失败: ' + error)
        showToast('更新失败: ' + error, 'error')
      }
    }

    const updateAll = async () => {
      loading.value = true
      loadingMessage.value = '正在更新所有软件包...'
      try {
        const result = await UpdateAllPackages()
        if (result.success) {
          showToast(result.message, 'success')
          await loadInstalledPackages()
          await loadOutdatedPackages()
        } else {
          showToast(result.message, 'error')
        }
      } catch (error) {
        showToast('更新全部失败: ' + error, 'error')
      } finally {
        loading.value = false
      }
    }

    const updateBrew = async () => {
      loading.value = true
      loadingMessage.value = '正在更新 Homebrew...'
      try {
        const result = await UpdateBrew()
        if (result.success) {
          showToast(result.message, 'success')
        } else {
          showToast(result.message, 'error')
        }
      } catch (error) {
        showToast('更新 Homebrew 失败: ' + error, 'error')
      } finally {
        loading.value = false
      }
    }

    const cleanupBrew = async () => {
      loading.value = true
      loadingMessage.value = '正在清理缓存...'
      try {
        const result = await CleanupBrew()
        if (result.success) {
          showToast(result.message, 'success')
        } else {
          showToast(result.message, 'error')
        }
      } catch (error) {
        showToast('清理失败: ' + error, 'error')
      } finally {
        loading.value = false
      }
    }

    const showPackageInfo = async (pkg) => {
      loading.value = true
      loadingMessage.value = `正在获取 ${pkg.name} 的详细信息...`
      try {
        const info = await GetPackageInfo(pkg.name)
        showModal(`${pkg.name} 详细信息`, info)
      } catch (error) {
        showToast('获取详情失败: ' + error, 'error')
      } finally {
        loading.value = false
      }
    }

    // 主题相关函数
    const applyTheme = (theme) => {
      const root = document.documentElement
      
      if (theme === 'auto') {
        // 检测系统主题
        const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
        root.setAttribute('data-theme', prefersDark ? 'dark' : 'light')
      } else {
        root.setAttribute('data-theme', theme)
      }
    }

    const setTheme = (theme) => {
      currentTheme.value = theme
      applyTheme(theme)
      // 保存到本地存储
      localStorage.setItem('brew-manager-theme', theme)
      showToast(`已切换到${theme === 'light' ? '浅色' : theme === 'dark' ? '深色' : '跟随系统'}主题`, 'success')
    }

    const cycleTheme = () => {
      const themes = ['light', 'dark', 'auto']
      const currentIndex = themes.indexOf(currentTheme.value)
      const nextIndex = (currentIndex + 1) % themes.length
      setTheme(themes[nextIndex])
    }

    const initTheme = () => {
      // 从本地存储读取主题设置
      const savedTheme = localStorage.getItem('brew-manager-theme') || 'dark'
      currentTheme.value = savedTheme
      applyTheme(savedTheme)

      // 如果是自动模式，监听系统主题变化
      if (savedTheme === 'auto') {
        const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
        mediaQuery.addEventListener('change', (e) => {
          if (currentTheme.value === 'auto') {
            applyTheme('auto')
          }
        })
      }
    }

    // 更新检测相关
    const loadUpdateSettings = () => {
      const savedAutoCheck = localStorage.getItem('brew-manager-auto-check-update')
      const savedUpdateMode = localStorage.getItem('brew-manager-update-mode')
      
      autoCheckUpdate.value = savedAutoCheck !== 'false'
      updateMode.value = savedUpdateMode || 'notify'
    }

    const saveUpdateSettings = () => {
      localStorage.setItem('brew-manager-auto-check-update', autoCheckUpdate.value.toString())
      localStorage.setItem('brew-manager-update-mode', updateMode.value)
      showToast('设置已保存', 'success')
    }

    const fetchCurrentVersion = async () => {
      try {
        // 尝试从GitHub获取当前发布的版本号
        const response = await fetch('https://api.github.com/repos/dqzboy/brew-manager/releases/latest', {
          headers: {
            'Accept': 'application/vnd.github.v3+json'
          }
        })
        
        if (response.ok) {
          const data = await response.json()
          if (data.tag_name) {
            appVersion.value = data.tag_name
            // 保存到本地，下次离线时使用
            localStorage.setItem('brew-manager-version', data.tag_name)
          }
        } else {
          // 如果没有发布版本，使用本地存储或默认值
          const savedVersion = localStorage.getItem('brew-manager-version')
          if (savedVersion) {
            appVersion.value = savedVersion
          }
        }
      } catch (error) {
        // 网络错误时使用本地存储的版本
        const savedVersion = localStorage.getItem('brew-manager-version')
        if (savedVersion) {
          appVersion.value = savedVersion
        }
      }
    }

    const checkForUpdates = async (silent = false) => {
      let shouldShowModal = false
      let modalTitle = ''
      let modalContent = ''
      
      if (!silent) {
        loading.value = true
        loadingMessage.value = '正在检查应用更新...'
      }
      
      try {
        // 从 GitHub API 获取最新版本信息
        const response = await fetch('https://api.github.com/repos/dqzboy/brew-manager/releases/latest', {
          headers: {
            'Accept': 'application/vnd.github.v3+json'
          }
        })
        
        if (!response.ok) {
          if (response.status === 404) {
            // 没有发布版本时，显示详细说明弹窗
            if (!silent) {
              shouldShowModal = true
              modalTitle = '暂无可用版本'
              modalContent = '当前 GitHub 仓库还没有发布任何 Release 版本。\n\n' +
                '当前运行版本：' + appVersion.value + '\n\n' +
                '💡 请稍后再试或访问 GitHub 查看最新动态。'
            }
            return
          }
          throw new Error(`HTTP ${response.status}: ${response.statusText}`)
        }
        
        const data = await response.json()
        
        const latestVersion = data.tag_name || data.name
        const currentVersion = appVersion.value
        
        // 比较版本号
        const isNewer = compareVersions(latestVersion, currentVersion)
        
        if (latestVersion && isNewer) {
          // 有新版本
          const publishedDate = new Date(data.published_at).toLocaleDateString('zh-CN')
          const asset = data.assets.find(a => 
            a.name.includes('darwin') || 
            a.name.includes('mac') || 
            a.name.includes('Brew-Manager') ||
            a.name.includes('brew-manager')
          )
          const fileSize = asset ? (asset.size / 1024 / 1024).toFixed(2) + ' MB' : '未知'
          
          updateDialog.value = {
            show: true,
            version: latestVersion,
            description: data.body || '查看 GitHub 了解更新详情',
            publishedAt: publishedDate,
            fileSize: fileSize,
            downloadUrl: asset ? asset.browser_download_url : data.html_url
          }
        } else {
          if (!silent) {
            shouldShowModal = true
            modalTitle = '已是最新版本'
            modalContent = '当前版本：' + currentVersion + '\n' +
              '最新版本：' + latestVersion + '\n\n' +
              '🎉 您使用的已经是最新版本！'
          }
        }
      } catch (error) {
        console.error('Update check error:', error)
        if (!silent) {
          shouldShowModal = true
          modalTitle = '检查更新失败'
          modalContent = '无法连接到 GitHub 服务器。\n\n' +
            '错误信息：' + error.message + '\n\n' +
            '请检查：\n' +
            '1. 网络连接是否正常\n' +
            '2. 是否能访问 GitHub\n' +
            '3. 防火墙设置\n\n' +
            '您也可以手动访问以下地址查看更新：\n' +
            'https://github.com/dqzboy/brew-manager/releases'
        }
      } finally {
        if (!silent) {
          loading.value = false
        }
        
        // 在 loading 关闭后显示弹窗
        if (shouldShowModal) {
          setTimeout(() => {
            console.log('Showing modal:', modalTitle)
            showModal(modalTitle, modalContent)
          }, 200)
        }
      }
    }

    // 版本号比较函数
    const compareVersions = (v1, v2) => {
      // 移除 'v' 前缀
      const ver1 = v1.replace(/^v/, '').split('.').map(Number)
      const ver2 = v2.replace(/^v/, '').split('.').map(Number)
      
      for (let i = 0; i < Math.max(ver1.length, ver2.length); i++) {
        const num1 = ver1[i] || 0
        const num2 = ver2[i] || 0
        
        if (num1 > num2) return true
        if (num1 < num2) return false
      }
      
      return false // 版本相同
    }

    const closeUpdateDialog = () => {
      updateDialog.value.show = false
    }

    const downloadUpdate = () => {
      if (updateDialog.value.downloadUrl) {
        BrowserOpenURL(updateDialog.value.downloadUrl)
        closeUpdateDialog()
        showToast('已在浏览器中打开下载页面', 'success')
      }
    }

    const openURL = (url) => {
      BrowserOpenURL(url)
    }

    // 推广轮播控制
    const scrollPromoLeft = () => {
      if (promoCarousel.value) {
        const scrollAmount = 340 // 卡片宽度 + gap
        promoCarousel.value.scrollBy({
          left: -scrollAmount,
          behavior: 'smooth'
        })
        updatePromoScrollPosition()
      }
    }

    const scrollPromoRight = () => {
      if (promoCarousel.value) {
        const scrollAmount = 340
        promoCarousel.value.scrollBy({
          left: scrollAmount,
          behavior: 'smooth'
        })
        updatePromoScrollPosition()
      }
    }

    const updatePromoScrollPosition = () => {
      setTimeout(() => {
        if (promoCarousel.value) {
          promoScrollPosition.value = promoCarousel.value.scrollLeft
          const maxScroll = promoCarousel.value.scrollWidth - promoCarousel.value.clientWidth
          promoAtEnd.value = promoScrollPosition.value >= maxScroll - 10
        }
      }, 100)
    }

    // 返回顶部相关
    const scrollToTop = () => {
      // 尝试找到当前标签页的滚动容器
      const packageList = document.querySelector('.tab-content-wrapper .package-list')
      const settingsContent = document.querySelector('.settings-content')
      
      if (packageList) {
        packageList.scrollTo({
          top: 0,
          behavior: 'smooth'
        })
      } else if (settingsContent) {
        settingsContent.scrollTo({
          top: 0,
          behavior: 'smooth'
        })
      }
    }

    const handleScroll = (e) => {
      showBackToTop.value = e.target.scrollTop > 300
    }

    const setupScrollListeners = () => {
      // 为包列表容器添加滚动监听
      const packageLists = document.querySelectorAll('.package-list')
      packageLists.forEach(list => {
        list.removeEventListener('scroll', handleScroll)
        list.addEventListener('scroll', handleScroll)
      })

      // 为设置页面添加滚动监听
      const settingsContent = document.querySelector('.settings-content')
      if (settingsContent) {
        settingsContent.removeEventListener('scroll', handleScroll)
        settingsContent.addEventListener('scroll', handleScroll)
      }
    }

    onMounted(async () => {
      initTheme()
      loadUpdateSettings()
      
      // 获取当前应用版本号
      await fetchCurrentVersion()
      
      loadInstalledPackages()
      // 默认加载可更新软件包列表
      loadOutdatedPackages()

      // 初始化滚动监听
      setTimeout(() => {
        setupScrollListeners()
      }, 100)

      // 如果启用了自动检查更新，在启动后5秒静默检查
      if (autoCheckUpdate.value) {
        setTimeout(() => {
          checkForUpdates(true) // silent=true，只在发现新版本时弹窗
        }, 5000)
      }
    })

    return {
      currentTab,
      loading,
      loadingMessage,
      installedPackages,
      searchResults,
      outdatedPackages,
      searchQuery,
      filterQuery,
      brewCoreUrl,
      brewBottleUrl,
      currentTheme,
      themeTooltip,
      showBackToTop,
      autoCheckUpdate,
      updateMode,
      appVersion,
      promoCarousel,
      promoScrollPosition,
      promoAtEnd,
      currentPackages,
      toast,
      modal,
      confirmDialog,
      operationModals,
      updateDialog,
      logsContainer,
      switchTab,
      searchPackages,
      installPackage,
      confirmUninstall,
      confirmReinstall,
      updatePackage,
      updateAll,
      updateBrew,
      cleanupBrew,
      showPackageInfo,
      closeModal,
      closeConfirmDialog,
      confirmAction,
      loadBrewSources,
      confirmSetMirror,
      closeOperationModal,
      toggleOperationMinimize,
      setTheme,
      cycleTheme,
      scrollToTop,
      setupScrollListeners,
      checkForUpdates,
      saveUpdateSettings,
      closeUpdateDialog,
      downloadUpdate,
      openURL,
      scrollPromoLeft,
      scrollPromoRight
    }
  }
}
</script>

