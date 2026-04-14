import { useEffect, useState } from "react";
import { useSessionStore } from "../../../store/agent-chat/sessionStore";
import { useMessageStore } from "../../../store/agent-chat/messageStore";
import {
  getFileTree,
  getFileContent,
  type FileNode,
  type FileContentResponse,
} from "../../../api/agent-chat/workspace";
import FileTree from "./FileTree";
import FileViewer from "./FileViewer";

export default function WorkspacePanel() {
  const username = useSessionStore((s) => s.username);
  const messages = useMessageStore((s) => s.messages);

  const [tree, setTree] = useState<FileNode[]>([]);
  const [selectedPath, setSelectedPath] = useState<string | null>(null);
  const [fileContent, setFileContent] = useState<FileContentResponse | null>(null);
  const [treeLoading, setTreeLoading] = useState(false);
  const [fileLoading, setFileLoading] = useState(false);

  async function refreshTree() {
    if (!username) return;
    setTreeLoading(true);
    try {
      const res = await getFileTree(username);
      setTree(res.tree);
    } catch {
      // silently ignore
    } finally {
      setTreeLoading(false);
    }
  }

  async function selectFile(node: FileNode) {
    setSelectedPath(node.path);
    setFileLoading(true);
    try {
      const res = await getFileContent(node.path, username);
      setFileContent(res);
    } catch {
      setFileContent(null);
    } finally {
      setFileLoading(false);
    }
  }

  useEffect(() => {
    refreshTree();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [username]);

  // Refresh file tree after each agent reply
  const agentMsgCount = messages.filter((m) => m.role === "agent").length;
  useEffect(() => {
    if (agentMsgCount > 0) refreshTree();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [agentMsgCount]);

  return (
    <div className="ac-panel ac-workspace">
      <div className="ac-panel__header">
        <span className="ac-panel__label">WORKSPACE</span>
        <button className="ac-panel__action" onClick={refreshTree} disabled={treeLoading}>
          {treeLoading ? "..." : "刷新"}
        </button>
      </div>
      <div className="ac-workspace__body">
        <div className="ac-workspace__tree">
          {tree.length === 0 && !treeLoading && (
            <p className="ac-workspace__empty">无文件</p>
          )}
          <FileTree nodes={tree} selectedPath={selectedPath} onSelect={selectFile} />
        </div>
        <div className="ac-workspace__viewer">
          <FileViewer file={fileContent} loading={fileLoading} />
        </div>
      </div>
    </div>
  );
}
