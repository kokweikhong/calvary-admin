"use client";

export default function Page() {
  async function uploadFile(data: FormData) {
    await fetch("http://localhost:8080/files/upload", {
      method: "POST",
      // no multipart boundary param in Content-Type
      headers: {
        "Content-Type": "multipart/form-data",
      },
      body: data,
    });
  }
  return (
    <div>
      {/* multipart-form */}
      <form encType="multipart/form-data">
        <input
          type="file"
          onChange={async (e) => {
            console.log(e.target.files);
            if (!e.target.files) return;
            const data = new FormData();
            console.log(e.target.files[0]);
            data.append("file", e.target.files[0]);
            data.append("filename", "file");
            data.append("saveDir", "");

            await uploadFile(data);
          }}
        />
      </form>
    </div>
  );
}
