import { render, screen, fireEvent } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { describe, it, expect } from "vitest";
import SignupForm from "./SignupForm";
import userEvent from "@testing-library/user-event";

// Helper to wrap component with QueryClient
function renderWithClient(component: React.ReactElement) {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: { retry: false },
      mutations: { retry: false },
    },
  });

  return render(<QueryClientProvider client={queryClient}>{component}</QueryClientProvider>);
}

describe("SignupForm", () => {
  it("renders email and name inputs", () => {
    renderWithClient(<SignupForm />);

    expect(screen.getByLabelText("Email")).toBeInTheDocument();
    expect(screen.getByLabelText("Name")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Sign Up" })).toBeInTheDocument();
  });

  it("shows error when submitting empty form", async () => {
    renderWithClient(<SignupForm />);
    const submitButton = screen.getByRole("button", { name: "Sign Up" });
    await userEvent.click(submitButton);

    expect(screen.getByText("All fields are required")).toBeInTheDocument();
  });

  // it("shows error when submitting invalid email", () => {
  //   renderWithClient(<SignupForm />);

  //   const emailInput = screen.getByLabelText("Email");
  //   const nameInput = screen.getByLabelText("Name");

  //   // Use fireEvent.change instead of userEvent.type
  //   fireEvent.change(emailInput, { target: { value: "notanemail" } });
  //   fireEvent.change(nameInput, { target: { value: "John Doe" } });

  //   const submitButton = screen.getByRole("button", { name: "Sign Up" });
  //   fireEvent.click(submitButton);

  //   expect(screen.getByText("Invalid email format")).toBeInTheDocument();
  // });

  // TODO: Test 3 - Shows error for invalid email
  // TODO: Test 4 - Shows "Submitting..." while loading
  // TODO: Test 5 - Shows "Signup successful!" on success
  // TODO: Test 6 - Shows error when API fails
});
