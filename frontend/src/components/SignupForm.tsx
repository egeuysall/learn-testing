"use client";

import { useState } from "react";
import { useMutation } from "@tanstack/react-query";

interface SignupData {
  email: string;
  name: string;
}

export default function SignupForm() {
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [error, setError] = useState("");

  const mutation = useMutation({
    mutationFn: async (data: SignupData) => {
      const res = await fetch("/api/signup", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });

      if (!res.ok) {
        throw new Error("Signup failed");
      }

      return res.json();
    },
    onSuccess: () => {
      setError("");
      setEmail("");
      setName("");
    },
    onError: () => {
      setError("Failed to sign up. Please try again.");
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    // Validation
    if (!email || !name) {
      setError("All fields are required");
      return;
    }

    if (!email.includes("@")) {
      setError("Invalid email format");
      return;
    }

    setError("");
    mutation.mutate({ email, name });
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} aria-label="Email" />
      </div>

      <div>
        <input type="text" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} aria-label="Name" />
      </div>

      {error && <div role="alert">{error}</div>}

      {mutation.isPending && <div>Submitting...</div>}
      {mutation.isSuccess && <div>Signup successful!</div>}

      <button type="submit" disabled={mutation.isPending}>
        Sign Up
      </button>
    </form>
  );
}
