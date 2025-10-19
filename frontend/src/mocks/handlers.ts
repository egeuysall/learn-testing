import { http, HttpResponse } from "msw";

export const handlers = [
  http.get("/api/users", () => {
    return HttpResponse.json([
      { id: 1, name: "John Doe", email: "john@example.com" },
      { id: 2, name: "Jane Smith", email: "jane@example.com" },
    ]);
  }),
  http.get("https://blogapi.egeuysal.com:8443/v1/blogs", () => {
    return HttpResponse.json({
      data: [
        {
          id: 1,
          title: "My First Blog",
          content: "This is the content of my first blog post.",
          slug: "my-first-blog",
          tags: ["testing", "development"],
          created_at: "2025-01-15T10:00:00Z",
          created_by: "Ege",
          cover_link: "https://example.com/cover1.jpg",
        },
        {
          id: 2,
          title: "Learning React",
          content: "React is awesome!",
          slug: "learning-react",
          tags: ["react", "frontend"],
          created_at: "2025-01-16T14:30:00Z",
          created_by: "Ege",
          cover_link: "https://example.com/cover2.jpg",
        },
      ],
    });
  }),
];
