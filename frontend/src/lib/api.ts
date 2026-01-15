const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

interface HealthResponse {
  status: string;
}

export async function checkHealth(): Promise<HealthResponse> {
  const response = await fetch(`${API_URL}/health`, {
    cache: "no-store",
  });

  if (!response.ok) {
    throw new Error(`Health check failed: ${response.status}`);
  }

  return response.json();
}

// Generic API client for future use
export async function apiClient<T>(
  endpoint: string,
  options?: RequestInit
): Promise<T> {
  const response = await fetch(`${API_URL}${endpoint}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...options?.headers,
    },
  });

  if (!response.ok) {
    throw new Error(`API error: ${response.status}`);
  }

  return response.json();
}
