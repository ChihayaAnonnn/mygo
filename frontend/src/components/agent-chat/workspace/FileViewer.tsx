import ReactMarkdown from "react-markdown";
import type { FileContentResponse } from "../../../api/agent-chat/workspace";

interface Props {
  file: FileContentResponse | null;
  loading: boolean;
}

export default function FileViewer({ file, loading }: Props) {
  if (loading) {
    return <div className="ac-file-viewer ac-file-viewer--loading">加载中...</div>;
  }

  if (!file) {
    return (
      <div className="ac-file-viewer ac-file-viewer--empty">
        <p>选择左侧文件预览内容</p>
      </div>
    );
  }

  const isMarkdown = file.suffix === ".md";

  return (
    <div className="ac-file-viewer">
      <div className="ac-file-viewer__path">{file.path}</div>
      <div className="ac-file-viewer__content">
        {isMarkdown ? (
          <div className="ac-bubble__markdown"><ReactMarkdown>{file.content}</ReactMarkdown></div>
        ) : (
          <pre className="ac-file-viewer__pre">{file.content}</pre>
        )}
      </div>
    </div>
  );
}
