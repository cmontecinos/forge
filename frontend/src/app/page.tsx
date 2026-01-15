"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";
import { checkHealth } from "@/lib/api";

export default function Home() {
  const { user, isAuthenticated, isLoading, logout } = useAuth();
  const [healthStatus, setHealthStatus] = useState("checking...");

  useEffect(() => {
    checkHealth()
      .then((health) => setHealthStatus(health.status))
      .catch(() => setHealthStatus("Backend not reachable"));
  }, []);

  const handleLogout = async () => {
    await logout();
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-4">
          Welcome to {"{{.ProjectName}}"}
        </h1>
        <p className="text-gray-600 mb-8">
          Your Next.js + Go + Supabase application
        </p>

        {/* Auth Status */}
        <div className="bg-white border rounded-lg p-4 mb-8 min-w-[300px]">
          {isLoading ? (
            <p className="text-gray-500">Loading...</p>
          ) : isAuthenticated && user ? (
            <div className="space-y-3">
              <p className="text-sm text-gray-500">Logged in as</p>
              <p className="font-medium">{user.email}</p>
              <button
                onClick={handleLogout}
                className="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700 transition-colors text-sm"
              >
                Logout
              </button>
            </div>
          ) : (
            <div className="space-y-3">
              <p className="text-sm text-gray-500">Not logged in</p>
              <div className="flex gap-2 justify-center">
                <Link
                  href="/login"
                  className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors text-sm"
                >
                  Sign In
                </Link>
                <Link
                  href="/register"
                  className="bg-gray-600 text-white px-4 py-2 rounded-md hover:bg-gray-700 transition-colors text-sm"
                >
                  Sign Up
                </Link>
              </div>
            </div>
          )}
        </div>

        {/* Health Check */}
        <div className="bg-gray-100 rounded-lg p-4 mb-8">
          <p className="text-sm text-gray-500">Backend Health</p>
          <p
            className={`font-mono ${
              healthStatus === "healthy" ? "text-green-600" : "text-red-600"
            }`}
          >
            {healthStatus}
          </p>
        </div>

        <div className="space-y-2 text-sm text-gray-500">
          <p>
            Edit{" "}
            <code className="bg-gray-100 px-1 rounded">src/app/page.tsx</code>{" "}
            to get started
          </p>
          <p>
            Backend API:{" "}
            <code className="bg-gray-100 px-1 rounded">
              {process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080"}
            </code>
          </p>
        </div>
      </div>
    </main>
  );
}
