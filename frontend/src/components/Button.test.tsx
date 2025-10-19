import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import Button from "./Button";

describe("Button", () => {
  it("shows the text inside", () => {
    render(<Button onClick={() => {}}>Click me</Button>);

    expect(screen.getByText("Click me")).toBeInTheDocument();
  });

  it("shows submit inside", () => {
    render(<Button onClick={() => {}}>Submit</Button>);

    expect(screen.getByText("Submit")).toBeInTheDocument();
  });

  it("is disabled when disabled prop is true", () => {
    render(
      <Button onClick={() => {}} disabled>
        Submit
      </Button>,
    );

    expect(screen.getByText("Submit")).toBeDisabled();
  });
});
