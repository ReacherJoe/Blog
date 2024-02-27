import axios from "axios";
import { config } from "../config/config";
import { User } from "../typings/types";

export const getInfiniteUsers = async ({ pageParam = 0 }) => {
  const res = await axios.get(
    `${config.apiBaseUrl}/users/infinite?cursor=${pageParam}`
  );

  const hasNext = pageParam <= res.data.totalPages;
  return {
    data: res.data.users,
    nextCursor: hasNext ? pageParam + 1 : undefined,
    totalPages: res.data.totalPages,
  };
};

export const getUser = async (id) => {
  const res = await axios.get(`${config.apiBaseUrl}/users/${id}`);
  return res.data ;
};

export const getUsers = async () => {
  const res = await axios.get(`${config.apiBaseUrl}/users`);

  return res.data;
};

export const getSearchedUsers = async (query) => {
  const res = await axios.get(`${config.apiBaseUrl}/users/search?q=${query}`);

  return res.data;
};

export const deleteUser = async (id) => {
  const accessToken = localStorage.getItem("accessToken");

  const res = await axios.delete(`${config.apiBaseUrl}/users/${id}`, {
    headers: { Authorization: `Bearer ${accessToken}` },
  });

  return res.data;
};

export const updateUser = async ({ data, userId }) => {
  const accessToken = localStorage.getItem("accessToken");

  const res = await axios.put(`${config.apiBaseUrl}/users/${userId}`, data, {
    headers: { Authorization: `Bearer ${accessToken}` },
  });

  return res.data;
};

export const createAdmin = async (data) => {
  const accessToken = localStorage.getItem("accessToken");

  const res = await axios.post(`${config.apiBaseUrl}/users`, data, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${accessToken}`,
    },
  });

  return res.data;
};
