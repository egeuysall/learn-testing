import { render, screen, waitFor } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { describe, it, expect } from "vitest";
import UserList from "./UserList";
import { http, HttpResponse } from "msw";
import { server } from "../mocks/server";

function renderWithClient(component: React.ReactElement) {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
    },
  });

  return render(<QueryClientProvider client={queryClient}>{component}</QueryClientProvider>);
}

describe("UserList", () => {
  it("shows loading state initially", () => {
    renderWithClient(<UserList />);

    expect(screen.getByText("Loading...")).toBeInTheDocument();
  });

  it("displays users after loading", async () => {
    renderWithClient(<UserList />);

    // Wait for users to appear (MSW will return mocked data)
    await waitFor(() => {
      expect(screen.getByText("John Doe")).toBeInTheDocument();
    });

    expect(screen.getByText("Jane Smith")).toBeInTheDocument();
  });

  it("shows error when API fails", async () => {
    // Override the handler to return error for this test only
    server.use(
      http.get("/api/users", () => {
        return new HttpResponse(null, { status: 500 });
      }),
    );

    renderWithClient(<UserList />);

    // Wait for error message to appear
    await waitFor(() => {
      expect(screen.getByText("Error loading users")).toBeInTheDocument();
    });
  });
});
