"use server";

import axios from "axios";

const fileApiURL = process.env.BACKEND_API_URL + "/files";

const axiosFile = axios.create({
  baseURL: fileApiURL,
});

export async function uploadFile(data: FormData) {
  const response = await axiosFile.post("/upload", data);
  return response.data;
}
