import matter from "gray-matter";

export interface IContent {
  content: string;
}

export interface IMarkdown<T> {
  content: string;
  data: T;
}

export type IMarkdownT<T> = T & IContent;

export function parseMarkdown<T>(markdown: string): IMarkdown<T> {
  const parsed = matter(markdown);
  return {
    content: parsed.content,
    data: parsed.data as T,
  };
}

export function parseMarkdownT<T>(markdown: string): IMarkdownT<T> {
  const parsed = matter(markdown);
  return {
    content: parsed.content,
    ...(parsed.data as T),
  };
}

export function extractData<T>(markdown: string): T {
  const parsed = matter(markdown);
  return parsed.data as T;
}
