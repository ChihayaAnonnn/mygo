import { agentClient } from "./client";

export interface FileNode {
  name: string;
  path: string;
  type: "file" | "dir";
  readable?: boolean;
  children?: FileNode[];
}

export interface FileTreeResponse {
  user: string;
  tree: FileNode[];
}

export interface FileContentResponse {
  path: string;
  content: string;
  suffix: string;
}

export async function getFileTree(user: string): Promise<FileTreeResponse> {
  const { data } = await agentClient.get<FileTreeResponse>("/workspace/files", {
    params: { user },
  });
  return data;
}

export async function getFileContent(
  filePath: string,
  user: string
): Promise<FileContentResponse> {
  const { data } = await agentClient.get<FileContentResponse>(
    `/workspace/files/${filePath}`,
    { params: { user } }
  );
  return data;
}
