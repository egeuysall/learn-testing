import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, it, expect, vi } from "vitest";
import Button from "./Button";

describe("Button", () => {
  it("renders with text", () => {
    render(<Button onClick={() => {}}>Click me</Button>);

    expect(screen.getByText("Click me")).toBeInTheDocument();
  });

  it("calls onClick when clicked", async () => {
    // Create a mock function to track if it was called
    const mockOnClick = vi.fn();

    // Render button with the mock
    render(<Button onClick={mockOnClick}>Click me</Button>);

    // Get the button
    const button = screen.getByText("Click me");

    // Simulate user clicking it
    await userEvent.click(button);

    // Assert the mock was called
    expect(mockOnClick).toHaveBeenCalled();
  });

  // TODO 2: Write test that button is disabled when disabled prop is true
  // Hint: render with disabled={true}, then use expect(button).toBeDisabled()
});
