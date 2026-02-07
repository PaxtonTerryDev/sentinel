/*
  Warnings:

  - You are about to drop the `UsersOnApplications` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "UsersOnApplications" DROP CONSTRAINT "UsersOnApplications_applicationId_fkey";

-- DropForeignKey
ALTER TABLE "UsersOnApplications" DROP CONSTRAINT "UsersOnApplications_userId_fkey";

-- DropTable
DROP TABLE "UsersOnApplications";

-- CreateTable
CREATE TABLE "users_on_applications" (
    "userId" TEXT NOT NULL,
    "applicationId" TEXT NOT NULL,

    CONSTRAINT "users_on_applications_pkey" PRIMARY KEY ("userId","applicationId")
);

-- AddForeignKey
ALTER TABLE "users_on_applications" ADD CONSTRAINT "users_on_applications_userId_fkey" FOREIGN KEY ("userId") REFERENCES "user"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "users_on_applications" ADD CONSTRAINT "users_on_applications_applicationId_fkey" FOREIGN KEY ("applicationId") REFERENCES "application"("id") ON DELETE CASCADE ON UPDATE CASCADE;
