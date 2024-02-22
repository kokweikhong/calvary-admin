import { cn } from "@/lib/utils";
import { uploadFile } from "@/services/file";
import { FC } from "react";
import { toast } from "sonner";

type FileUploadButtonProps = {
  filename: string;
  saveDir?: string;
  className?: string;
  setValue?: (value: string) => void;
};

const FileUploadButton: FC<FileUploadButtonProps> = ({
  filename,
  saveDir = "",
  className = "",
  setValue,
}) => {
  async function handleFileUpload(e: React.ChangeEvent<HTMLInputElement>) {
    if (!e.target.files) return;
    const data = new FormData();
    data.append("file", e.target.files[0]);
    data.append("filename", filename);
    data.append("saveDir", saveDir);
    toast(
      "This will be uploaded to server and replace existing file, are you confirm to upload?",
      {
        action: {
          label: "Confirm",
          onClick: () => {
            toast.promise(uploadFile(data), {
              loading: "Uploading...",
              success(data) {
                console.log(data);
                if (setValue) {
                  setValue(data);
                }
                return "File uploaded successfully!";
              },
              error: "Failed to upload file!",
            });
          },
        },
        cancel: {
          label: "Cancel",
          onClick: () => console.log("Cancel!"),
        },
      }
    );
    // await uploadFile(data);
  }

  return (
    <div>
      <label
        htmlFor="file-upload"
        className={cn(
          "cursor-pointer inline-flex items-center justify-center px-4 py-2 border border-transparent text-base",
          "leading-6 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-700",
          "focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo transition duration-150 ease-in-out",
          "whitespace-nowrap",
          className
        )}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth="1.5"
          stroke="currentColor"
          className="-ml-1 mr-1 h-5 w-5"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M12 16.5V9.75m0 0 3 3m-3-3-3 3M6.75 19.5a4.5 4.5 0 0 1-1.41-8.775 5.25 5.25 0 0 1 10.233-2.33 3 3 0 0 1 3.758 3.848A3.752 3.752 0 0 1 18 19.5H6.75Z"
          />
        </svg>
        <span>Upload File</span>
      </label>
      <input
        id="file-upload"
        type="file"
        className="hidden"
        onChange={async (e) => {
          await handleFileUpload(e);
        }}
      />
    </div>
  );
};

export default FileUploadButton;
