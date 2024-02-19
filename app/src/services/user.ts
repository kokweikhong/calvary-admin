"use server";

import axios from "axios";

const userApiURL = process.env.BACKEND_API_URL;

const axiosUser = axios.create({
  baseURL: `${userApiURL}/users`,
});

export async function getUsers() {
  const response = await axiosUser.get("/");
  return response.data;
}

export async function getUser(id: string) {
  const response = await axiosUser.get(`/${id}`);
  return response.data;
}

export async function createUser(data: any) {
  const response = await axiosUser.post("/", data);
  return response.data;
}

export async function updateUser(id: string, data: any) {
  const response = await axiosUser.put(`/${id}`, data);
  return response.data;
}

export async function deleteUser(id: string) {
  const response = await axiosUser.delete(`/${id}`);
  return response.data;
}
