import type { FileNode } from "../../../api/agent-chat/workspace";

interface Props {
  nodes: FileNode[];
  selectedPath: string | null;
  onSelect: (node: FileNode) => void;
  depth?: number;
}

export default function FileTree({ nodes, selectedPath, onSelect, depth = 0 }: Props) {
  return (
    <ul className="ac-file-tree" style={{ paddingLeft: depth > 0 ? "1rem" : 0 }}>
      {nodes.map((node) => (
        <li key={node.path} className="ac-file-tree__item">
          {node.type === "dir" ? (
            <>
              <span className="ac-file-tree__dir">
                <span className="ac-file-tree__icon">▸</span>
                {node.name}
              </span>
              {node.children && node.children.length > 0 && (
                <FileTree
                  nodes={node.children}
                  selectedPath={selectedPath}
                  onSelect={onSelect}
                  depth={depth + 1}
                />
              )}
            </>
          ) : (
            <button
              className={`ac-file-tree__file ${selectedPath === node.path ? "ac-file-tree__file--active" : ""} ${!node.readable ? "ac-file-tree__file--disabled" : ""}`}
              onClick={() => node.readable && onSelect(node)}
              disabled={!node.readable}
            >
              <span className="ac-file-tree__icon">◦</span>
              {node.name}
            </button>
          )}
        </li>
      ))}
    </ul>
  );
}
