import { randomBytes, scrypt, timingSafeEqual } from "node:crypto";
import { promisify } from "node:util";

const scryptAsync = promisify(scrypt);

const SALT_LENGTH = 32;
const KEY_LENGTH = 64;

export class Password {
  static async hash(password: string): Promise<string> {
    const salt = randomBytes(SALT_LENGTH).toString("hex");
    const derivedKey = (await scryptAsync(password, salt, KEY_LENGTH)) as Buffer;
    return `${salt}:${derivedKey.toString("hex")}`;
  }

  static async verify(password: string, stored: string): Promise<boolean> {
    const [salt, storedHash] = stored.split(":");
    if (!salt || !storedHash) {
      return false;
    }

    const derivedKey = (await scryptAsync(password, salt, KEY_LENGTH)) as Buffer;
    const storedBuffer = Buffer.from(storedHash, "hex");

    return timingSafeEqual(derivedKey, storedBuffer);
  }
}
