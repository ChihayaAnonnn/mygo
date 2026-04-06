import React, { useState, useEffect } from "react";
import { 
  Wifi, WifiOff, Clock, Activity,
  Battery, BatteryCharging, Cpu,
  Server, CheckCircle, XCircle,
  AlertCircle, Moon, Sun
} from "lucide-react";

interface StatusIndicatorProps {
  status: 'online' | 'busy' | 'away' | 'sleeping';
  activity: string;
  connected: boolean;
  showDetails?: boolean;
  className?: string;
}

const StatusIndicator: React.FC<StatusIndicatorProps> = ({
  status,
  activity,
  connected,
  showDetails = true,
  className = ""
}) => {
  const [currentTime, setCurrentTime] = useState<string>('');
  const [uptime, setUptime] = useState<string>('');
  const [cpuUsage, setCpuUsage] = useState<number>(0);
  const [memoryUsage, setMemoryUsage] = useState<number>(0);

  // 更新时间
  useEffect(() => {
    const updateTime = () => {
      const now = new Date();
      setCurrentTime(now.toLocaleTimeString([], { 
        hour: '2-digit', 
        minute: '2-digit',
        second: '2-digit'
      }));
    };

    updateTime();
    const interval = setInterval(updateTime, 1000);

    return () => clearInterval(interval);
  }, []);

  // 模拟系统指标
  useEffect(() => {
    const updateMetrics = () => {
      // 模拟运行时间（假设从某个时间开始）
      const startTime = Date.now() - 3600000; // 1小时前
      const uptimeMs = Date.now() - startTime;
      const hours = Math.floor(uptimeMs / 3600000);
      const minutes = Math.floor((uptimeMs % 3600000) / 60000);
      setUptime(`${hours}h ${minutes}m`);

      // 模拟CPU和内存使用率
      setCpuUsage(20 + Math.random() * 30);
      setMemoryUsage(40 + Math.random() * 30);
    };

    updateMetrics();
    const interval = setInterval(updateMetrics, 5000);

    return () => clearInterval(interval);
  }, []);

  // 获取状态颜色
  const getStatusColor = () => {
    switch (status) {
      case 'online':
        return connected ? 'bg-green-500' : 'bg-yellow-500';
      case 'busy':
        return 'bg-yellow-500';
      case 'away':
        return 'bg-gray-500';
      case 'sleeping':
        return 'bg-blue-500';
      default:
        return 'bg-gray-500';
    }
  };

  // 获取状态文本
  const getStatusText = () => {
    switch (status) {
      case 'online':
        return connected ? '在线' : '连接中';
      case 'busy':
        return '忙碌';
      case 'away':
        return '离开';
      case 'sleeping':
        return '睡眠';
      default:
        return '未知';
    }
  };

  // 获取状态图标
  const getStatusIcon = () => {
    switch (status) {
      case 'online':
        return connected ? <CheckCircle className="w-4 h-4" /> : <AlertCircle className="w-4 h-4" />;
      case 'busy':
        return <Activity className="w-4 h-4" />;
      case 'away':
        return <Moon className="w-4 h-4" />;
      case 'sleeping':
        return <Moon className="w-4 h-4" />;
      default:
        return <Server className="w-4 h-4" />;
    }
  };

  // 获取连接状态
  const getConnectionStatus = () => {
    if (!connected) {
      return {
        text: '连接断开',
        icon: <WifiOff className="w-4 h-4" />,
        color: 'text-red-500'
      };
    }

    switch (status) {
      case 'online':
        return {
          text: '连接正常',
          icon: <Wifi className="w-4 h-4" />,
          color: 'text-green-500'
        };
      case 'busy':
        return {
          text: '处理中',
          icon: <Activity className="w-4 h-4" />,
          color: 'text-yellow-500'
        };
      default:
        return {
          text: '连接正常',
          icon: <Wifi className="w-4 h-4" />,
          color: 'text-green-500'
        };
    }
  };

  const connectionStatus = getConnectionStatus();

  return (
    <div className={`flex flex-col space-y-4 ${className}`}>
      {/* 主要状态指示器 */}
      <div className="flex items-center space-x-4">
        {/* 状态灯 */}
        <div className="relative">
          <div className={`w-3 h-3 rounded-full ${getStatusColor()} animate-pulse`} />
          <div 
            className={`absolute inset-0 rounded-full ${getStatusColor()} animate-ping`}
            style={{ animationDuration: '2s' }}
          />
        </div>

        {/* 状态文本 */}
        <div>
          <div className="flex items-center space-x-2">
            <span className="font-medium text-stone-900 dark:text-stone-100">
              {getStatusText()}
            </span>
            <div className={`p-1 rounded ${getStatusColor().replace('bg-', 'bg-').replace('500', '100')} dark:${getStatusColor().replace('bg-', 'bg-').replace('500', '900/30')}`}>
              {getStatusIcon()}
            </div>
          </div>
          <div className="text-sm text-stone-600 dark:text-stone-400">
            {activity}
          </div>
        </div>
      </div>

      {/* 详细状态信息 */}
      {showDetails && (
        <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
          {/* 连接状态 */}
          <div className="bg-stone-50 dark:bg-stone-900 rounded-lg p-3">
            <div className="flex items-center justify-between mb-1">
              <div className={`flex items-center ${connectionStatus.color}`}>
                {connectionStatus.icon}
              </div>
              <div className="text-xs text-stone-500 dark:text-stone-400">
                连接
              </div>
            </div>
            <div className="font-medium text-stone-900 dark:text-stone-100">
              {connectionStatus.text}
            </div>
            <div className="text-xs text-stone-500 dark:text-stone-400 mt-1">
              {connected ? 'WebSocket已连接' : '正在重连...'}
            </div>
          </div>

          {/* 系统时间 */}
          <div className="bg-stone-50 dark:bg-stone-900 rounded-lg p-3">
            <div className="flex items-center justify-between mb-1">
              <div className="text-blue-500">
                <Clock className="w-4 h-4" />
              </div>
              <div className="text-xs text-stone-500 dark:text-stone-400">
                时间
              </div>
            </div>
            <div className="font-medium text-stone-900 dark:text-stone-100 font-mono">
              {currentTime}
            </div>
            <div className="text-xs text-stone-500 dark:text-stone-400 mt-1">
              运行: {uptime}
            </div>
          </div>

          {/* CPU使用率 */}
          <div className="bg-stone-50 dark:bg-stone-900 rounded-lg p-3">
            <div className="flex items-center justify-between mb-1">
              <div className="text-purple-500">
                <Cpu className="w-4 h-4" />
              </div>
              <div className="text-xs text-stone-500 dark:text-stone-400">
                CPU
              </div>
            </div>
            <div className="font-medium text-stone-900 dark:text-stone-100">
              {cpuUsage.toFixed(1)}%
            </div>
            <div className="mt-1">
              <div className="h-1.5 bg-stone-200 dark:bg-stone-700 rounded-full overflow-hidden">
                <div 
                  className="h-full bg-gradient-to-r from-purple-400 to-pink-400 rounded-full"
                  style={{ width: `${cpuUsage}%` }}
                />
              </div>
            </div>
          </div>

          {/* 内存使用率 */}
          <div className="bg-stone-50 dark:bg-stone-900 rounded-lg p-3">
            <div className="flex items-center justify-between mb-1">
              <div className="text-green-500">
                <Server className="w-4 h-4" />
              </div>
              <div className="text-xs text-stone-500 dark:text-stone-400">
                内存
              </div>
            </div>
            <div className="font-medium text-stone-900 dark:text-stone-100">
              {memoryUsage.toFixed(1)}%
            </div>
            <div className="mt-1">
              <div className="h-1.5 bg-stone-200 dark:bg-stone-700 rounded-full overflow-hidden">
                <div 
                  className="h-full bg-gradient-to-r from-green-400 to-cyan-400 rounded-full"
                  style={{ width: `${memoryUsage}%` }}
                />
              </div>
            </div>
          </div>
        </div>
      )}

      {/* 状态详情展开 */}
      <div className="pt-4 border-t border-stone-200 dark:border-stone-700">
        <div className="flex items-center justify-between">
          <div className="text-sm text-stone-600 dark:text-stone-400">
            系统状态详情
          </div>
          
          <div className="flex items-center space-x-4">
            {/* 电池状态（模拟） */}
            <div className="flex items-center text-sm">
              <BatteryCharging className="w-4 h-4 text-green-500 mr-1" />
              <span className="text-stone-700 dark:text-stone-300">充电中</span>
              <span className="mx-1">·</span>
              <span className="text-stone-500 dark:text-stone-400">85%</span>
            </div>

            {/* 网络延迟（模拟） */}
            <div className="flex items-center text-sm">
              <Activity className="w-4 h-4 text-blue-500 mr-1" />
              <span className="text-stone-700 dark:text-stone-300">延迟</span>
              <span className="mx-1">·</span>
              <span className="text-stone-500 dark:text-stone-400">32ms</span>
            </div>
          </div>
        </div>

        {/* 状态进度条 */}
        <div className="mt-3 space-y-2">
          {/* 响应时间 */}
          <div>
            <div className="flex justify-between text-xs text-stone-600 dark:text-stone-400 mb-1">
              <span>响应时间</span>
              <span>快速</span>
            </div>
            <div className="h-1.5 bg-stone-200 dark:bg-stone-700 rounded-full overflow-hidden">
              <div className="h-full bg-gradient-to-r from-green-400 to-emerald-400 rounded-full w-3/4" />
            </div>
          </div>

          {/* 任务队列 */}
          <div>
            <div className="flex justify-between text-xs text-stone-600 dark:text-stone-400 mb-1">
              <span>任务队列</span>
              <span>3个等待</span>
            </div>
            <div className="h-1.5 bg-stone-200 dark:bg-stone-700 rounded-full overflow-hidden">
              <div className="h-full bg-gradient-to-r from-yellow-400 to-orange-400 rounded-full w-1/4" />
            </div>
          </div>

          {/* 错误率 */}
          <div>
            <div className="flex justify-between text-xs text-stone-600 dark:text-stone-400 mb-1">
              <span>错误率</span>
              <span>0.2%</span>
            </div>
            <div className="h-1.5 bg-stone-200 dark:bg-stone-700 rounded-full overflow-hidden">
              <div className="h-full bg-gradient-to-r from-red-400 to-pink-400 rounded-full w-1" />
            </div>
          </div>
        </div>
      </div>

      {/* 状态提示 */}
      <div className={`p-3 rounded-lg text-sm ${
        status === 'online' && connected 
          ? 'bg-green-50 dark:bg-green-900/20 text-green-800 dark:text-green-300 border border-green-200 dark:border-green-800'
          : status === 'busy'
          ? 'bg-yellow-50 dark:bg-yellow-900/20 text-yellow-800 dark:text-yellow-300 border border-yellow-200 dark:border-yellow-800'
          : status === 'sleeping'
          ? 'bg-blue-50 dark:bg-blue-900/20 text-blue-800 dark:text-blue-300 border border-blue-200 dark:border-blue-800'
          : 'bg-gray-50 dark:bg-gray-900/20 text-gray-800 dark:text-gray-300 border border-gray-200 dark:border-gray-800'
      }`}>
        <div className="flex items-start">
          {status === 'online' && connected ? (
            <CheckCircle className="w-4 h-4 mt-0.5 mr-2 flex-shrink-0" />
          ) : status === 'busy' ? (
            <Activity className="w-4 h-4 mt-0.5 mr-2 flex-shrink-0" />
          ) : (
            <AlertCircle className="w-4 h-4 mt-0.5 mr-2 flex-shrink-0" />
          )}
          
          <div>
            {status === 'online' && connected ? (
              <p>OpenClaw运行正常，随时为您服务</p>
            ) : status === 'busy' ? (
              <p>OpenClaw正在处理任务，响应可能稍有延迟</p>
            ) : status === 'sleeping' ? (
              <p>OpenClaw处于睡眠模式，唤醒后立即响应</p>
            ) : (
              <p>OpenClaw暂时离开，返回后继续服务</p>
            )}
            
            <div className="mt-1 text-xs opacity-80">
              最后活动: {new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default StatusIndicator;