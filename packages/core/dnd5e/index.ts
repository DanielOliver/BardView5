import matter from "gray-matter";

export interface IMarkdown<T> {
  content: string;
  data: T;
}

export function parseMarkdown<T>(markdown: string): IMarkdown<T> {
  const parsed = matter(markdown);
  return {
    content: parsed.content,
    data: parsed.data as T,
  };
}
