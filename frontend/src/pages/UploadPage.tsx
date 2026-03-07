import { useState, useRef } from "react";
import { Link } from "react-router-dom";
import { API_CONFIG, formatFileSize } from "../config/api";

export default function UploadPage() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploadStatus, setUploadStatus] = useState<"idle" | "uploading" | "success" | "error">(
    "idle"
  );
  const [errorMessage, setErrorMessage] = useState<string>("");
  const [successMessage, setSuccessMessage] = useState<string>("");
  const [isDragging, setIsDragging] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  const validate = (file: File): string | null => {
    if (file.size > API_CONFIG.maxFileSize) {
      return `文件大小超出限制（最大 ${formatFileSize(API_CONFIG.maxFileSize)}）`;
    }
    if (file.type !== "text/markdown" && !file.name.endsWith(".md")) {
      return "请选择 Markdown 文件（.md）";
    }
    return null;
  };

  const setFile = (file: File) => {
    setErrorMessage("");
    setSuccessMessage("");
    const err = validate(file);
    if (err) {
      setErrorMessage(err);
      setSelectedFile(null);
    } else {
      setSelectedFile(file);
    }
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files?.[0]) setFile(e.target.files[0]);
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragging(true);
  };

  const handleDragLeave = () => setIsDragging(false);

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragging(false);
    if (e.dataTransfer.files?.[0]) setFile(e.dataTransfer.files[0]);
  };

  const handleUpload = async () => {
    if (!selectedFile) return;

    setUploadStatus("uploading");
    setErrorMessage("");
    setSuccessMessage("");

    try {
      const formData = new FormData();
      formData.append("file", selectedFile);

      const response = await fetch(`${API_CONFIG.baseUrl}${API_CONFIG.endpoints.uploadFile}`, {
        method: "POST",
        body: formData,
      });

      const result = await response.json();
      if (!response.ok) throw new Error(result.error || "上传失败");

      setSuccessMessage(`"${result.filename}" 上传成功，${formatFileSize(result.size)}`);
      setUploadStatus("success");
      setSelectedFile(null);
      if (inputRef.current) inputRef.current.value = "";

      setTimeout(() => {
        setUploadStatus("idle");
        setSuccessMessage("");
      }, 4000);
    } catch (error) {
      setErrorMessage(error instanceof Error ? error.message : "上传失败，请重试");
      setUploadStatus("error");
    }
  };

  return (
    <div className="upload-page">
      <div className="upload-page-grid">

        {/* 左侧信息面板 */}
        <aside className="upload-info-panel">
          <div>
            <span className="eyebrow">Admin</span>
            <h1 className="upload-page-title">上传文章</h1>
            <p className="upload-page-desc">
              上传 Markdown 文件以发布新内容。文件格式需包含 front matter 元信息。
            </p>
          </div>

          <div className="upload-constraints">
            <p className="upload-constraints-label">限制与说明</p>
            <ul>
              <li>
                <span className="uc-key">格式</span>
                <span className="uc-val">.md</span>
              </li>
              <li>
                <span className="uc-key">大小上限</span>
                <span className="uc-val">{formatFileSize(API_CONFIG.maxFileSize)}</span>
              </li>
              <li>
                <span className="uc-key">保存路径</span>
                <span className="uc-val">/workspace/data/posts/</span>
              </li>
              <li>
                <span className="uc-key">生效方式</span>
                <span className="uc-val">重新构建前端后可见</span>
              </li>
            </ul>
          </div>

          <Link to="/" className="text-link upload-back-link">
            ← 返回主站
          </Link>
        </aside>

        {/* 右侧上传区 */}
        <div className="upload-form-panel">
          <div
            className={`upload-dropzone${isDragging ? " upload-dropzone--active" : ""}${selectedFile ? " upload-dropzone--selected" : ""}`}
            onDragOver={handleDragOver}
            onDragLeave={handleDragLeave}
            onDrop={handleDrop}
            onClick={() => inputRef.current?.click()}
            aria-hidden="false"
          >
            <input
              ref={inputRef}
              id="file-input"
              type="file"
              accept=".md,text/markdown"
              onChange={handleFileChange}
              style={{ display: "none" }}
            />

            {selectedFile ? (
              <div className="dropzone-file-info">
                <div className="dropzone-file-icon" aria-hidden="true">✦</div>
                <p className="dropzone-filename">{selectedFile.name}</p>
                <p className="dropzone-filesize">{formatFileSize(selectedFile.size)}</p>
                <p className="dropzone-hint">点击更换文件</p>
              </div>
            ) : (
              <div className="dropzone-placeholder">
                <div className="dropzone-icon" aria-hidden="true">↑</div>
                <p className="dropzone-main-text">拖放文件至此处</p>
                <p className="dropzone-sub-text">或点击选择 .md 文件</p>
              </div>
            )}
          </div>

          {errorMessage && (
            <p className="upload-feedback upload-feedback--error">{errorMessage}</p>
          )}
          {successMessage && (
            <p className="upload-feedback upload-feedback--success">{successMessage}</p>
          )}

          <button
            onClick={handleUpload}
            disabled={!selectedFile || uploadStatus === "uploading"}
            className="upload-submit-btn"
          >
            {uploadStatus === "uploading" ? (
              <span className="upload-btn-inner">
                <span className="upload-spinner" aria-hidden="true" />
                上传中…
              </span>
            ) : (
              "上传文件"
            )}
          </button>
        </div>

      </div>
    </div>
  );
}
