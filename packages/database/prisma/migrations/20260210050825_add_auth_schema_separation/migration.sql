/*
  Warnings:

  - You are about to drop the `credential` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `user_credential` table. If the table is not empty, all the data it contains will be lost.

*/
-- CreateSchema
CREATE SCHEMA IF NOT EXISTS "auth";

-- DropForeignKey
ALTER TABLE "user_credential" DROP CONSTRAINT "user_credential_credential_id_fkey";

-- DropForeignKey
ALTER TABLE "user_credential" DROP CONSTRAINT "user_credential_user_id_fkey";

-- DropTable
DROP TABLE "credential";

-- DropTable
DROP TABLE "user_credential";

-- CreateTable
CREATE TABLE "auth"."credential" (
    "id" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "email_verified" BOOLEAN NOT NULL DEFAULT false,
    "password_hash" TEXT NOT NULL,

    CONSTRAINT "credential_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "auth"."user_credential" (
    "id" TEXT NOT NULL,
    "user_id" TEXT NOT NULL,
    "credential_id" TEXT NOT NULL,

    CONSTRAINT "user_credential_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "credential_email_key" ON "auth"."credential"("email");

-- CreateIndex
CREATE UNIQUE INDEX "user_credential_user_id_key" ON "auth"."user_credential"("user_id");

-- CreateIndex
CREATE UNIQUE INDEX "user_credential_credential_id_key" ON "auth"."user_credential"("credential_id");

-- AddForeignKey
ALTER TABLE "auth"."user_credential" ADD CONSTRAINT "user_credential_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "auth"."user_credential" ADD CONSTRAINT "user_credential_credential_id_fkey" FOREIGN KEY ("credential_id") REFERENCES "auth"."credential"("id") ON DELETE CASCADE ON UPDATE CASCADE;
