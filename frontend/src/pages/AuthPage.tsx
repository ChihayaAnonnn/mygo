import { useState, FormEvent, ChangeEvent } from "react";

// API 响应类型
interface ApiResponse<T> {
  code: number;
  message: string;
  data?: T;
}

interface LoginResponse {
  session_id: string;
  user_id: number;
  username: string;
}

interface RegisterResponse {
  user_id: number;
  username: string;
  email: string;
}

export default function AuthPage() {
  const [isLogin, setIsLogin] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  // 表单数据
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  // 表单验证错误
  const [fieldErrors, setFieldErrors] = useState({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
    setFieldErrors((prev) => ({ ...prev, [name]: "" }));
    setError("");
  };

  const validateForm = (): boolean => {
    const errors = {
      username: "",
      email: "",
      password: "",
      confirmPassword: "",
    };
    let isValid = true;

    if (!formData.username.trim()) {
      errors.username = "请输入用户名";
      isValid = false;
    } else if (formData.username.length < 3) {
      errors.username = "用户名至少 3 个字符";
      isValid = false;
    }

    if (!isLogin) {
      if (!formData.email.trim()) {
        errors.email = "请输入邮箱";
        isValid = false;
      } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
        errors.email = "请输入有效的邮箱地址";
        isValid = false;
      }
    }

    if (!formData.password) {
      errors.password = "请输入密码";
      isValid = false;
    } else if (formData.password.length < 6) {
      errors.password = "密码至少 6 个字符";
      isValid = false;
    }

    if (!isLogin) {
      if (!formData.confirmPassword) {
        errors.confirmPassword = "请确认密码";
        isValid = false;
      } else if (formData.password !== formData.confirmPassword) {
        errors.confirmPassword = "两次输入的密码不一致";
        isValid = false;
      }
    }

    setFieldErrors(errors);
    return isValid;
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    if (!validateForm()) return;

    setIsLoading(true);

    try {
      const endpoint = isLogin ? "/api/users/login" : "/api/users/register";
      const body = isLogin
        ? { username: formData.username, password: formData.password }
        : {
            username: formData.username,
            email: formData.email,
            password: formData.password,
          };

      const response = await fetch(endpoint, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
      });

      const result: ApiResponse<LoginResponse | RegisterResponse> =
        await response.json();

      if (result.code === 0) {
        if (isLogin) {
          const loginData = result.data as LoginResponse;
          // 存储 session
          localStorage.setItem("sessionId", loginData.session_id);
          localStorage.setItem("userId", String(loginData.user_id));
          localStorage.setItem("username", loginData.username);
          setSuccess(`欢迎回来，${loginData.username}！`);
          // 可以在这里重定向到主页
          setTimeout(() => {
            window.location.href = "/";
          }, 1500);
        } else {
          setSuccess("注册成功！请登录您的账号。");
          setTimeout(() => {
            setIsLogin(true);
            setFormData((prev) => ({
              ...prev,
              email: "",
              password: "",
              confirmPassword: "",
            }));
          }, 1500);
        }
      } else {
        setError(result.message || "操作失败，请重试");
      }
    } catch {
      setError("网络错误，请检查连接后重试");
    } finally {
      setIsLoading(false);
    }
  };

  const toggleMode = () => {
    setIsLogin(!isLogin);
    setError("");
    setSuccess("");
    setFieldErrors({
      username: "",
      email: "",
      password: "",
      confirmPassword: "",
    });
  };

  return (
    <main className="page">
      {/* 背景发光效果 */}
      <div className="bg" aria-hidden="true">
        <div className="glow glow-1" />
        <div className="glow glow-2" />
        <div className="glow glow-3" />
      </div>

      {/* 浮动装饰粒子 */}
      <div className="particles" aria-hidden="true">
        {[...Array(6)].map((_, i) => (
          <div key={i} className={`particle particle-${i + 1}`} />
        ))}
      </div>

      <section className="auth-card">
        {/* 卡片顶部装饰线 */}
        <div className="card-accent-line" aria-hidden="true" />

        <header className="auth-header">
          <div className="logo-mark">
            <svg
              viewBox="0 0 40 40"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <circle
                cx="20"
                cy="20"
                r="18"
                stroke="currentColor"
                strokeWidth="2"
              />
              <path
                d="M12 20c0-4.4 3.6-8 8-8s8 3.6 8 8-3.6 8-8 8"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
              />
              <circle cx="20" cy="20" r="3" fill="currentColor" />
            </svg>
          </div>
          <p className="brand-name">mygo</p>
        </header>

        <div className="auth-title-wrapper">
          <h1 className="auth-title">{isLogin ? "登录" : "创建账号"}</h1>
          <p className="auth-subtitle">
            {isLogin ? "欢迎回来，请登录您的账号" : "加入我们，开启新旅程"}
          </p>
        </div>

        {/* 模式切换标签 */}
        <div className="auth-tabs">
          <button
            type="button"
            className={`auth-tab ${isLogin ? "active" : ""}`}
            onClick={() => !isLogin && toggleMode()}
          >
            登录
          </button>
          <button
            type="button"
            className={`auth-tab ${!isLogin ? "active" : ""}`}
            onClick={() => isLogin && toggleMode()}
          >
            注册
          </button>
          <div
            className="tab-indicator"
            style={{ transform: isLogin ? "translateX(0)" : "translateX(100%)" }}
          />
        </div>

        {/* 消息提示 */}
        {error && (
          <div className="auth-message error">
            <svg viewBox="0 0 20 20" fill="currentColor">
              <path
                fillRule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                clipRule="evenodd"
              />
            </svg>
            {error}
          </div>
        )}
        {success && (
          <div className="auth-message success">
            <svg viewBox="0 0 20 20" fill="currentColor">
              <path
                fillRule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                clipRule="evenodd"
              />
            </svg>
            {success}
          </div>
        )}

        {/* 表单 */}
        <form className="auth-form" onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="username" className="form-label">
              用户名
            </label>
            <div className="input-wrapper">
              <svg
                className="input-icon"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fillRule="evenodd"
                  d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
                  clipRule="evenodd"
                />
              </svg>
              <input
                type="text"
                id="username"
                name="username"
                className={`form-input ${fieldErrors.username ? "error" : ""}`}
                placeholder="请输入用户名"
                value={formData.username}
                onChange={handleInputChange}
                autoComplete="username"
              />
            </div>
            {fieldErrors.username && (
              <span className="field-error">{fieldErrors.username}</span>
            )}
          </div>

          {!isLogin && (
            <div className="form-group">
              <label htmlFor="email" className="form-label">
                邮箱
              </label>
              <div className="input-wrapper">
                <svg
                  className="input-icon"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                  <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                </svg>
                <input
                  type="email"
                  id="email"
                  name="email"
                  className={`form-input ${fieldErrors.email ? "error" : ""}`}
                  placeholder="请输入邮箱地址"
                  value={formData.email}
                  onChange={handleInputChange}
                  autoComplete="email"
                />
              </div>
              {fieldErrors.email && (
                <span className="field-error">{fieldErrors.email}</span>
              )}
            </div>
          )}

          <div className="form-group">
            <label htmlFor="password" className="form-label">
              密码
            </label>
            <div className="input-wrapper">
              <svg
                className="input-icon"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fillRule="evenodd"
                  d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
                  clipRule="evenodd"
                />
              </svg>
              <input
                type="password"
                id="password"
                name="password"
                className={`form-input ${fieldErrors.password ? "error" : ""}`}
                placeholder="请输入密码"
                value={formData.password}
                onChange={handleInputChange}
                autoComplete={isLogin ? "current-password" : "new-password"}
              />
            </div>
            {fieldErrors.password && (
              <span className="field-error">{fieldErrors.password}</span>
            )}
          </div>

          {!isLogin && (
            <div className="form-group">
              <label htmlFor="confirmPassword" className="form-label">
                确认密码
              </label>
              <div className="input-wrapper">
                <svg
                  className="input-icon"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fillRule="evenodd"
                    d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                    clipRule="evenodd"
                  />
                </svg>
                <input
                  type="password"
                  id="confirmPassword"
                  name="confirmPassword"
                  className={`form-input ${fieldErrors.confirmPassword ? "error" : ""}`}
                  placeholder="请再次输入密码"
                  value={formData.confirmPassword}
                  onChange={handleInputChange}
                  autoComplete="new-password"
                />
              </div>
              {fieldErrors.confirmPassword && (
                <span className="field-error">{fieldErrors.confirmPassword}</span>
              )}
            </div>
          )}

          {isLogin && (
            <div className="form-options">
              <label className="remember-me">
                <input type="checkbox" />
                <span className="checkbox-mark" />
                <span>记住我</span>
              </label>
              <a href="#" className="forgot-link" onClick={(e) => e.preventDefault()}>
                忘记密码？
              </a>
            </div>
          )}

          <button type="submit" className="submit-btn" disabled={isLoading}>
            {isLoading ? (
              <span className="loading-spinner" />
            ) : isLogin ? (
              "登录"
            ) : (
              "创建账号"
            )}
          </button>
        </form>

        <footer className="auth-footer">
          <span className="footer-text">
            {isLogin ? "还没有账号？" : "已有账号？"}
          </span>
          <button type="button" className="switch-mode-btn" onClick={toggleMode}>
            {isLogin ? "立即注册" : "去登录"}
          </button>
        </footer>
      </section>
    </main>
  );
}
