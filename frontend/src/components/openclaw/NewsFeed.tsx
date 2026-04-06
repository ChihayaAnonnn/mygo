import React, { useState } from "react";
import { Newspaper, ExternalLink, Clock, Tag, ChevronRight, RefreshCw } from "lucide-react";

interface NewsItem {
  id: string;
  title: string;
  source: string;
  summary: string;
  url: string;
  category: string;
  published_at: string;
  relevance: number;
}

interface NewsFeedProps {
  news?: NewsItem[];
  maxItems?: number;
  onRefresh?: () => void;
}

const NewsFeed: React.FC<NewsFeedProps> = ({
  news = [],
  maxItems = 5,
  onRefresh
}) => {
  const [expandedId, setExpandedId] = useState<string | null>(null);

  // 模拟数据
  const defaultNews: NewsItem[] = [
    {
      id: 'news_001',
      title: 'AI领域最新突破：新型语言模型发布',
      source: '科技新闻',
      summary: '研究人员在自然语言处理方面取得重大进展，新型模型在多项基准测试中表现优异...',
      url: 'https://example.com/news/001',
      category: 'technology',
      published_at: new Date().toISOString(),
      relevance: 0.9
    },
    {
      id: 'news_002',
      title: '气候变化研究新发现',
      source: '科学杂志',
      summary: '最新研究表明全球气温上升速度超出预期，科学家呼吁采取紧急行动...',
      url: 'https://example.com/news/002',
      category: 'science',
      published_at: new Date(Date.now() - 86400000).toISOString(),
      relevance: 0.8
    },
    {
      id: 'news_003',
      title: '太空探索新里程碑',
      source: '航天周刊',
      summary: '新型探测器成功登陆火星，传回高清图像和科学数据...',
      url: 'https://example.com/news/003',
      category: 'space',
      published_at: new Date(Date.now() - 172800000).toISOString(),
      relevance: 0.7
    }
  ];

  const displayNews = news.length > 0 ? news.slice(0, maxItems) : defaultNews;

  // 格式化时间
  const formatTime = (timestamp: string) => {
    const date = new Date(timestamp);
    const now = new Date();
    const diff = now.getTime() - date.getTime();
    
    if (diff < 3600000) { // 1小时内
      const minutes = Math.floor(diff / 60000);
      return `${minutes}分钟前`;
    } else if (diff < 86400000) { // 1天内
      const hours = Math.floor(diff / 3600000);
      return `${hours}小时前`;
    } else {
      return date.toLocaleDateString();
    }
  };

  // 获取类别颜色
  const getCategoryColor = (category: string) => {
    switch (category.toLowerCase()) {
      case 'technology':
        return 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300';
      case 'science':
        return 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300';
      case 'space':
        return 'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-300';
      case 'business':
        return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300';
      case 'entertainment':
        return 'bg-pink-100 text-pink-800 dark:bg-pink-900 dark:text-pink-300';
      default:
        return 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-300';
    }
  };

  // 获取类别图标
  const getCategoryIcon = (category: string) => {
    switch (category.toLowerCase()) {
      case 'technology':
        return '💻';
      case 'science':
        return '🔬';
      case 'space':
        return '🚀';
      case 'business':
        return '💼';
      case 'entertainment':
        return '🎬';
      default:
        return '📰';
    }
  };

  return (
    <div className="bg-white dark:bg-stone-800 rounded-2xl shadow-xl p-6">
      {/* 头部 */}
      <div className="flex items-center justify-between mb-6">
        <div className="flex items-center">
          <Newspaper className="w-6 h-6 text-orange-500 mr-3" />
          <div>
            <h3 className="text-xl font-bold text-stone-900 dark:text-stone-100">
              今日新闻
            </h3>
            <p className="text-sm text-stone-600 dark:text-stone-400">
              OpenClaw收集的最新资讯
            </p>
          </div>
        </div>
        
        {onRefresh && (
          <button
            onClick={onRefresh}
            className="p-2 rounded-lg hover:bg-stone-100 dark:hover:bg-stone-700 transition-colors"
            title="刷新新闻"
          >
            <RefreshCw className="w-5 h-5 text-stone-500" />
          </button>
        )}
      </div>

      {/* 新闻列表 */}
      <div className="space-y-4">
        {displayNews.map((item) => (
          <div
            key={item.id}
            className={`bg-stone-50 dark:bg-stone-900 rounded-xl p-4 border border-stone-200 dark:border-stone-700 transition-all ${
              expandedId === item.id ? 'ring-2 ring-orange-500' : 'hover:border-orange-300 dark:hover:border-orange-800'
            }`}
          >
            <div className="flex items-start justify-between">
              <div className="flex-1">
                {/* 标题和来源 */}
                <div className="flex items-start mb-2">
                  <div className="mr-3 text-2xl">
                    {getCategoryIcon(item.category)}
                  </div>
                  <div className="flex-1">
                    <h4 className="font-bold text-stone-900 dark:text-stone-100 mb-1">
                      {item.title}
                    </h4>
                    <div className="flex items-center text-sm text-stone-600 dark:text-stone-400">
                      <span className="mr-3">{item.source}</span>
                      <Clock className="w-3 h-3 mr-1" />
                      <span>{formatTime(item.published_at)}</span>
                    </div>
                  </div>
                </div>

                {/* 摘要 */}
                <p className="text-stone-700 dark:text-stone-300 mb-3 line-clamp-2">
                  {item.summary}
                </p>

                {/* 标签和操作 */}
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-2">
                    <span className={`px-2 py-1 rounded-full text-xs font-medium ${getCategoryColor(item.category)}`}>
                      <Tag className="w-3 h-3 inline mr-1" />
                      {item.category}
                    </span>
                    <span className="text-xs text-stone-500 dark:text-stone-400">
                      相关性: {(item.relevance * 100).toFixed(0)}%
                    </span>
                  </div>

                  <div className="flex items-center space-x-2">
                    <button
                      onClick={() => setExpandedId(expandedId === item.id ? null : item.id)}
                      className="text-sm text-orange-600 dark:text-orange-400 hover:text-orange-700 dark:hover:text-orange-300"
                    >
                      {expandedId === item.id ? '收起' : '详情'}
                    </button>
                    <a
                      href={item.url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-sm text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300 flex items-center"
                    >
                      阅读原文
                      <ExternalLink className="w-3 h-3 ml-1" />
                    </a>
                  </div>
                </div>

                {/* 展开的详情 */}
                {expandedId === item.id && (
                  <div className="mt-4 pt-4 border-t border-stone-200 dark:border-stone-700">
                    <div className="prose dark:prose-invert max-w-none">
                      <p className="text-stone-700 dark:text-stone-300">
                        {item.summary} 这里是更详细的新闻内容。OpenClaw通过智能算法从多个来源收集和整理这些信息，确保您获得最新、最相关的资讯。
                      </p>
                      <div className="mt-3 text-sm text-stone-500 dark:text-stone-400">
                        <p>📊 数据来源: {item.source}</p>
                        <p>⏰ 发布时间: {new Date(item.published_at).toLocaleString()}</p>
                        <p>🎯 相关性评分: {(item.relevance * 100).toFixed(1)}%</p>
                      </div>
                    </div>
                  </div>
                )}
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* 底部 */}
      <div className="mt-6 pt-6 border-t border-stone-200 dark:border-stone-700">
        <div className="flex items-center justify-between text-sm text-stone-600 dark:text-stone-400">
          <div className="flex items-center">
            <div className="flex items-center mr-4">
              <div className="w-2 h-2 bg-green-500 rounded-full mr-1" />
              <span>实时更新</span>
            </div>
            <span>共 {displayNews.length} 条新闻</span>
          </div>
          
          <button className="flex items-center text-orange-600 dark:text-orange-400 hover:text-orange-700 dark:hover:text-orange-300">
            查看更多
            <ChevronRight className="w-4 h-4 ml-1" />
          </button>
        </div>
      </div>
    </div>
  );
};

export default NewsFeed;