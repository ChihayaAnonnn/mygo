import type { Post } from "../types/post";

const postModules = import.meta.glob('../../posts/*.md', { as: 'raw', eager: true });

interface FrontMatter {
  title: string;
  summary: string;
  date: string;
  tags: string[];
}

function parseMarkdownFile(raw: string): { meta: FrontMatter; content: string } {
  const matched = raw.match(/^---\n([\s\S]*?)\n---\n?([\s\S]*)$/);

  if (!matched) {
    return {
      meta: { title: "", summary: "", date: "", tags: [] },
      content: raw.trim(),
    };
  }

  const [, frontMatterBlock, contentBlock] = matched;
  const lineMap = frontMatterBlock.split("\n").reduce<Record<string, string>>((acc, line) => {
    const [key, ...rest] = line.split(":");
    if (!key || rest.length === 0) {
      return acc;
    }
    acc[key.trim()] = rest.join(":").trim();
    return acc;
  }, {});

  return {
    meta: {
      title: lineMap.title ?? "",
      summary: lineMap.summary ?? "",
      date: lineMap.date ?? "",
      tags: (lineMap.tags ?? "")
        .split(",")
        .map((tag) => tag.trim())
        .filter(Boolean),
    },
    content: contentBlock.trim(),
  };
}

export const posts: Post[] = Object.entries(postModules)
  .map(([path, raw]) => {
    const slug = path.split('/').pop()?.replace('.md', '') || '';
    const { meta, content } = parseMarkdownFile(raw as string);
    return {
      slug,
      title: meta.title,
      summary: meta.summary,
      date: meta.date,
      tags: meta.tags,
      content,
    };
  })
  .sort((a, b) => b.date.localeCompare(a.date));

export function getPostBySlug(slug: string): Post | undefined {
  return posts.find((post) => post.slug === slug);
}
