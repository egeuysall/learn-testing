"use client";

import { useQuery } from "@tanstack/react-query";

type Blog = {
  id: number;
  title: string;
  content: string;
  slug: string;
  tags: string[];
  created_at: string;
  created_by: string;
  cover_link: string;
};

type BlogResponse = {
  data: Blog[];
};

export default function BlogList() {
  const {
    data: response,
    isLoading,
    error,
  } = useQuery({
    queryKey: ["blogs"],
    queryFn: async () => {
      const res = await fetch("https://blogapi.egeuysal.com:8443/v1/blogs");
      if (!res.ok) throw new Error("Failed to fetch blogs");
      return res.json() as Promise<BlogResponse>;
    },
  });

  if (isLoading) return <div>Loading blogs...</div>;
  if (error) return <div>Error loading blogs</div>;

  const blogs = response?.data;

  if (!blogs || blogs.length === 0) {
    return (
      <div>
        <h1>My Blogs</h1>
        <p>No blogs found</p>
      </div>
    );
  }

  return (
    <div>
      <h1>My Blogs</h1>
      <ul>
        {blogs?.map((blog) => (
          <li key={blog.id}>
            <h2>{blog.title}</h2>
            <p>{blog.content.substring(0, 100)}...</p>
            <span>By {blog.created_by}</span>
            <div>Tags: {blog.tags.join(", ")}</div>
          </li>
        ))}
      </ul>
    </div>
  );
}
