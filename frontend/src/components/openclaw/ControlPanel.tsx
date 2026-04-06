import React, { useState, useRef, useEffect } from "react";
import { 
  Send, Mic, Keyboard, Zap, 
  History, Star, Settings, HelpCircle,
  ChevronRight, Search, Terminal,
  MessageSquare, Bot, Globe, Cpu
} from "lucide-react";

interface QuickAction {
  id: string;
  label: string;
  icon: React.ReactNode;
  description?: string;
}

interface CommandHistory {
  id: string;
  command: string;
  timestamp: number;
  success: boolean;
  response?: string;
}

interface ControlPanelProps {
  onCommand: (command: string, params?: any) => void;
  quickActions?: QuickAction[];
  showHistory?: boolean;
  className?: string;
}

const ControlPanel: React.FC<ControlPanelProps> = ({
  onCommand,
  quickActions = [],
  showHistory = true,
  className = ""
}) => {
  const [input, setInput] = useState("");
  const [isListening, setIsListening] = useState(false);
  const [commandHistory, setCommandHistory] = useState<CommandHistory[]>([]);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [selectedSuggestion, setSelectedSuggestion] = useState(0);
  const inputRef = useRef<HTMLInputElement>(null);
  const suggestionsRef = useRef<HTMLDivElement>(null);

  // 默认快捷动作
  const defaultQuickActions: QuickAction[] = [
    {
      id: 'weather',
      label: '查看天气',
      icon: <Globe className="w-5 h-5" />,
      description: '获取当前天气信息'
    },
    {
      id: 'news',
      label: '今日新闻',
      icon: <MessageSquare className="w-5 h-5" />,
      description: '浏览最新新闻'
    },
    {
      id: 'ai-updates',
      label: 'AI动态',
      icon: <Cpu className="w-5 h-5" />,
      description: '查看AI领域最新进展'
    },
    {
      id: 'joke',
      label: '讲个笑话',
      icon: <Bot className="w-5 h-5" />,
      description: '让OpenClaw讲个笑话'
    }
  ];

  const actions = quickActions.length > 0 ? quickActions : defaultQuickActions;

  // 命令建议
  const commandSuggestions = [
    { command: 'show weather', description: '显示天气信息' },
    { command: 'tell me a joke', description: '讲个笑话' },
    { command: 'what\'s new in AI', description: 'AI领域最新动态' },
    { command: 'play music', description: '播放音乐' },
    { command: 'set reminder', description: '设置提醒' },
    { command: 'search web for', description: '搜索网络信息' },
    { command: 'translate', description: '翻译文本' },
    { command: 'calculate', description: '计算数学表达式' },
  ];

  // 加载历史记录
  useEffect(() => {
    const savedHistory = localStorage.getItem('openclaw_command_history');
    if (savedHistory) {
      try {
        setCommandHistory(JSON.parse(savedHistory).slice(0, 10));
      } catch (e) {
        console.error('Failed to parse command history:', e);
      }
    }
  }, []);

  // 保存历史记录
  const saveToHistory = (command: string, success: boolean, response?: string) => {
    const newEntry: CommandHistory = {
      id: Date.now().toString(),
      command,
      timestamp: Date.now(),
      success,
      response
    };

    const updatedHistory = [newEntry, ...commandHistory].slice(0, 20);
    setCommandHistory(updatedHistory);
    
    try {
      localStorage.setItem('openclaw_command_history', JSON.stringify(updatedHistory));
    } catch (e) {
      console.error('Failed to save command history:', e);
    }
  };

  // 处理命令提交
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!input.trim()) return;

    // 执行命令
    onCommand(input.trim());
    
    // 保存到历史记录
    saveToHistory(input.trim(), true);
    
    // 清空输入
    setInput("");
    setShowSuggestions(false);
    
    // 聚焦输入框
    inputRef.current?.focus();
  };

  // 处理快捷动作
  const handleQuickAction = (actionId: string) => {
    const action = actions.find(a => a.id === actionId);
    if (action) {
      onCommand(action.label);
      saveToHistory(action.label, true);
    }
  };

  // 处理语音输入
  const handleVoiceInput = () => {
    if (!('webkitSpeechRecognition' in window || 'SpeechRecognition' in window)) {
      alert('您的浏览器不支持语音识别功能');
      return;
    }

    setIsListening(true);
    
    // 这里应该实现实际的语音识别
    // 由于浏览器API限制，这里只做演示
    setTimeout(() => {
      setIsListening(false);
      const demoCommands = [
        "今天天气怎么样",
        "讲个笑话",
        "搜索人工智能最新进展",
        "播放一些轻松的音乐"
      ];
      const randomCommand = demoCommands[Math.floor(Math.random() * demoCommands.length)];
      setInput(randomCommand);
      inputRef.current?.focus();
    }, 2000);
  };

  // 处理键盘导航
  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (!showSuggestions) return;

    switch (e.key) {
      case 'ArrowDown':
        e.preventDefault();
        setSelectedSuggestion(prev => 
          prev < commandSuggestions.length - 1 ? prev + 1 : prev
        );
        break;
      case 'ArrowUp':
        e.preventDefault();
        setSelectedSuggestion(prev => prev > 0 ? prev - 1 : prev);
        break;
      case 'Enter':
        if (showSuggestions && selectedSuggestion >= 0) {
          e.preventDefault();
          const suggestion = commandSuggestions[selectedSuggestion];
          setInput(suggestion.command);
          setShowSuggestions(false);
        }
        break;
      case 'Escape':
        setShowSuggestions(false);
        break;
    }
  };

  // 点击外部关闭建议
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        suggestionsRef.current && 
        !suggestionsRef.current.contains(event.target as Node) &&
        inputRef.current !== event.target
      ) {
        setShowSuggestions(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  // 格式化时间
  const formatTime = (timestamp: number) => {
    const date = new Date(timestamp);
    const now = new Date();
    const diff = now.getTime() - timestamp;
    
    if (diff < 60000) return '刚刚';
    if (diff < 3600000) return `${Math.floor(diff / 60000)}分钟前`;
    if (diff < 86400000) return `${Math.floor(diff / 3600000)}小时前`;
    
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
  };

  return (
    <div className={`bg-white dark:bg-stone-800 rounded-2xl shadow-xl p-6 ${className}`}>
      {/* 面板标题 */}
      <div className="flex items-center justify-between mb-6">
        <div>
          <h3 className="text-xl font-bold text-stone-900 dark:text-stone-100 flex items-center">
            <Terminal className="w-6 h-6 text-orange-500 mr-2" />
            控制面板
          </h3>
          <p className="text-sm text-stone-600 dark:text-stone-400 mt-1">
            与OpenClaw交互，发送命令或使用快捷操作
          </p>
        </div>
        
        <div className="flex items-center space-x-2">
          <button
            className="p-2 rounded-lg hover:bg-stone-100 dark:hover:bg-stone-700 transition-colors"
            title="帮助"
          >
            <HelpCircle className="w-5 h-5 text-stone-500" />
          </button>
          <button
            className="p-2 rounded-lg hover:bg-stone-100 dark:hover:bg-stone-700 transition-colors"
            title="设置"
          >
            <Settings className="w-5 h-5 text-stone-500" />
          </button>
        </div>
      </div>

      {/* 命令输入区域 */}
      <div className="mb-6">
        <form onSubmit={handleSubmit} className="relative">
          <div className="flex items-center">
            <div className="flex-1 relative">
              <input
                ref={inputRef}
                type="text"
                value={input}
                onChange={(e) => {
                  setInput(e.target.value);
                  setShowSuggestions(e.target.value.length > 0);
                  setSelectedSuggestion(0);
                }}
                onKeyDown={handleKeyDown}
                onFocus={() => setShowSuggestions(input.length > 0)}
                placeholder="输入命令，例如：今天天气怎么样？"
                className="w-full px-4 py-3 pl-12 bg-stone-50 dark:bg-stone-900 border border-stone-200 dark:border-stone-700 rounded-xl focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent text-stone-900 dark:text-stone-100"
              />
              
              <div className="absolute left-4 top-1/2 transform -translate-y-1/2">
                <Keyboard className="w-5 h-5 text-stone-400" />
              </div>
              
              {/* 命令建议 */}
              {showSuggestions && commandSuggestions.length > 0 && (
                <div 
                  ref={suggestionsRef}
                  className="absolute top-full left-0 right-0 mt-1 bg-white dark:bg-stone-800 border border-stone-200 dark:border-stone-700 rounded-xl shadow-lg z-10 overflow-hidden"
                >
                  {commandSuggestions.map((suggestion, index) => (
                    <button
                      key={index}
                      type="button"
                      onClick={() => {
                        setInput(suggestion.command);
                        setShowSuggestions(false);
                        inputRef.current?.focus();
                      }}
                      className={`w-full text-left px-4 py-3 hover:bg-stone-100 dark:hover:bg-stone-700 transition-colors flex items-center justify-between ${
                        index === selectedSuggestion ? 'bg-stone-100 dark:bg-stone-700' : ''
                      }`}
                    >
                      <div>
                        <div className="font-medium text-stone-900 dark:text-stone-100">
                          {suggestion.command}
                        </div>
                        <div className="text-sm text-stone-500 dark:text-stone-400 mt-1">
                          {suggestion.description}
                        </div>
                      </div>
                      <ChevronRight className="w-4 h-4 text-stone-400" />
                    </button>
                  ))}
                </div>
              )}
            </div>
            
            <div className="flex ml-2">
              <button
                type="button"
                onClick={handleVoiceInput}
                disabled={isListening}
                className={`p-3 rounded-xl mr-2 transition-all ${
                  isListening 
                    ? 'bg-red-100 dark:bg-red-900/30 text-red-600 dark:text-red-400 animate-pulse' 
                    : 'bg-stone-100 dark:bg-stone-900 hover:bg-stone-200 dark:hover:bg-stone-700 text-stone-600 dark:text-stone-400'
                }`}
                title="语音输入"
              >
                <Mic className="w-5 h-5" />
              </button>
              
              <button
                type="submit"
                disabled={!input.trim()}
                className={`p-3 rounded-xl transition-all ${
                  input.trim()
                    ? 'bg-orange-500 hover:bg-orange-600 text-white' 
                    : 'bg-stone-100 dark:bg-stone-900 text-stone-400 dark:text-stone-600'
                }`}
                title="发送命令"
              >
                <Send className="w-5 h-5" />
              </button>
            </div>
          </div>
          
          {/* 输入提示 */}
          <div className="mt-2 text-xs text-stone-500 dark:text-stone-400 flex items-center">
            <Zap className="w-3 h-3 mr-1" />
            支持自然语言命令，如"天气"、"新闻"、"讲个笑话"等
          </div>
        </form>
      </div>

      {/* 快捷操作区域 */}
      <div className="mb-8">
        <h4 className="text-lg font-medium text-stone-900 dark:text-stone-100 mb-4 flex items-center">
          <Zap className="w-5 h-5 text-orange-500 mr-2" />
          快捷操作
        </h4>
        
        <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
          {actions.map((action) => (
            <button
              key={action.id}
              onClick={() => handleQuickAction(action.id)}
              className="group p-4 bg-stone-50 dark:bg-stone-900 hover:bg-orange-50 dark:hover:bg-orange-900/20 rounded-xl border border-stone-200 dark:border-stone-700 hover:border-orange-300 dark:hover:border-orange-800 transition-all duration-200"
            >
              <div className="flex flex-col items-center text-center">
                <div className="p-2 rounded-lg bg-white dark:bg-stone-800 group-hover:bg-orange-100 dark:group-hover:bg-orange-900/30 mb-3 transition-colors">
                  <div className="text-orange-500">
                    {action.icon}
                  </div>
                </div>
                <div className="font-medium text-stone-900 dark:text-stone-100 group-hover:text-orange-600 dark:group-hover:text-orange-400 transition-colors">
                  {action.label}
                </div>
                {action.description && (
                  <div className="text-xs text-stone-500 dark:text-stone-400 mt-1">
                    {action.description}
                  </div>
                )}
              </div>
            </button>
          ))}
        </div>
      </div>

      {/* 命令历史记录 */}
      {showHistory && commandHistory.length > 0 && (
        <div>
          <h4 className="text-lg font-medium text-stone-900 dark:text-stone-100 mb-4 flex items-center">
            <History className="w-5 h-5 text-orange-500 mr-2" />
            最近命令
          </h4>
          
          <div className="space-y-2 max-h-60 overflow-y-auto pr-2">
            {commandHistory.slice(0, 5).map((item) => (
              <div
                key={item.id}
                className="p-3 bg-stone-50 dark:bg-stone-900 rounded-lg border border-stone-200 dark:border-stone-700"
              >
                <div className="flex items-start justify-between">
                  <div className="flex-1">
                    <div className="flex items-center">
                      <div className={`w-2 h-2 rounded-full mr-2 ${
                        item.success ? 'bg-green-500' : 'bg-red-500'
                      }`} />
                      <div className="font-medium text-stone-900 dark:text-stone-100">
                        {item.command}
                      </div>
                    </div>
                    {item.response && (
                      <div className="mt-2 text-sm text-stone-600 dark:text-stone-400 pl-4 border-l-2 border-stone-300 dark:border-stone-600">
                        {item.response}
                      </div>
                    )}
                  </div>
                  <div className="text-xs text-stone-500 dark:text-stone-400 ml-2 whitespace-nowrap">
                    {formatTime(item.timestamp)}
                  </div>
                </div>
              </div>
            ))}
          </div>
          
          {commandHistory.length > 5 && (
            <div className="text-center mt-3">
              <button
                onClick={() => {
                  // 显示更多历史记录
                  console.log('Show more history');
                }}
                className="text-sm text-orange-600 dark:text-orange-400 hover:text-orange-700 dark:hover:text-orange-300"
              >
                查看全部 {commandHistory.length} 条记录
              </button>
            </div>
          )}
        </div>
      )}

      {/* 面板底部 */}
      <div className="mt-6 pt-6 border-t border-stone-200 dark:border-stone-700">
        <div className="flex items-center justify-between text-sm text-stone-500 dark:text-stone-400">
          <div className="flex items-center">
            <div className="flex items-center mr-4">
              <div className="w-2 h-2 bg-green-500 rounded-full mr-1" />
              <span>连接正常</span>
            </div>
            <div className="flex items-center">
              <Star className="w-3 h-3 mr-1" />
              <span>常用命令已加载</span>
            </div>
          </div>
          
          <button
            onClick={() => {
              setCommandHistory([]);
              localStorage.removeItem('openclaw_command_history');
            }}
            className="text-sm text-stone-500 dark:text-stone-400 hover:text-stone-700 dark:hover:text-stone-300"
          >
            清空历史
          </button>
        </div>
      </div>
    </div>
  );
};

export default ControlPanel;