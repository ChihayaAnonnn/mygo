import React, { useState, useEffect } from "react";
import { 
  Sun, Cloud, CloudRain, CloudSnow, 
  Zap, Wind, Thermometer, Droplets,
  Calendar, MapPin, RefreshCw
} from "lucide-react";

interface WeatherCardProps {
  location: string;
  temperature: number;
  condition: string;
  humidity: number;
  forecast: string;
  icon: string;
  updatedAt: number;
  onRefresh?: () => void;
}

const WeatherCard: React.FC<WeatherCardProps> = ({
  location,
  temperature,
  condition,
  humidity,
  forecast,
  icon,
  updatedAt,
  onRefresh
}) => {
  const [timeSinceUpdate, setTimeSinceUpdate] = useState<string>('刚刚');
  const [isRefreshing, setIsRefreshing] = useState(false);

  // 更新时间显示
  useEffect(() => {
    const updateTimeDisplay = () => {
      const now = Date.now();
      const diff = now - updatedAt;
      
      if (diff < 60000) { // 1分钟内
        setTimeSinceUpdate('刚刚');
      } else if (diff < 3600000) { // 1小时内
        const minutes = Math.floor(diff / 60000);
        setTimeSinceUpdate(`${minutes}分钟前`);
      } else if (diff < 86400000) { // 1天内
        const hours = Math.floor(diff / 3600000);
        setTimeSinceUpdate(`${hours}小时前`);
      } else {
        const days = Math.floor(diff / 86400000);
        setTimeSinceUpdate(`${days}天前`);
      }
    };

    updateTimeDisplay();
    const interval = setInterval(updateTimeDisplay, 60000); // 每分钟更新一次

    return () => clearInterval(interval);
  }, [updatedAt]);

  // 获取天气图标
  const getWeatherIcon = () => {
    const iconName = icon.toLowerCase();
    
    if (iconName.includes('sun') || iconName.includes('clear')) {
      return <Sun className="w-12 h-12 text-yellow-500" />;
    } else if (iconName.includes('cloud')) {
      return <Cloud className="w-12 h-12 text-gray-500" />;
    } else if (iconName.includes('rain')) {
      return <CloudRain className="w-12 h-12 text-blue-500" />;
    } else if (iconName.includes('snow')) {
      return <CloudSnow className="w-12 h-12 text-blue-300" />;
    } else if (iconName.includes('storm') || iconName.includes('thunder')) {
      return <Zap className="w-12 h-12 text-purple-500" />;
    } else if (iconName.includes('wind')) {
      return <Wind className="w-12 h-12 text-gray-400" />;
    } else {
      return <Sun className="w-12 h-12 text-gray-400" />;
    }
  };

  // 获取温度颜色
  const getTemperatureColor = (temp: number) => {
    if (temp >= 30) return 'text-red-600 dark:text-red-400';
    if (temp >= 25) return 'text-orange-600 dark:text-orange-400';
    if (temp >= 20) return 'text-yellow-600 dark:text-yellow-400';
    if (temp >= 15) return 'text-green-600 dark:text-green-400';
    if (temp >= 10) return 'text-blue-600 dark:text-blue-400';
    return 'text-indigo-600 dark:text-indigo-400';
  };

  // 获取天气建议
  const getWeatherAdvice = () => {
    const conditionLower = condition.toLowerCase();
    
    if (conditionLower.includes('雨')) {
      return "记得带伞哦！";
    } else if (conditionLower.includes('雪')) {
      return "注意保暖，小心路滑！";
    } else if (conditionLower.includes('晴') && temperature > 30) {
      return "天气炎热，注意防晒！";
    } else if (conditionLower.includes('晴') && temperature < 10) {
      return "阳光虽好，注意保暖！";
    } else if (conditionLower.includes('风')) {
      return "风有点大，注意防风！";
    } else if (conditionLower.includes('云')) {
      return "多云天气，适合外出！";
    } else {
      return "天气不错，享受美好的一天！";
    }
  };

  // 处理刷新
  const handleRefresh = () => {
    if (isRefreshing || !onRefresh) return;
    
    setIsRefreshing(true);
    onRefresh();
    
    // 模拟刷新延迟
    setTimeout(() => {
      setIsRefreshing(false);
    }, 1000);
  };

  // 格式化温度显示
  const formatTemperature = (temp: number) => {
    return `${Math.round(temp)}°C`;
  };

  // 获取湿度等级
  const getHumidityLevel = (humidity: number) => {
    if (humidity < 30) return '干燥';
    if (humidity < 60) return '舒适';
    if (humidity < 80) return '潮湿';
    return '非常潮湿';
  };

  return (
    <div className="bg-gradient-to-br from-blue-50 to-cyan-50 dark:from-blue-900/30 dark:to-cyan-900/30 rounded-2xl shadow-xl p-6 border border-blue-200 dark:border-blue-800">
      {/* 卡片头部 */}
      <div className="flex items-start justify-between mb-6">
        <div>
          <div className="flex items-center mb-2">
            <MapPin className="w-5 h-5 text-blue-600 dark:text-blue-400 mr-2" />
            <h3 className="text-xl font-bold text-stone-900 dark:text-stone-100">
              {location}
            </h3>
          </div>
          <div className="flex items-center text-sm text-stone-600 dark:text-stone-400">
            <Calendar className="w-4 h-4 mr-1" />
            <span>更新于 {timeSinceUpdate}</span>
          </div>
        </div>
        
        <button
          onClick={handleRefresh}
          disabled={isRefreshing}
          className={`p-2 rounded-full transition-all ${
            isRefreshing 
              ? 'bg-blue-100 dark:bg-blue-900/50' 
              : 'bg-white dark:bg-stone-800 hover:bg-blue-50 dark:hover:bg-blue-900/30'
          }`}
          title="刷新天气"
        >
          <RefreshCw className={`w-5 h-5 text-blue-600 dark:text-blue-400 ${
            isRefreshing ? 'animate-spin' : ''
          }`} />
        </button>
      </div>

      {/* 主要天气信息 */}
      <div className="flex items-center justify-between mb-6">
        <div>
          <div className="flex items-baseline">
            <span className={`text-5xl font-bold ${getTemperatureColor(temperature)}`}>
              {formatTemperature(temperature)}
            </span>
            <span className="ml-2 text-lg text-stone-600 dark:text-stone-400">
              {condition}
            </span>
          </div>
          <p className="mt-2 text-stone-700 dark:text-stone-300">
            {getWeatherAdvice()}
          </p>
        </div>
        
        <div className="text-center">
          {getWeatherIcon()}
          <div className="mt-2 text-sm text-stone-600 dark:text-stone-400">
            实时天气
          </div>
        </div>
      </div>

      {/* 详细数据 */}
      <div className="grid grid-cols-2 gap-4 mb-6">
        <div className="bg-white/50 dark:bg-stone-800/50 rounded-xl p-4">
          <div className="flex items-center mb-2">
            <Droplets className="w-5 h-5 text-blue-500 mr-2" />
            <span className="text-sm font-medium text-stone-700 dark:text-stone-300">湿度</span>
          </div>
          <div className="flex items-baseline">
            <span className="text-2xl font-bold text-blue-600 dark:text-blue-400">
              {humidity}%
            </span>
            <span className="ml-2 text-sm text-stone-600 dark:text-stone-400">
              {getHumidityLevel(humidity)}
            </span>
          </div>
          
          {/* 湿度进度条 */}
          <div className="mt-2">
            <div className="h-2 bg-blue-100 dark:bg-blue-900/30 rounded-full overflow-hidden">
              <div 
                className="h-full bg-gradient-to-r from-blue-400 to-cyan-400 rounded-full"
                style={{ width: `${humidity}%` }}
              />
            </div>
          </div>
        </div>

        <div className="bg-white/50 dark:bg-stone-800/50 rounded-xl p-4">
          <div className="flex items-center mb-2">
            <Thermometer className="w-5 h-5 text-orange-500 mr-2" />
            <span className="text-sm font-medium text-stone-700 dark:text-stone-300">体感温度</span>
          </div>
          <div className="flex items-baseline">
            <span className={`text-2xl font-bold ${getTemperatureColor(temperature)}`}>
              {formatTemperature(temperature)}
            </span>
            <span className="ml-2 text-sm text-stone-600 dark:text-stone-400">
              与气温相近
            </span>
          </div>
          
          {/* 温度指示器 */}
          <div className="mt-2">
            <div className="h-2 bg-gradient-to-r from-blue-400 via-green-400 to-red-400 rounded-full" />
            <div className="flex justify-between text-xs text-stone-500 dark:text-stone-400 mt-1">
              <span>寒冷</span>
              <span>舒适</span>
              <span>炎热</span>
            </div>
          </div>
        </div>
      </div>

      {/* 天气预报 */}
      <div className="bg-white/70 dark:bg-stone-800/70 rounded-xl p-4">
        <div className="flex items-center mb-3">
          <Calendar className="w-5 h-5 text-blue-600 dark:text-blue-400 mr-2" />
          <h4 className="font-medium text-stone-900 dark:text-stone-100">天气预报</h4>
        </div>
        
        <p className="text-stone-700 dark:text-stone-300">
          {forecast}
        </p>
        
        {/* 未来几小时预测（模拟） */}
        <div className="mt-4 grid grid-cols-4 gap-2">
          {['现在', '3h', '6h', '9h'].map((time, index) => {
            // 模拟温度变化
            const tempChange = Math.sin(index * 0.5) * 2;
            const futureTemp = temperature + tempChange;
            const futureIcon = icon; // 简化处理，实际应该有不同的图标
            
            return (
              <div key={time} className="text-center">
                <div className="text-sm font-medium text-stone-700 dark:text-stone-300 mb-1">
                  {time}
                </div>
                <div className="mb-1 flex justify-center">
                  {getWeatherIcon()}
                </div>
                <div className={`text-sm font-bold ${getTemperatureColor(futureTemp)}`}>
                  {formatTemperature(futureTemp)}
                </div>
              </div>
            );
          })}
        </div>
      </div>

      {/* 数据来源和额外信息 */}
      <div className="mt-4 pt-4 border-t border-blue-200 dark:border-blue-800">
        <div className="flex justify-between items-center text-xs text-stone-500 dark:text-stone-400">
          <div>
            数据来源: OpenClaw天气收集
          </div>
          <div className="flex items-center">
            <span className="flex items-center">
              <div className="w-2 h-2 bg-green-500 rounded-full mr-1" />
              实时更新
            </span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default WeatherCard;