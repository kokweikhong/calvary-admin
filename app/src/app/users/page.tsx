"use client";

import FileUploadButton from "@/components/FileUploadButton";

export default function Page() {
  async function uploadFile(data: FormData) {
    await fetch("http://localhost:8080/files/upload", {
      method: "POST",
      // no multipart boundary param in Content-Type

      body: data,
    });
  }
  return (
    <div>
      <FileUploadButton filename="test" />
    </div>
  );
}
