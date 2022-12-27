import matter from "gray-matter";
import { Bv5Obj } from "./classes";

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
