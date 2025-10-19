import { render, screen, waitFor } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { describe, it, expect } from "vitest";
import { http, HttpResponse } from "msw";
import { server } from "../mocks/server";
import BlogList from "./BlogList";

function renderWithClient(component: React.ReactElement) {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
    },
  });

  return render(<QueryClientProvider client={queryClient}>{component}</QueryClientProvider>);
}

describe("BlogList", () => {
  it("shows loading state initially", () => {
    renderWithClient(<BlogList />);

    expect(screen.getByText("Loading blogs...")).toBeInTheDocument();
  });

  it("displays blogs after loading", async () => {
    renderWithClient(<BlogList />);

    // Wait for the blogs to load and appear
    await waitFor(() => {
      expect(screen.getByText("My First Blog")).toBeInTheDocument();
    });

    // Check that both blogs are displayed
    expect(screen.getByText("Learning React")).toBeInTheDocument();
    expect(screen.getAllByText(/By Ege/)).toHaveLength(2);
  });

  // TODO: Test that shows error when API fails
  // Hint: Use server.use() to return error, then check for "Error loading blogs"
  it("shows error when API fails", async () => {
    server.use(
      http.get("https://blogapi.egeuysal.com:8443/v1/blogs", () => {
        return new HttpResponse(null, { status: 500 });
      }),
    );

    renderWithClient(<BlogList />);

    await waitFor(() => {
      expect(screen.getByText("Error loading blogs")).toBeInTheDocument();
    });
  });

  it("shows message when no blogs exist", async () => {
    server.use(
      http.get("https://blogapi.egeuysal.com:8443/v1/blogs", () => {
        return HttpResponse.json({
          data: [],
        });
      }),
    );

    renderWithClient(<BlogList />);

    await waitFor(() => {
      expect(screen.getByText("No blogs found")).toBeInTheDocument();
    });
  });
});
