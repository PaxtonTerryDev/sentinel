import dotenv from 'dotenv';

dotenv.config();

class ConfigError extends Error {
  override message: string;

  constructor(message: string) {
    super(message)
    this.message = message;
  }
}

function env(variable: string, fallback?: string): string {
  if (process.env[variable]) {
    return process.env[variable]!;
  }
  if (fallback) {
    return fallback;
  }
  throw new ConfigError(`Missing environment variable: ${variable}`);
}

export const config = {
  port: parseInt(env("SERVER_PORT"))
}

export type Config = typeof config;
