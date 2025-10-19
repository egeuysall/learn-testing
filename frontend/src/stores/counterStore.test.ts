import { describe, it, expect, beforeEach } from "vitest";
import { useCounterStore } from "./counterStore";

describe("Counter Store", () => {
  beforeEach(() => {
    // Reset store before each test
    useCounterStore.getState().reset();
  });

  it("starts at 0", () => {
    const { count } = useCounterStore.getState();
    expect(count).toBe(0);
  });

  it("can increment count", () => {
    useCounterStore.getState().increment();
    const { count } = useCounterStore.getState();
    expect(count).toBe(1);
  });

  it("can decrement count", () => {
    useCounterStore.getState().decrement();
    const { count } = useCounterStore.getState();
    expect(count).toBe(-1);
  });
});
