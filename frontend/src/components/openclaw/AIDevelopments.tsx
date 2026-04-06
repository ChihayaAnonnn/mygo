import React, { useState } from "react";
import { Cpu, Zap, TrendingUp, AlertCircle, Clock, ExternalLink, ChevronRight, RefreshCw } from "lucide-react";

interface AIUpdate {
  id: string;
  title: string;
  description: string;
  source: string;
  impact: 'low' | 'medium' | 'high';
  date: string;
}

interface AIDevelopmentsProps {
  updates?: AIUpdate[];
  maxItems?: number;
  onRefresh?: () => void;
}

const AIDevelopments: React.FC<AIDevelopmentsProps> = ({
  updates = [],
  maxItems = 5,
  onRefresh
}) => {
  const [expandedId, setExpandedId] = useState<string | null>(null);

  // 模拟数据
  const defaultUpdates: AIUpdate[] = [
    {
      id: 'ai_001',
      title: '新型多模态AI模型发布',
      description: '最新研究展示了能够同时处理文本、图像和音频的多模态AI模型，在多项任务上表现优异。',
      source: 'AI研究期刊',
      impact: 'high',
      date: new Date().toISOString()
    },
    {
      id: 'ai_002',
      title: 'AI在医疗诊断中的突破',
      description: '深度学习算法在早期癌症检测方面达到专家级准确率，有望改变医疗诊断流程。',
      source: '医学AI进展',
      impact: 'high',
      date: new Date(Date.now() - 86400000).toISOString()
    },
    {
      id: 'ai_003',
      title: '开源AI工具包更新',
      description: '流行的开源机器学习框架发布重大更新，新增多项实用功能和性能优化。',
      source: '开源社区',
      impact: 'medium',
      date: new Date(Date.now() - 172800000).toISOString()
    },
    {
      id: 'ai_004',
      title: 'AI伦理框架讨论',
      description: '学术界和产业界就AI伦理标准展开深入讨论，提出新的治理框架建议。',
      source: 'AI伦理论坛',
      impact: 'medium',
      date: new Date(Date.now() - 259200000).toISOString()
    }
  ];

  const displayUpdates = updates.length > 0 ? updates.slice(0, maxItems) : defaultUpdates;

  // 格式化时间
  const formatTime = (timestamp: string) => {
    const date = new Date(timestamp);
    const now = new Date();
    const diff = now.getTime() - date.getTime();
    
    if (diff < 86400000) { // 1天内
      const hours = Math.floor(diff / 3600000);
      return `${hours}小时前`;
    } else {
      const days = Math.floor(diff / 86400000);
      return `${days}天前`;
    }
  };

  // 获取影响级别颜色
  const getImpactColor = (impact: string) => {
    switch (impact) {
      case 'high':
        return 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300';
      case 'medium':
        return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300';
      case 'low':
        return 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300';
      default:
        return 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-300';
    }
  };

  // 获取影响级别图标
  const getImpactIcon = (impact: string) => {
    switch (impact) {
      case 'high':
        return <AlertCircle className="w-4 h-4" />;
      case 'medium':
        return <TrendingUp className="w-4 h-4" />;
      case 'low':
        return <Zap className="w-4 h-4" />;
      default:
        return <Cpu className="w-4 h-4" />;
    }
  };

  // 获取影响级别文本
  const getImpactText = (impact: string) => {
    switch (impact) {
      case 'high':
        return '高影响';
      case 'medium':
        return '中影响';
      case 'low':
        return '低影响';
      default:
        return '未知';
    }
  };

  return (
    <div className="bg-white dark:bg-stone-800 rounded-2xl shadow-xl p-6">
      {/* 头部 */}
      <div className="flex items-center justify-between mb-6">
        <div className="flex items-center">
          <Cpu className="w-6 h-6 text-purple-500 mr-3" />
          <div>
            <h3 className="text-xl font-bold text-stone-900 dark:text-stone-100">
              AI发展动态
            </h3>
            <p className="text-sm text-stone-600 dark:text-stone-400">
              OpenClaw监控的最新AI进展
            </p>
          </div>
        </div>
        
        {onRefresh && (
          <button
            onClick={onRefresh}
            className="p-2 rounded-lg hover:bg-stone-100 dark:hover:bg-stone-700 transition-colors"
            title="刷新AI动态"
          >
            <RefreshCw className="w-5 h-5 text-stone-500" />
          </button>
        )}
      </div>

      {/* AI动态列表 */}
      <div className="space-y-4">
        {displayUpdates.map((update) => (
          <div
            key={update.id}
            className={`bg-stone-50 dark:bg-stone-900 rounded-xl p-4 border border-stone-200 dark:border-stone-700 transition-all ${
              expandedId === update.id ? 'ring-2 ring-purple-500' : 'hover:border-purple-300 dark:hover:border-purple-800'
            }`}
          >
            <div className="flex items-start justify-between">
              <div className="flex-1">
                {/* 标题和影响级别 */}
                <div className="flex items-start mb-3">
                  <div className={`p-2 rounded-lg mr-3 ${getImpactColor(update.impact)}`}>
                    {getImpactIcon(update.impact)}
                  </div>
                  <div className="flex-1">
                    <div className="flex items-center justify-between mb-1">
                      <h4 className="font-bold text-stone-900 dark:text-stone-100">
                        {update.title}
                      </h4>
                      <span className={`px-2 py-1 rounded-full text-xs font-medium ${getImpactColor(update.impact)} flex items-center`}>
                        {getImpactIcon(update.impact)}
                        <span className="ml-1">{getImpactText(update.impact)}</span>
                      </span>
                    </div>
                    <div className="flex items-center text-sm text-stone-600 dark:text-stone-400">
                      <span className="mr-3">{update.source}</span>
                      <Clock className="w-3 h-3 mr-1" />
                      <span>{formatTime(update.date)}</span>
                    </div>
                  </div>
                </div>

                {/* 描述 */}
                <p className="text-stone-700 dark:text-stone-300 mb-4">
                  {update.description}
                </p>

                {/* 操作按钮 */}
                <div className="flex items-center justify-between">
                  <div className="text-sm text-stone-500 dark:text-stone-400">
                    <span>AI领域 · 技术进展</span>
                  </div>

                  <div className="flex items-center space-x-3">
                    <button
                      onClick={() => setExpandedId(expandedId === update.id ? null : update.id)}
                      className="text-sm text-purple-600 dark:text-purple-400 hover:text-purple-700 dark:hover:text-purple-300"
                    >
                      {expandedId === update.id ? '收起详情' : '查看详情'}
                    </button>
                    <button className="text-sm text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300 flex items-center">
                      相关研究
                      <ExternalLink className="w-3 h-3 ml-1" />
                    </button>
                  </div>
                </div>

                {/* 展开的详情 */}
                {expandedId === update.id && (
                  <div className="mt-4 pt-4 border-t border-stone-200 dark:border-stone-700">
                    <div className="prose dark:prose-invert max-w-none">
                      <h5 className="font-bold text-stone-900 dark:text-stone-100 mb-2">
                        详细内容
                      </h5>
                      <p className="text-stone-700 dark:text-stone-300 mb-3">
                        {update.description} 这里是更详细的技术细节和分析。OpenClaw通过监控学术论文、技术博客和行业报告，为您提供最前沿的AI发展信息。
                      </p>
                      
                      <div className="grid grid-cols-2 gap-4 mt-4">
                        <div className="bg-blue-50 dark:bg-blue-900/20 p-3 rounded-lg">
                          <div className="text-sm font-medium text-blue-800 dark:text-blue-300 mb-1">
                            技术类别
                          </div>
                          <div className="text-blue-600 dark:text-blue-400">
                            机器学习 · 深度学习
                          </div>
                        </div>
                        
                        <div className="bg-green-50 dark:bg-green-900/20 p-3 rounded-lg">
                          <div className="text-sm font-medium text-green-800 dark:text-green-300 mb-1">
                            应用领域
                          </div>
                          <div className="text-green-600 dark:text-green-400">
                            医疗健康 · 自然语言处理
                          </div>
                        </div>
                      </div>
                      
                      <div className="mt-4 text-sm text-stone-500 dark:text-stone-400">
                        <p>📊 数据来源: {update.source}</p>
                        <p>⏰ 发布时间: {new Date(update.date).toLocaleDateString()}</p>
                        <p>🎯 影响评估: {getImpactText(update.impact)}级别</p>
                        <p>🔍 监控频率: 每日更新</p>
                      </div>
                    </div>
                  </div>
                )}
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* 趋势图表（简化） */}
      <div className="mt-6 bg-gradient-to-r from-purple-50 to-blue-50 dark:from-purple-900/20 dark:to-blue-900/20 rounded-xl p-4">
        <div className="flex items-center justify-between mb-3">
          <h4 className="font-medium text-stone-900 dark:text-stone-100">
            AI发展热度趋势
          </h4>
          <span className="text-sm text-stone-600 dark:text-stone-400">
            最近30天
          </span>
        </div>
        
        <div className="h-32 flex items-end space-x-1">
          {Array.from({ length: 30 }).map((_, i) => {
            const height = 20 + Math.sin(i * 0.3) * 40 + Math.random() * 20;
            const opacity = 0.5 + Math.sin(i * 0.2) * 0.3;
            
            return (
              <div
                key={i}
                className="flex-1 bg-gradient-to-t from-purple-400 to-blue-400 rounded-t"
                style={{
                  height: `${height}%`,
                  opacity: opacity
                }}
              />
            );
          })}
        </div>
        
        <div className="flex justify-between text-xs text-stone-500 dark:text-stone-400 mt-2">
          <span>30天前</span>
          <span>15天前</span>
          <span>今天</span>
        </div>
      </div>

      {/* 底部 */}
      <div className="mt-6 pt-6 border-t border-stone-200 dark:border-stone-700">
        <div className="flex items-center justify-between text-sm text-stone-600 dark:text-stone-400">
          <div className="flex items-center">
            <div className="flex items-center mr-4">
              <div className="w-2 h-2 bg-green-500 rounded-full mr-1" />
              <span>自动监控</span>
            </div>
            <span>共 {displayUpdates.length} 条动态</span>
          </div>
          
          <button className="flex items-center text-purple-600 dark:text-purple-400 hover:text-purple-700 dark:hover:text-purple-300">
            查看历史
            <ChevronRight className="w-4 h-4 ml-1" />
          </button>
        </div>
      </div>
    </div>
  );
};

export default AIDevelopments;